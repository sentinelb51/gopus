//go:build cgo && !opus_shared

/* celt's and silk's Neon paths, the arm64 counterpart of simd_amd64.c and apart
   from the amalgamation for the same reason: one translation unit is shared by
   every architecture, and a GOARCH filename suffix is the only constraint cgo
   reads from a C file.

   silk has more to gain here than it does from SSE. Its quantiser and its
   biquad are fixed-point in a float build too, so silk_NSQ_del_dec,
   silk_biquad_alt_stride2 and silk_LPC_inverse_pred_gain are hot either way —
   and upstream writes them for Neon but only for SSE4.1, which config.h does
   not presume. The asymmetry is upstream's, not this build's.

   arm_silk_map.c is not here: it is the runtime dispatch table, and its whole
   body is behind OPUS_HAVE_RTCD, which stays undefined. */

#include "opus-1.6.1/celt/arm/celt_neon_intr.c"
#include "opus-1.6.1/celt/arm/pitch_neon_intr.c"
#include "opus-1.6.1/silk/arm/biquad_alt_neon_intr.c"
#include "opus-1.6.1/silk/arm/LPC_inv_pred_gain_neon_intr.c"
#include "opus-1.6.1/silk/arm/NSQ_del_dec_neon_intr.c"
#include "opus-1.6.1/silk/arm/NSQ_neon.c"
