/* gopus: upstream's plc_data.c reduced to what USE_WEIGHTS_FILE
   compiles. The weights it would otherwise carry are in
   opus_data.bin. Regenerate with update-opus.sh. */

#include "config.h"
#include "plc_data.h"
#include "nnet.h"

#ifndef DUMP_BINARY_WEIGHTS
int init_plcmodel(PLCModel *model, const WeightArray *arrays) {
    if (linear_init(&model->plc_dense_in, arrays, "plc_dense_in_bias", NULL, NULL,"plc_dense_in_weights_float", NULL, NULL, NULL, 57, 128)) return 1;
    if (linear_init(&model->plc_dense_out, arrays, "plc_dense_out_bias", NULL, NULL,"plc_dense_out_weights_float", NULL, NULL, NULL, 192, 20)) return 1;
    if (linear_init(&model->plc_gru1_input, arrays, "plc_gru1_input_bias", "plc_gru1_input_subias", "plc_gru1_input_weights_int8","plc_gru1_input_weights_float", NULL, NULL, "plc_gru1_input_scale", 128, 576)) return 1;
    if (linear_init(&model->plc_gru1_recurrent, arrays, "plc_gru1_recurrent_bias", "plc_gru1_recurrent_subias", "plc_gru1_recurrent_weights_int8","plc_gru1_recurrent_weights_float", NULL, NULL, "plc_gru1_recurrent_scale", 192, 576)) return 1;
    if (linear_init(&model->plc_gru2_input, arrays, "plc_gru2_input_bias", "plc_gru2_input_subias", "plc_gru2_input_weights_int8","plc_gru2_input_weights_float", NULL, NULL, "plc_gru2_input_scale", 192, 576)) return 1;
    if (linear_init(&model->plc_gru2_recurrent, arrays, "plc_gru2_recurrent_bias", "plc_gru2_recurrent_subias", "plc_gru2_recurrent_weights_int8","plc_gru2_recurrent_weights_float", NULL, NULL, "plc_gru2_recurrent_scale", 192, 576)) return 1;
    return 0;
}
#endif /* DUMP_BINARY_WEIGHTS */
