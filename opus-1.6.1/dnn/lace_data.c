/* gopus: upstream's lace_data.c reduced to what USE_WEIGHTS_FILE
   compiles. The weights it would otherwise carry are in
   opus_data.bin. Regenerate with update-opus.sh. */

#include "config.h"
#include "lace_data.h"
#include "nnet.h"

#ifndef DUMP_BINARY_WEIGHTS
int init_lacelayers(LACELayers *model, const WeightArray *arrays) {
    if (linear_init(&model->lace_pitch_embedding, arrays, "lace_pitch_embedding_bias", NULL, NULL,"lace_pitch_embedding_weights_float", NULL, NULL, NULL, 301, 64)) return 1;
    if (linear_init(&model->lace_fnet_conv1, arrays, "lace_fnet_conv1_bias", NULL, NULL,"lace_fnet_conv1_weights_float", NULL, NULL, NULL, 173, 96)) return 1;
    if (linear_init(&model->lace_fnet_conv2, arrays, "lace_fnet_conv2_bias", "lace_fnet_conv2_subias", "lace_fnet_conv2_weights_int8","lace_fnet_conv2_weights_float", NULL, NULL, "lace_fnet_conv2_scale", 768, 128)) return 1;
    if (linear_init(&model->lace_fnet_tconv, arrays, "lace_fnet_tconv_bias", "lace_fnet_tconv_subias", "lace_fnet_tconv_weights_int8","lace_fnet_tconv_weights_float", NULL, NULL, "lace_fnet_tconv_scale", 128, 512)) return 1;
    if (linear_init(&model->lace_fnet_gru_input, arrays, "lace_fnet_gru_input_bias", "lace_fnet_gru_input_subias", "lace_fnet_gru_input_weights_int8","lace_fnet_gru_input_weights_float", NULL, NULL, "lace_fnet_gru_input_scale", 128, 384)) return 1;
    if (linear_init(&model->lace_fnet_gru_recurrent, arrays, "lace_fnet_gru_recurrent_bias", "lace_fnet_gru_recurrent_subias", "lace_fnet_gru_recurrent_weights_int8","lace_fnet_gru_recurrent_weights_float", NULL, NULL, "lace_fnet_gru_recurrent_scale", 128, 384)) return 1;
    if (linear_init(&model->lace_cf1_kernel, arrays, "lace_cf1_kernel_bias", "lace_cf1_kernel_subias", "lace_cf1_kernel_weights_int8","lace_cf1_kernel_weights_float", NULL, NULL, "lace_cf1_kernel_scale", 128, 16)) return 1;
    if (linear_init(&model->lace_cf1_gain, arrays, "lace_cf1_gain_bias", NULL, NULL,"lace_cf1_gain_weights_float", NULL, NULL, NULL, 128, 1)) return 1;
    if (linear_init(&model->lace_cf1_global_gain, arrays, "lace_cf1_global_gain_bias", NULL, NULL,"lace_cf1_global_gain_weights_float", NULL, NULL, NULL, 128, 1)) return 1;
    if (linear_init(&model->lace_cf2_kernel, arrays, "lace_cf2_kernel_bias", "lace_cf2_kernel_subias", "lace_cf2_kernel_weights_int8","lace_cf2_kernel_weights_float", NULL, NULL, "lace_cf2_kernel_scale", 128, 16)) return 1;
    if (linear_init(&model->lace_cf2_gain, arrays, "lace_cf2_gain_bias", NULL, NULL,"lace_cf2_gain_weights_float", NULL, NULL, NULL, 128, 1)) return 1;
    if (linear_init(&model->lace_cf2_global_gain, arrays, "lace_cf2_global_gain_bias", NULL, NULL,"lace_cf2_global_gain_weights_float", NULL, NULL, NULL, 128, 1)) return 1;
    if (linear_init(&model->lace_af1_kernel, arrays, "lace_af1_kernel_bias", "lace_af1_kernel_subias", "lace_af1_kernel_weights_int8","lace_af1_kernel_weights_float", NULL, NULL, "lace_af1_kernel_scale", 128, 16)) return 1;
    if (linear_init(&model->lace_af1_gain, arrays, "lace_af1_gain_bias", NULL, NULL,"lace_af1_gain_weights_float", NULL, NULL, NULL, 128, 1)) return 1;
    return 0;
}
#endif /* DUMP_BINARY_WEIGHTS */
