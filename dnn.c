//go:build cgo && !opus_shared

/* libopus 1.5's Deep PLC, compiled apart from the amalgamation in
   opus_nonshared.go.

   These are upstream's own DEEP_PLC_SOURCES and are written to be separate
   translation units: dnn/burg.c defines a static silk_energy_FLP of its own and
   dnn/freq.c a static eband5ms, both of which clash with the silk and celt
   copies when everything shares one unit. cgo compiles every .c in the package
   directory, so putting them here is the whole fix, and the build tag matches
   the one on the amalgamation so -tags opus_shared still links the system
   library and compiles none of this. */

#include "opus-1.5.2/dnn/burg.c"
#include "opus-1.5.2/dnn/freq.c"
#include "opus-1.5.2/dnn/fargan.c"
#include "opus-1.5.2/dnn/fargan_data.c"
#include "opus-1.5.2/dnn/lpcnet_enc.c"
#include "opus-1.5.2/dnn/lpcnet_plc.c"
#include "opus-1.5.2/dnn/lpcnet_tables.c"
#include "opus-1.5.2/dnn/nnet.c"
#include "opus-1.5.2/dnn/nnet_default.c"
#include "opus-1.5.2/dnn/plc_data.c"
#include "opus-1.5.2/dnn/parse_lpcnet_weights.c"
#include "opus-1.5.2/dnn/pitchdnn.c"
#include "opus-1.5.2/dnn/pitchdnn_data.c"
