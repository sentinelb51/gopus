/* Hand-written for the vendored build. libopus normally generates this with
   autoconf; nothing here runs configure, so this states the same answers
   directly.

   The build is float, pure C, and without the neural extensions: ENABLE_DEEP_PLC,
   ENABLE_DRED and ENABLE_OSCE are deliberately left undefined, which is what
   keeps the 11 MB dnn/ tree out of the module. */

#define OPUS_BUILD 1
#define PACKAGE_NAME "opus"
#define PACKAGE_VERSION "1.5.2"
#define PACKAGE_STRING "opus 1.5.2"

/* Float build: FIXED_POINT stays undefined. */

/* C99 variable-length arrays rather than alloca, which needs a header that is
   not portable across the platforms this module is built on. */
#define VAR_ARRAYS 1

#define HAVE_LRINT 1
#define HAVE_LRINTF 1
#define HAVE_STDINT_H 1
#define HAVE_STRING_H 1
#define HAVE_STDLIB_H 1
#define STDC_HEADERS 1

/* No run-time CPU detection and no intrinsics: only the generic C paths are
   compiled, so the same object works on every amd64 and arm64 target without a
   dispatch table. Encode is about a millisecond a frame either way at 48 kHz
   mono, which is what makes this trade cheap. */
/* #undef OPUS_HAVE_RTCD */
/* #undef OPUS_X86_MAY_HAVE_SSE */
/* #undef OPUS_X86_MAY_HAVE_SSE2 */
/* #undef OPUS_X86_MAY_HAVE_SSE4_1 */
/* #undef OPUS_X86_MAY_HAVE_AVX2 */
/* #undef OPUS_ARM_MAY_HAVE_NEON_INTR */

/* The neural extensions, all off. */
/* #undef ENABLE_DEEP_PLC */
/* #undef ENABLE_DRED */
/* #undef ENABLE_OSCE */

#define restrict __restrict
