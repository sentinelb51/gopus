//go:build cgo && !opus_shared && !opus_baseline

// The x86-64-v3 floor — AVX2, FMA and BMI, which is every amd64 part from
// Haswell and Excavator (2013) onward.
//
// It is here for Deep PLC, not for the codec. `dnn/vec_avx.h` picks its vector
// width from the compiler's own `__AVX2__` and `__FMA__` at compile time, not
// from libopus's config, so without this the neural model runs that header's
// SSE2 fallback: concealing one frame costs ~690 µs instead of ~220, and a
// decoder at complexity >= 5 pays 36 % more on every *good* frame besides,
// having to feed the model. celt and silk gain about 9 % from the same flag,
// which on its own would not be worth a floor.
//
// Build with `-tags opus_baseline` for a plain x86-64 object. Below the floor
// the failure is SIGILL, not a slow decode, so the tag is the answer for any
// machine older than Haswell.
//
// Deliberately not `-march=native`: on Zen 5 it is 13 % *slower* than this,
// because `vec_avx.h:623` takes the hardware VNNI dot product wherever one
// exists and that instruction loses to the AVX2 sequence it replaces —
// measured, and confirmed by `-mno-avxvnni` recovering every bit of it. Nor
// `x86-64-v4`, which is a further 3 % and rules out every Intel consumer part
// since Rocket Lake.

package gopus

// #cgo CFLAGS: -march=x86-64-v3
import "C"
