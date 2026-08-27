//go:build cgo && !opus_shared

// The vendored build, and the default on every architecture: libopus is
// compiled from the tree beside this file, so a fresh clone needs no system
// library and no pkg-config. It used to be limited to amd64 and 386, which sent
// arm64 — Apple Silicon, and Linux arm64 — down the shared path and made a
// system libopus a build requirement there. Nothing about the C needs that.
//
// The architecture-specific intrinsics are in simd_amd64.c and simd_arm64.c,
// which cannot live in the amalgamation below: this file has no GOARCH
// constraint and they must. What selects them is config.h.
//
// Build with -tags opus_shared to link the system library instead.

package gopus

// #cgo linux darwin freebsd openbsd LDFLAGS: -lm
//
// #cgo CFLAGS: -Iopus-1.5.2 -Iopus-1.5.2/include -Iopus-1.5.2/celt -Iopus-1.5.2/silk -Iopus-1.5.2/silk/float
// #cgo CFLAGS: -Iopus-1.5.2/dnn
// #cgo CFLAGS: -DOPUS_BUILD -DHAVE_CONFIG_H -Wno-unused-function
// #cgo CFLAGS: -DCELT_ENCODER_C= -DCELT_DECODER_C=
//
// /* cgo compiles with `-g -O2` and appends whatever is here, so the later -O
//    wins. Worth 4-8 % of the codec on its own.
//
//    What is *not* here, so nobody spends the afternoon again. `-Ofast` does not
//    compile: celt/arch.h:201 refuses it outright — "Cannot build libopus with
//    -ffast-math unless FLOAT_APPROX is defined. This could result in crashes on
//    extreme (e.g. NaN) input". `-fno-math-errno` and `-fno-trapping-math` are
//    the two halves of -ffast-math that *are* safe here (libopus reads errno
//    from no math function and arms no FP trap) and are worth another ~6 % of a
//    concealed frame, but cgo's flag allowlist rejects every `-f` flag outside
//    its own list, so a library cannot carry them; they need CGO_CFLAGS from
//    whoever is building. `-flto` passes the allowlist and produces a binary
//    Windows will not load, and would buy nothing regardless — the amalgamation
//    below is already one translation unit. */
// #cgo CFLAGS: -O3
//
// /* Deep PLC: libopus 1.5's neural packet-loss concealment, which reconstructs a
//    lost frame as speech rather than extrapolating the last pitch period. It is
//    compiled in and switched at *runtime* — opus_decoder.c reads
//    `complexity >= 5`, and a decoder's complexity is 0 unless it is set — so this
//    costs nothing until Decoder.SetComplexity asks for it.
//
//    DRED and OSCE are deliberately left out. DRED needs the sender to enable it
//    too, so it buys nothing against a far end that is not also this library, and
//    the two of them are 10 MB of model data against Deep PLC's 5.
//
//    The sources are in dnn.c rather than in the amalgamation below: they are
//    written to be their own translation units and collide with celt and silk
//    when they are not — dnn/burg.c has its own static silk_energy_FLP, and
//    dnn/freq.c its own eband5ms. */
// #cgo CFLAGS: -DENABLE_DEEP_PLC
//
// /* Every source below is included into one translation unit, so celt.h is
//    parsed exactly once, by whichever file reaches it first. Its custom-mode
//    prototypes are gated on CELT_ENCODER_C / CELT_DECODER_C, which upstream
//    defines at the top of celt_encoder.c and celt_decoder.c and which therefore
//    arrive too late here. Defining both up front is what lets the amalgamation
//    see them; the two macros gate nothing else in the tree.
//
//    The trailing `=` is not a typo. A bare -D defines the macro as 1, which the
//    source's own empty #define then contradicts and gcc reports; an empty value
//    matches it exactly and is a benign redefinition. */
//
// #include "opus-1.5.2/config.h"
//
// #include "opus-1.5.2/celt/bands.c"
// #include "opus-1.5.2/celt/celt.c"
// #include "opus-1.5.2/celt/celt_encoder.c"
// #include "opus-1.5.2/celt/celt_decoder.c"
// #include "opus-1.5.2/celt/cwrs.c"
// #include "opus-1.5.2/celt/entcode.c"
// #include "opus-1.5.2/celt/entdec.c"
// #include "opus-1.5.2/celt/entenc.c"
// #include "opus-1.5.2/celt/kiss_fft.c"
// #include "opus-1.5.2/celt/laplace.c"
// #include "opus-1.5.2/celt/mathops.c"
// #include "opus-1.5.2/celt/mdct.c"
// #include "opus-1.5.2/celt/modes.c"
// #include "opus-1.5.2/celt/pitch.c"
// #include "opus-1.5.2/celt/celt_lpc.c"
// #include "opus-1.5.2/celt/quant_bands.c"
// #include "opus-1.5.2/celt/rate.c"
// #include "opus-1.5.2/celt/vq.c"
// #include "opus-1.5.2/silk/CNG.c"
// #include "opus-1.5.2/silk/code_signs.c"
// #include "opus-1.5.2/silk/init_decoder.c"
// #include "opus-1.5.2/silk/decode_core.c"
// #include "opus-1.5.2/silk/decode_frame.c"
// #include "opus-1.5.2/silk/decode_parameters.c"
// #include "opus-1.5.2/silk/decode_indices.c"
// #include "opus-1.5.2/silk/decode_pulses.c"
// #include "opus-1.5.2/silk/decoder_set_fs.c"
// #include "opus-1.5.2/silk/dec_API.c"
// #include "opus-1.5.2/silk/enc_API.c"
// #include "opus-1.5.2/silk/encode_indices.c"
// #include "opus-1.5.2/silk/encode_pulses.c"
// #include "opus-1.5.2/silk/gain_quant.c"
// #include "opus-1.5.2/silk/interpolate.c"
// #include "opus-1.5.2/silk/LP_variable_cutoff.c"
// #include "opus-1.5.2/silk/NLSF_decode.c"
// #include "opus-1.5.2/silk/NSQ.c"
// #include "opus-1.5.2/silk/NSQ_del_dec.c"
// #include "opus-1.5.2/silk/PLC.c"
// #include "opus-1.5.2/silk/shell_coder.c"
// #include "opus-1.5.2/silk/tables_gain.c"
// #include "opus-1.5.2/silk/tables_LTP.c"
// #include "opus-1.5.2/silk/tables_NLSF_CB_NB_MB.c"
// #include "opus-1.5.2/silk/tables_NLSF_CB_WB.c"
// #include "opus-1.5.2/silk/tables_other.c"
// #include "opus-1.5.2/silk/tables_pitch_lag.c"
// #include "opus-1.5.2/silk/tables_pulses_per_block.c"
// #include "opus-1.5.2/silk/VAD.c"
// #include "opus-1.5.2/silk/control_audio_bandwidth.c"
// #include "opus-1.5.2/silk/quant_LTP_gains.c"
// #include "opus-1.5.2/silk/VQ_WMat_EC.c"
// #include "opus-1.5.2/silk/HP_variable_cutoff.c"
// #include "opus-1.5.2/silk/NLSF_encode.c"
// #include "opus-1.5.2/silk/NLSF_VQ.c"
// #include "opus-1.5.2/silk/NLSF_unpack.c"
// #include "opus-1.5.2/silk/NLSF_del_dec_quant.c"
// #include "opus-1.5.2/silk/process_NLSFs.c"
// #include "opus-1.5.2/silk/stereo_LR_to_MS.c"
// #include "opus-1.5.2/silk/stereo_MS_to_LR.c"
// #include "opus-1.5.2/silk/check_control_input.c"
// #include "opus-1.5.2/silk/control_SNR.c"
// #include "opus-1.5.2/silk/init_encoder.c"
// #include "opus-1.5.2/silk/control_codec.c"
// #include "opus-1.5.2/silk/A2NLSF.c"
// #include "opus-1.5.2/silk/ana_filt_bank_1.c"
// #include "opus-1.5.2/silk/biquad_alt.c"
// #include "opus-1.5.2/silk/bwexpander_32.c"
// #include "opus-1.5.2/silk/bwexpander.c"
// #include "opus-1.5.2/silk/debug.c"
// #include "opus-1.5.2/silk/decode_pitch.c"
// #include "opus-1.5.2/silk/inner_prod_aligned.c"
// #include "opus-1.5.2/silk/lin2log.c"
// #include "opus-1.5.2/silk/log2lin.c"
// #include "opus-1.5.2/silk/LPC_analysis_filter.c"
// #include "opus-1.5.2/silk/LPC_inv_pred_gain.c"
// #include "opus-1.5.2/silk/table_LSF_cos.c"
// #include "opus-1.5.2/silk/NLSF2A.c"
// #include "opus-1.5.2/silk/NLSF_stabilize.c"
// #include "opus-1.5.2/silk/NLSF_VQ_weights_laroia.c"
// #include "opus-1.5.2/silk/pitch_est_tables.c"
// #include "opus-1.5.2/silk/resampler.c"
// #include "opus-1.5.2/silk/resampler_down2_3.c"
// #include "opus-1.5.2/silk/resampler_down2.c"
// #include "opus-1.5.2/silk/resampler_private_AR2.c"
// #include "opus-1.5.2/silk/resampler_private_down_FIR.c"
// #include "opus-1.5.2/silk/resampler_private_IIR_FIR.c"
// #include "opus-1.5.2/silk/resampler_private_up2_HQ.c"
// #include "opus-1.5.2/silk/resampler_rom.c"
// #include "opus-1.5.2/silk/sigm_Q15.c"
// #include "opus-1.5.2/silk/sort.c"
// #include "opus-1.5.2/silk/sum_sqr_shift.c"
// #include "opus-1.5.2/silk/stereo_decode_pred.c"
// #include "opus-1.5.2/silk/stereo_encode_pred.c"
// #include "opus-1.5.2/silk/stereo_find_predictor.c"
// #include "opus-1.5.2/silk/stereo_quant_pred.c"
// #include "opus-1.5.2/silk/LPC_fit.c"
// #include "opus-1.5.2/silk/float/apply_sine_window_FLP.c"
// #include "opus-1.5.2/silk/float/corrMatrix_FLP.c"
// #include "opus-1.5.2/silk/float/encode_frame_FLP.c"
// #include "opus-1.5.2/silk/float/find_LPC_FLP.c"
// #include "opus-1.5.2/silk/float/find_LTP_FLP.c"
// #include "opus-1.5.2/silk/float/find_pitch_lags_FLP.c"
// #include "opus-1.5.2/silk/float/find_pred_coefs_FLP.c"
// #include "opus-1.5.2/silk/float/LPC_analysis_filter_FLP.c"
// #include "opus-1.5.2/silk/float/LTP_analysis_filter_FLP.c"
// #include "opus-1.5.2/silk/float/LTP_scale_ctrl_FLP.c"
// #include "opus-1.5.2/silk/float/noise_shape_analysis_FLP.c"
// #include "opus-1.5.2/silk/float/process_gains_FLP.c"
// #include "opus-1.5.2/silk/float/regularize_correlations_FLP.c"
// #include "opus-1.5.2/silk/float/residual_energy_FLP.c"
// #include "opus-1.5.2/silk/float/warped_autocorrelation_FLP.c"
// #include "opus-1.5.2/silk/float/wrappers_FLP.c"
// #include "opus-1.5.2/silk/float/autocorrelation_FLP.c"
// #include "opus-1.5.2/silk/float/burg_modified_FLP.c"
// #include "opus-1.5.2/silk/float/bwexpander_FLP.c"
// #include "opus-1.5.2/silk/float/energy_FLP.c"
// #include "opus-1.5.2/silk/float/inner_product_FLP.c"
// #include "opus-1.5.2/silk/float/k2a_FLP.c"
// #include "opus-1.5.2/silk/float/LPC_inv_pred_gain_FLP.c"
// #include "opus-1.5.2/silk/float/pitch_analysis_core_FLP.c"
// #include "opus-1.5.2/silk/float/scale_copy_vector_FLP.c"
// #include "opus-1.5.2/silk/float/scale_vector_FLP.c"
// #include "opus-1.5.2/silk/float/schur_FLP.c"
// #include "opus-1.5.2/silk/float/sort_FLP.c"
// #include "opus-1.5.2/src/opus.c"
// #include "opus-1.5.2/src/opus_decoder.c"
// #include "opus-1.5.2/src/opus_encoder.c"
// #include "opus-1.5.2/src/extensions.c"
// #include "opus-1.5.2/src/opus_multistream.c"
// #include "opus-1.5.2/src/opus_multistream_encoder.c"
// #include "opus-1.5.2/src/opus_multistream_decoder.c"
// #include "opus-1.5.2/src/repacketizer.c"
// #include "opus-1.5.2/src/opus_projection_encoder.c"
// #include "opus-1.5.2/src/opus_projection_decoder.c"
// #include "opus-1.5.2/src/mapping_matrix.c"
// #include "opus-1.5.2/src/analysis.c"
// #include "opus-1.5.2/src/mlp.c"
// #include "opus-1.5.2/src/mlp_data.c"
//
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
// mid-call.
//
// It is not free on a clean stream, whatever this comment used to say. The
// model has to be fed on every frame that arrives (celt_decoder.c calls
// update_plc_state on each one), so a decoder at ComplexityDeepPLC costs about
// a quarter more per good frame than one at ComplexityOff — 17.6 µs against
// 14.2 on a 9950X3D, 48 kHz mono. Concealing is where the rest of it goes: one
// hidden frame is ~227 µs against ~17, thirteen times the price of the
// extrapolation it replaces. Both figures are with march_amd64.go's AVX2 floor;
// without it concealment is ~690 µs.
//
// So it is worth having and it is worth being a setting, but a client running
// it for every participant of a large call is spending real time on it.
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
