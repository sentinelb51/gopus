/* gopus: upstream's bbwenet_data.c reduced to what USE_WEIGHTS_FILE
   compiles. The weights it would otherwise carry are in
   opus_data.bin. Regenerate with update-opus.sh. */

#include "config.h"
#include "bbwenet_data.h"
#include "nnet.h"

#ifndef DUMP_BINARY_WEIGHTS
int init_bbwenetlayers(BBWENETLayers *model, const WeightArray *arrays) {
    if (linear_init(&model->bbwenet_fnet_conv1, arrays, "bbwenet_fnet_conv1_bias", NULL, NULL,"bbwenet_fnet_conv1_weights_float", NULL, NULL, NULL, 342, 128)) return 1;
    if (linear_init(&model->bbwenet_fnet_conv2, arrays, "bbwenet_fnet_conv2_bias", "bbwenet_fnet_conv2_subias", "bbwenet_fnet_conv2_weights_int8","bbwenet_fnet_conv2_weights_float", NULL, NULL, "bbwenet_fnet_conv2_scale", 384, 128)) return 1;
    if (linear_init(&model->bbwenet_fnet_gru_input, arrays, "bbwenet_fnet_gru_input_bias", "bbwenet_fnet_gru_input_subias", "bbwenet_fnet_gru_input_weights_int8","bbwenet_fnet_gru_input_weights_float", NULL, NULL, "bbwenet_fnet_gru_input_scale", 128, 384)) return 1;
    if (linear_init(&model->bbwenet_fnet_gru_recurrent, arrays, "bbwenet_fnet_gru_recurrent_bias", "bbwenet_fnet_gru_recurrent_subias", "bbwenet_fnet_gru_recurrent_weights_int8","bbwenet_fnet_gru_recurrent_weights_float", NULL, NULL, "bbwenet_fnet_gru_recurrent_scale", 128, 384)) return 1;
    if (linear_init(&model->bbwenet_fnet_tconv, arrays, "bbwenet_fnet_tconv_bias", "bbwenet_fnet_tconv_subias", "bbwenet_fnet_tconv_weights_int8","bbwenet_fnet_tconv_weights_float", NULL, NULL, "bbwenet_fnet_tconv_scale", 128, 256)) return 1;
    if (linear_init(&model->bbwenet_tdshape1_alpha1_f, arrays, "bbwenet_tdshape1_alpha1_f_bias", "bbwenet_tdshape1_alpha1_f_subias", "bbwenet_tdshape1_alpha1_f_weights_int8","bbwenet_tdshape1_alpha1_f_weights_float", NULL, NULL, "bbwenet_tdshape1_alpha1_f_scale", 256, 80)) return 1;
    if (linear_init(&model->bbwenet_tdshape1_alpha1_t, arrays, "bbwenet_tdshape1_alpha1_t_bias", NULL, NULL,"bbwenet_tdshape1_alpha1_t_weights_float", NULL, NULL, NULL, 42, 80)) return 1;
    if (linear_init(&model->bbwenet_tdshape1_alpha2, arrays, "bbwenet_tdshape1_alpha2_bias", NULL, NULL,"bbwenet_tdshape1_alpha2_weights_float", NULL, NULL, NULL, 160, 80)) return 1;
    if (linear_init(&model->bbwenet_tdshape2_alpha1_f, arrays, "bbwenet_tdshape2_alpha1_f_bias", "bbwenet_tdshape2_alpha1_f_subias", "bbwenet_tdshape2_alpha1_f_weights_int8","bbwenet_tdshape2_alpha1_f_weights_float", NULL, NULL, "bbwenet_tdshape2_alpha1_f_scale", 256, 120)) return 1;
    if (linear_init(&model->bbwenet_tdshape2_alpha1_t, arrays, "bbwenet_tdshape2_alpha1_t_bias", NULL, NULL,"bbwenet_tdshape2_alpha1_t_weights_float", NULL, NULL, NULL, 42, 120)) return 1;
    if (linear_init(&model->bbwenet_tdshape2_alpha2, arrays, "bbwenet_tdshape2_alpha2_bias", NULL, NULL,"bbwenet_tdshape2_alpha2_weights_float", NULL, NULL, NULL, 240, 120)) return 1;
    if (linear_init(&model->bbwenet_af1_kernel, arrays, "bbwenet_af1_kernel_bias", "bbwenet_af1_kernel_subias", "bbwenet_af1_kernel_weights_int8","bbwenet_af1_kernel_weights_float", NULL, NULL, "bbwenet_af1_kernel_scale", 128, 48)) return 1;
    if (linear_init(&model->bbwenet_af1_gain, arrays, "bbwenet_af1_gain_bias", NULL, NULL,"bbwenet_af1_gain_weights_float", NULL, NULL, NULL, 128, 3)) return 1;
    if (linear_init(&model->bbwenet_af2_kernel, arrays, "bbwenet_af2_kernel_bias", "bbwenet_af2_kernel_subias", "bbwenet_af2_kernel_weights_int8","bbwenet_af2_kernel_weights_float", NULL, NULL, "bbwenet_af2_kernel_scale", 128, 288)) return 1;
    if (linear_init(&model->bbwenet_af2_gain, arrays, "bbwenet_af2_gain_bias", NULL, NULL,"bbwenet_af2_gain_weights_float", NULL, NULL, NULL, 128, 3)) return 1;
    if (linear_init(&model->bbwenet_af3_kernel, arrays, "bbwenet_af3_kernel_bias", "bbwenet_af3_kernel_subias", "bbwenet_af3_kernel_weights_int8","bbwenet_af3_kernel_weights_float", NULL, NULL, "bbwenet_af3_kernel_scale", 128, 48)) return 1;
    if (linear_init(&model->bbwenet_af3_gain, arrays, "bbwenet_af3_gain_bias", NULL, NULL,"bbwenet_af3_gain_weights_float", NULL, NULL, NULL, 128, 1)) return 1;
    return 0;
}
#endif /* DUMP_BINARY_WEIGHTS */
