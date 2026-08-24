//go:build cgo && !opus_shared

/* Deep PLC's linear layers, built for SSE2.

   Its own translation unit rather than dnn.c's, because nnet_arch.h is a
   template instantiated by whatever RTCD_ARCH names — dnn.c already
   instantiates it once as `c` through nnet_default.c, and a second pass in the
   same unit would redefine every static in it.

   This buys no speed: nnet_arch.h reaches its arithmetic through vec.h, which
   picks vec_avx.h from the compiler's own __SSE2__ and so was already
   vectorised before any of this. It is here because config.h presuming SSE2
   makes nnet.h route compute_linear through dnn/x86/dnn_x86.h, and the symbol
   it names then has to exist. */

#include "opus-1.5.2/dnn/x86/nnet_sse2.c"
