//go:build cgo && !opus_shared

/* celt's SSE paths, compiled apart from the amalgamation in opus_nonshared.go.

   Apart because the amalgamation is one translation unit for every
   architecture, and these are intrinsics: the `_amd64` suffix is what keeps
   them out of an arm64 or 386 build, and a filename is the only place cgo takes
   that constraint from. Go applies GOARCH suffixes to C files as it does to Go
   ones, so nothing here needs a build tag beyond the one above.

   What reaches these instead of the generic C is decided in config.h, which
   presumes SSE and SSE2 on amd64 and so resolves celt's dispatch macros
   straight to the symbols below — xcorr_kernel, celt_inner_prod,
   dual_inner_prod and comb_filter_const from pitch_sse.c, op_pvq_search from
   vq_sse2.c. Both files are guarded on !FIXED_POINT upstream, which this build
   is. The SSE4.1 and AVX2 levels above them are absent on purpose; config.h
   says why. */

#include "opus-1.6.1/celt/x86/pitch_sse.c"
#include "opus-1.6.1/celt/x86/vq_sse2.c"
