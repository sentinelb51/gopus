//go:build cgo && !opus_shared

/* libopus's neural extensions, compiled apart from the amalgamation in
   opus_nonshared.go.

   These are upstream's own DEEP_PLC_SOURCES and OSCE_SOURCES, written to be
   separate translation units: dnn/burg.c defines a static silk_energy_FLP of
   its own and dnn/freq.c a static eband5ms, both of which clash with the silk
   and celt copies when everything shares one unit. cgo compiles every .c in the
   package directory, so putting them here is the whole fix, and the build tag
   matches the one on the amalgamation so -tags opus_shared still links the
   system library and compiles none of this.

   The *_data.c files are the ones update-opus.sh reduces to their init
   functions: under USE_WEIGHTS_FILE the weights come from opus_data.bin
   instead. */

#include "opus-1.6.1/dnn/burg.c"
/* burg.c names its own SILK_MAX_ORDER_LPC 16, uses it for four stack arrays and
   nothing else; silk/SigProc_FIX.h, which osce.h reaches through silk/structs.h
   further down, calls it 24. Upstream never has the two in one unit. Dropping
   burg.c's the moment it is done with it leaves each definition alone rather
   than leaving a redefinition warning to be read past every build. */
#undef SILK_MAX_ORDER_LPC
#include "opus-1.6.1/dnn/freq.c"
#include "opus-1.6.1/dnn/fargan.c"
#include "opus-1.6.1/dnn/fargan_data.c"
#include "opus-1.6.1/dnn/lpcnet_enc.c"
#include "opus-1.6.1/dnn/lpcnet_plc.c"
#include "opus-1.6.1/dnn/lpcnet_tables.c"
#include "opus-1.6.1/dnn/nnet.c"
#include "opus-1.6.1/dnn/nnet_default.c"
#include "opus-1.6.1/dnn/plc_data.c"
#include "opus-1.6.1/dnn/parse_lpcnet_weights.c"
#include "opus-1.6.1/dnn/pitchdnn.c"
#include "opus-1.6.1/dnn/pitchdnn_data.c"

/* OSCE: LACE and NoLACE enhance decoded SILK, and bbwenet extends a wideband
   frame to fullband. New in 1.6 is the last of those. */
#include "opus-1.6.1/dnn/nndsp.c"
#include "opus-1.6.1/dnn/osce.c"
#include "opus-1.6.1/dnn/osce_features.c"
#include "opus-1.6.1/dnn/lace_data.c"
#include "opus-1.6.1/dnn/nolace_data.c"
#include "opus-1.6.1/dnn/bbwenet_data.c"
