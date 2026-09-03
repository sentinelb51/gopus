/* gopus: upstream's nolace_data.c reduced to what USE_WEIGHTS_FILE
   compiles. The weights it would otherwise carry are in
   opus_data.bin. Regenerate with update-opus.sh. */

#include "config.h"
#include "nolace_data.h"
#include "nnet.h"

#ifndef DUMP_BINARY_WEIGHTS
int init_nolacelayers(NOLACELayers *model, const WeightArray *arrays) {
    if (linear_init(&model->nolace_pitch_embedding, arrays, "nolace_pitch_embedding_bias", NULL, NULL,"nolace_pitch_embedding_weights_float", NULL, NULL, NULL, 301, 64)) return 1;
    if (linear_init(&model->nolace_fnet_conv1, arrays, "nolace_fnet_conv1_bias", NULL, NULL,"nolace_fnet_conv1_weights_float", NULL, NULL, NULL, 173, 96)) return 1;
    if (linear_init(&model->nolace_fnet_conv2, arrays, "nolace_fnet_conv2_bias", "nolace_fnet_conv2_subias", "nolace_fnet_conv2_weights_int8","nolace_fnet_conv2_weights_float", NULL, NULL, "nolace_fnet_conv2_scale", 768, 160)) return 1;
    if (linear_init(&model->nolace_fnet_tconv, arrays, "nolace_fnet_tconv_bias", "nolace_fnet_tconv_subias", "nolace_fnet_tconv_weights_int8","nolace_fnet_tconv_weights_float", NULL, NULL, "nolace_fnet_tconv_scale", 160, 640)) return 1;
    if (linear_init(&model->nolace_fnet_gru_input, arrays, "nolace_fnet_gru_input_bias", "nolace_fnet_gru_input_subias", "nolace_fnet_gru_input_weights_int8","nolace_fnet_gru_input_weights_float", NULL, NULL, "nolace_fnet_gru_input_scale", 160, 480)) return 1;
    if (linear_init(&model->nolace_fnet_gru_recurrent, arrays, "nolace_fnet_gru_recurrent_bias", "nolace_fnet_gru_recurrent_subias", "nolace_fnet_gru_recurrent_weights_int8","nolace_fnet_gru_recurrent_weights_float", NULL, NULL, "nolace_fnet_gru_recurrent_scale", 160, 480)) return 1;
    if (linear_init(&model->nolace_cf1_kernel, arrays, "nolace_cf1_kernel_bias", "nolace_cf1_kernel_subias", "nolace_cf1_kernel_weights_int8","nolace_cf1_kernel_weights_float", NULL, NULL, "nolace_cf1_kernel_scale", 160, 16)) return 1;
    if (linear_init(&model->nolace_cf1_gain, arrays, "nolace_cf1_gain_bias", NULL, NULL,"nolace_cf1_gain_weights_float", NULL, NULL, NULL, 160, 1)) return 1;
    if (linear_init(&model->nolace_cf1_global_gain, arrays, "nolace_cf1_global_gain_bias", NULL, NULL,"nolace_cf1_global_gain_weights_float", NULL, NULL, NULL, 160, 1)) return 1;
    if (linear_init(&model->nolace_cf2_kernel, arrays, "nolace_cf2_kernel_bias", "nolace_cf2_kernel_subias", "nolace_cf2_kernel_weights_int8","nolace_cf2_kernel_weights_float", NULL, NULL, "nolace_cf2_kernel_scale", 160, 16)) return 1;
    if (linear_init(&model->nolace_cf2_gain, arrays, "nolace_cf2_gain_bias", NULL, NULL,"nolace_cf2_gain_weights_float", NULL, NULL, NULL, 160, 1)) return 1;
    if (linear_init(&model->nolace_cf2_global_gain, arrays, "nolace_cf2_global_gain_bias", NULL, NULL,"nolace_cf2_global_gain_weights_float", NULL, NULL, NULL, 160, 1)) return 1;
    if (linear_init(&model->nolace_af1_kernel, arrays, "nolace_af1_kernel_bias", "nolace_af1_kernel_subias", "nolace_af1_kernel_weights_int8","nolace_af1_kernel_weights_float", NULL, NULL, "nolace_af1_kernel_scale", 160, 32)) return 1;
    if (linear_init(&model->nolace_af1_gain, arrays, "nolace_af1_gain_bias", NULL, NULL,"nolace_af1_gain_weights_float", NULL, NULL, NULL, 160, 2)) return 1;
    if (linear_init(&model->nolace_tdshape1_alpha1_f, arrays, "nolace_tdshape1_alpha1_f_bias", "nolace_tdshape1_alpha1_f_subias", "nolace_tdshape1_alpha1_f_weights_int8","nolace_tdshape1_alpha1_f_weights_float", NULL, NULL, "nolace_tdshape1_alpha1_f_scale", 320, 80)) return 1;
    if (linear_init(&model->nolace_tdshape1_alpha1_t, arrays, "nolace_tdshape1_alpha1_t_bias", NULL, NULL,"nolace_tdshape1_alpha1_t_weights_float", NULL, NULL, NULL, 42, 80)) return 1;
    if (linear_init(&model->nolace_tdshape1_alpha2, arrays, "nolace_tdshape1_alpha2_bias", NULL, NULL,"nolace_tdshape1_alpha2_weights_float", NULL, NULL, NULL, 160, 80)) return 1;
    if (linear_init(&model->nolace_tdshape2_alpha1_f, arrays, "nolace_tdshape2_alpha1_f_bias", "nolace_tdshape2_alpha1_f_subias", "nolace_tdshape2_alpha1_f_weights_int8","nolace_tdshape2_alpha1_f_weights_float", NULL, NULL, "nolace_tdshape2_alpha1_f_scale", 320, 80)) return 1;
    if (linear_init(&model->nolace_tdshape2_alpha1_t, arrays, "nolace_tdshape2_alpha1_t_bias", NULL, NULL,"nolace_tdshape2_alpha1_t_weights_float", NULL, NULL, NULL, 42, 80)) return 1;
    if (linear_init(&model->nolace_tdshape2_alpha2, arrays, "nolace_tdshape2_alpha2_bias", NULL, NULL,"nolace_tdshape2_alpha2_weights_float", NULL, NULL, NULL, 160, 80)) return 1;
    if (linear_init(&model->nolace_tdshape3_alpha1_f, arrays, "nolace_tdshape3_alpha1_f_bias", "nolace_tdshape3_alpha1_f_subias", "nolace_tdshape3_alpha1_f_weights_int8","nolace_tdshape3_alpha1_f_weights_float", NULL, NULL, "nolace_tdshape3_alpha1_f_scale", 320, 80)) return 1;
    if (linear_init(&model->nolace_tdshape3_alpha1_t, arrays, "nolace_tdshape3_alpha1_t_bias", NULL, NULL,"nolace_tdshape3_alpha1_t_weights_float", NULL, NULL, NULL, 42, 80)) return 1;
    if (linear_init(&model->nolace_tdshape3_alpha2, arrays, "nolace_tdshape3_alpha2_bias", NULL, NULL,"nolace_tdshape3_alpha2_weights_float", NULL, NULL, NULL, 160, 80)) return 1;
    if (linear_init(&model->nolace_af2_kernel, arrays, "nolace_af2_kernel_bias", "nolace_af2_kernel_subias", "nolace_af2_kernel_weights_int8","nolace_af2_kernel_weights_float", NULL, NULL, "nolace_af2_kernel_scale", 160, 64)) return 1;
    if (linear_init(&model->nolace_af2_gain, arrays, "nolace_af2_gain_bias", NULL, NULL,"nolace_af2_gain_weights_float", NULL, NULL, NULL, 160, 2)) return 1;
    if (linear_init(&model->nolace_af3_kernel, arrays, "nolace_af3_kernel_bias", "nolace_af3_kernel_subias", "nolace_af3_kernel_weights_int8","nolace_af3_kernel_weights_float", NULL, NULL, "nolace_af3_kernel_scale", 160, 64)) return 1;
    if (linear_init(&model->nolace_af3_gain, arrays, "nolace_af3_gain_bias", NULL, NULL,"nolace_af3_gain_weights_float", NULL, NULL, NULL, 160, 2)) return 1;
    if (linear_init(&model->nolace_af4_kernel, arrays, "nolace_af4_kernel_bias", "nolace_af4_kernel_subias", "nolace_af4_kernel_weights_int8","nolace_af4_kernel_weights_float", NULL, NULL, "nolace_af4_kernel_scale", 160, 32)) return 1;
    if (linear_init(&model->nolace_af4_gain, arrays, "nolace_af4_gain_bias", NULL, NULL,"nolace_af4_gain_weights_float", NULL, NULL, NULL, 160, 1)) return 1;
    if (linear_init(&model->nolace_post_cf1, arrays, "nolace_post_cf1_bias", "nolace_post_cf1_subias", "nolace_post_cf1_weights_int8","nolace_post_cf1_weights_float", NULL, NULL, "nolace_post_cf1_scale", 320, 160)) return 1;
    if (linear_init(&model->nolace_post_cf2, arrays, "nolace_post_cf2_bias", "nolace_post_cf2_subias", "nolace_post_cf2_weights_int8","nolace_post_cf2_weights_float", NULL, NULL, "nolace_post_cf2_scale", 320, 160)) return 1;
    if (linear_init(&model->nolace_post_af1, arrays, "nolace_post_af1_bias", "nolace_post_af1_subias", "nolace_post_af1_weights_int8","nolace_post_af1_weights_float", NULL, NULL, "nolace_post_af1_scale", 320, 160)) return 1;
    if (linear_init(&model->nolace_post_af2, arrays, "nolace_post_af2_bias", "nolace_post_af2_subias", "nolace_post_af2_weights_int8","nolace_post_af2_weights_float", NULL, NULL, "nolace_post_af2_scale", 320, 160)) return 1;
    if (linear_init(&model->nolace_post_af3, arrays, "nolace_post_af3_bias", "nolace_post_af3_subias", "nolace_post_af3_weights_int8","nolace_post_af3_weights_float", NULL, NULL, "nolace_post_af3_scale", 320, 160)) return 1;
    return 0;
}
#endif /* DUMP_BINARY_WEIGHTS */
