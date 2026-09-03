//go:build cgo && !opus_shared

/* Deep PLC's linear layers, built for Neon: the arm64 counterpart of
   dnn_amd64.c, its own translation unit for the same reason, and buying the
   same nothing — vec.h picks vec_neon.h from the compiler's __ARM_NEON without
   being asked.

   Only compute_linear is routed here. dnn_arm.h gates compute_activation and
   compute_conv2d on OPUS_ARM_PRESUME_NEON, which is the ARM *assembly* switch
   and not this build's; they stay generic C. */

#include "opus-1.6.1/dnn/arm/nnet_neon.c"
