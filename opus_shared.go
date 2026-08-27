//go:build cgo && opus_shared

// The system-library build, opt-in through -tags opus_shared. It links whatever
// pkg-config points at rather than the vendored tree, which is what a
// distribution packaging this would want.

package gopus

// #cgo !nopkgconfig pkg-config: opus
//
// #include <opus.h>
// enum {
//   gopus_ok = OPUS_OK,
//   gopus_bad_arg = OPUS_BAD_ARG,
//   gopus_small_buffer = OPUS_BUFFER_TOO_SMALL,
//   gopus_internal = OPUS_INTERNAL_ERROR,
//   gopus_invalid_packet = OPUS_INVALID_PACKET,
//   gopus_unimplemented = OPUS_UNIMPLEMENTED,
//   gopus_invalid_state = OPUS_INVALID_STATE,
//   gopus_alloc_fail = OPUS_ALLOC_FAIL,
// };
//
//
// enum {
//   gopus_application_voip    = OPUS_APPLICATION_VOIP,
//   gopus_application_audio   = OPUS_APPLICATION_AUDIO,
//   gopus_restricted_lowdelay = OPUS_APPLICATION_RESTRICTED_LOWDELAY,
//   gopus_bitrate_max         = OPUS_BITRATE_MAX,
// };
//
//
// void gopus_setvbr(OpusEncoder *encoder, int vbr) {
//   opus_encoder_ctl(encoder, OPUS_SET_VBR(vbr));
// }
//
// int gopus_setinbandfec(OpusEncoder *encoder, int enabled) {
//   return opus_encoder_ctl(encoder, OPUS_SET_INBAND_FEC(enabled));
// }
//
// int gopus_setpacketlossperc(OpusEncoder *encoder, int percent) {
//   return opus_encoder_ctl(encoder, OPUS_SET_PACKET_LOSS_PERC(percent));
// }
//
// int gopus_setdtx(OpusEncoder *encoder, int enabled) {
//   return opus_encoder_ctl(encoder, OPUS_SET_DTX(enabled));
// }
//
// void gopus_setbitrate(OpusEncoder *encoder, int bitrate) {
//   opus_encoder_ctl(encoder, OPUS_SET_BITRATE(bitrate));
// }
//
// opus_int32 gopus_bitrate(OpusEncoder *encoder) {
//   opus_int32 bitrate;
//   opus_encoder_ctl(encoder, OPUS_GET_BITRATE(&bitrate));
//   return bitrate;
// }
//
// void gopus_setapplication(OpusEncoder *encoder, int application) {
//   opus_encoder_ctl(encoder, OPUS_SET_APPLICATION(application));
// }
//
// opus_int32 gopus_application(OpusEncoder *encoder) {
//   opus_int32 application;
//   opus_encoder_ctl(encoder, OPUS_GET_APPLICATION(&application));
//   return application;
// }
//
// void gopus_encoder_resetstate(OpusEncoder *encoder) {
//   opus_encoder_ctl(encoder, OPUS_RESET_STATE);
// }
//
// int gopus_setdecodercomplexity(OpusDecoder *decoder, int complexity) {
//   return opus_decoder_ctl(decoder, OPUS_SET_COMPLEXITY(complexity));
// }
//
// void gopus_decoder_resetstate(OpusDecoder *decoder) {
//   opus_decoder_ctl(decoder, OPUS_RESET_STATE);
// }
import "C"

import (
	"errors"
	"unsafe"
)

type Application int

const (
	Voip               Application = C.gopus_application_voip
	Audio              Application = C.gopus_application_audio
	RestrictedLowDelay Application = C.gopus_restricted_lowdelay
)

const (
	BitrateMaximum = C.gopus_bitrate_max
)

type Encoder struct {
	data     []byte
	cEncoder *C.struct_OpusEncoder
}

// Version is the libopus build this package is linked against, straight from
// opus_get_version_string. Worth having where the library is vendored: it is the
// only way to tell at runtime which copy actually got compiled in.
func Version() string {
	return C.GoString(C.opus_get_version_string())
}

func NewEncoder(sampleRate, channels int, application Application) (*Encoder, error) {
	encoder := &Encoder{}
	encoder.data = make([]byte, int(C.opus_encoder_get_size(C.int(channels))))
	encoder.cEncoder = (*C.struct_OpusEncoder)(unsafe.Pointer(&encoder.data[0]))

	ret := C.opus_encoder_init(encoder.cEncoder, C.opus_int32(sampleRate), C.int(channels), C.int(application))
	if err := getErr(ret); err != nil {
		return nil, err
	}
	return encoder, nil
}

func (e *Encoder) Encode(pcm []int16, frameSize, maxDataBytes int) ([]byte, error) {
	data := make([]byte, maxDataBytes)

	n, err := e.EncodeIn(pcm, frameSize, data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

// EncodeIn encodes into data rather than allocating, and reports how many bytes
// the packet took. len(data) caps the packet, so data should be as large as the
// maxDataBytes the caller would have passed to Encode.
//
// It exists for the same caller DecodeIn does: fifty frames a second through
// Encode is fifty allocations a second, on a path that otherwise makes none.
func (e *Encoder) EncodeIn(pcm []int16, frameSize int, data []byte) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("gopus: empty packet buffer")
	}

	pcmPtr := (*C.opus_int16)(unsafe.Pointer(&pcm[0]))
	dataPtr := (*C.uchar)(unsafe.Pointer(&data[0]))

	encodedC := C.opus_encode(e.cEncoder, pcmPtr, C.int(frameSize), dataPtr, C.opus_int32(len(data)))
	encoded := int(encodedC)

	if encoded < 0 {
		return 0, getErr(C.int(encodedC))
	}
	return encoded, nil
}

func (e *Encoder) SetVbr(vbr bool) {
	var cVbr C.int
	if vbr {
		cVbr = 1
	} else {
		cVbr = 0
	}
	C.gopus_setvbr(e.cEncoder, cVbr)
}

func (e *Encoder) SetBitrate(bitrate int) {
	C.gopus_setbitrate(e.cEncoder, C.int(bitrate))
}

func (e *Encoder) Bitrate() int {
	return int(C.gopus_bitrate(e.cEncoder))
}

func (e *Encoder) SetApplication(application Application) {
	C.gopus_setapplication(e.cEncoder, C.int(application))
}

func (e *Encoder) Application() Application {
	return Application(C.gopus_application(e.cEncoder))
}

// SetInBandFEC turns on Opus's in-band forward error correction: each packet
// carries a low-bitrate copy of the one before it, so a single lost packet can be
// recovered from its successor rather than concealed.
//
// It costs bitrate only in proportion to SetPacketLossPerc, so the usual
// arrangement is to leave this on and tune the loss estimate instead.
func (e *Encoder) SetInBandFEC(enabled bool) error {
	return getErr(C.gopus_setinbandfec(e.cEncoder, cbool(enabled)))
}

// SetPacketLossPerc tells the encoder how much loss to expect, 0-100. It decides
// how much redundancy in-band FEC actually adds; at 0 the FEC flag does nothing.
func (e *Encoder) SetPacketLossPerc(percent int) error {
	return getErr(C.gopus_setpacketlossperc(e.cEncoder, C.int(percent)))
}

// SetDTX turns on discontinuous transmission: a silent frame goes as a few bytes
// of comfort noise rather than a full frame, so a gated microphone costs about
// 5 kbps instead of the full bitrate and the track stays alive.
func (e *Encoder) SetDTX(enabled bool) error {
	return getErr(C.gopus_setdtx(e.cEncoder, cbool(enabled)))
}

// cbool is C's idea of a boolean, which each CTL above takes as an int.
func cbool(v bool) C.int {
	if v {
		return 1
	}

	return 0
}

func (e *Encoder) ResetState() {
	C.gopus_encoder_resetstate(e.cEncoder)
}

type Decoder struct {
	data     []byte
	cDecoder *C.struct_OpusDecoder
	channels int
}

func NewDecoder(sampleRate, channels int) (*Decoder, error) {
	decoder := &Decoder{}
	decoder.data = make([]byte, int(C.opus_decoder_get_size(C.int(channels))))
	decoder.cDecoder = (*C.struct_OpusDecoder)(unsafe.Pointer(&decoder.data[0]))

	ret := C.opus_decoder_init(decoder.cDecoder, C.opus_int32(sampleRate), C.int(channels))
	if err := getErr(ret); err != nil {
		return nil, err
	}
	decoder.channels = channels

	return decoder, nil
}

func (d *Decoder) Decode(data []byte, frameSize int, fec bool) ([]int16, error) {
	output := make([]int16, d.channels*frameSize)

	n, err := d.DecodeIn(data, frameSize, output, fec)
	if err != nil {
		return nil, err
	}
	return output[:n], nil
}

// DecodeIn decodes into pcm rather than allocating, and reports how many
// samples it wrote — channels*frameSize for a whole frame. pcm must hold at
// least that many samples. A nil data conceals a lost packet and fec recovers
// one from its successor, exactly as with Decode.
//
// It exists for callers decoding tens of frames a second per stream, where
// Decode's per-call allocation is the receive path's dominant garbage.
func (d *Decoder) DecodeIn(data []byte, frameSize int, pcm []int16, fec bool) (int, error) {
	if len(pcm) < d.channels*frameSize {
		return 0, errors.New("gopus: pcm shorter than channels*frameSize")
	}

	var dataPtr *C.uchar
	if len(data) > 0 {
		dataPtr = (*C.uchar)(unsafe.Pointer(&data[0]))
	}
	dataLen := C.opus_int32(len(data))

	outputPtr := (*C.opus_int16)(unsafe.Pointer(&pcm[0]))

	var cFec C.int
	if fec {
		cFec = 1
	} else {
		cFec = 0
	}

	cRet := C.opus_decode(d.cDecoder, dataPtr, dataLen, outputPtr, C.int(frameSize), cFec)
	ret := int(cRet)

	if ret < 0 {
		return 0, getErr(cRet)
	}
	return ret * d.channels, nil
}

// Decoder complexity levels. Complexity is a decoder-side dial in libopus 1.5
// and above, and what it buys is loss concealment: below DeepPLC a lost packet
// is concealed the classic way — the last pitch period extrapolated and faded,
// which is convincing for one frame and robotic by the third — and at DeepPLC or
// above a neural model reconstructs it as speech in the talker's own voice.
//
// ComplexityOff is libopus's own default, so nothing changes for a caller that
// never asks. Deep PLC has to be compiled in as well; on a build without it, or
// on a system libopus older than 1.5, these are accepted and do nothing.
const (
	ComplexityOff     = 0
	ComplexityDeepPLC = 5
)

// SetComplexity sets how much work the decoder may do, 0-10. The level that
// matters is ComplexityDeepPLC, at and above which libopus conceals lost packets
// with its neural model rather than by extrapolation.
//
// It is per decoder and takes effect on the next frame, so it can be moved
// mid-call. The cost is paid only while concealing: a stream losing nothing
// decodes at exactly the same price either way.
func (d *Decoder) SetComplexity(complexity int) error {
	return getErr(C.gopus_setdecodercomplexity(d.cDecoder, C.int(complexity)))
}

func (d *Decoder) ResetState() {
	C.gopus_decoder_resetstate(d.cDecoder)
}

func CountFrames(data []byte) (int, error) {
	dataPtr := (*C.uchar)(unsafe.Pointer(&data[0]))
	cLen := C.opus_int32(len(data))

	cRet := C.opus_packet_get_nb_frames(dataPtr, cLen)
	if err := getErr(cRet); err != nil {
		return 0, err
	}
	return int(cRet), nil
}

var (
	ErrBadArgument   = errors.New("bad argument")
	ErrSmallBuffer   = errors.New("buffer is too small")
	ErrInternal      = errors.New("internal error")
	ErrInvalidPacket = errors.New("invalid packet")
	ErrUnimplemented = errors.New("unimplemented")
	ErrInvalidState  = errors.New("invalid state")
	ErrAllocFail     = errors.New("allocation failed")
	ErrUnknown       = errors.New("unknown error")
)

func getErr(code C.int) error {
	switch code {
	case C.gopus_ok:
		return nil
	case C.gopus_bad_arg:
		return ErrBadArgument
	case C.gopus_small_buffer:
		return ErrSmallBuffer
	case C.gopus_internal:
		return ErrInternal
	case C.gopus_invalid_packet:
		return ErrInvalidPacket
	case C.gopus_unimplemented:
		return ErrUnimplemented
	case C.gopus_invalid_state:
		return ErrInvalidState
	case C.gopus_alloc_fail:
		return ErrAllocFail
	default:
		return ErrUnknown
	}
}
