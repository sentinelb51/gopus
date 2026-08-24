/* Hand-written for the vendored build. libopus normally generates this with
   autoconf; nothing here runs configure, so this states the same answers
   directly.

   The build is float. Deep PLC is compiled in; the define for it arrives on the
   command line from opus_nonshared.go, next to the dnn/ sources it needs. DRED
   and OSCE are the extensions left out, and their 10 MB of weights with them. */

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

/* Vector instructions, decided here at compile time and never at run time.
   OPUS_HAVE_RTCD stays undefined, so libopus's PRESUME_* branches turn every
   dispatch into a direct call: no function-pointer table, no cpuid, and nothing
   to pay on a call that was already going to happen.

   What may be presumed is what the architecture itself guarantees — SSE and
   SSE2 on amd64, Neon on arm64 — so this raises no CPU floor and needs no
   -march. SSE4.1 and AVX2 are deliberately absent: neither is guaranteed, and
   reaching them means either a -march that drops every machine before 2013 or
   the runtime dispatch this build exists to avoid.

   The sources these resolve to are compiled by simd_amd64.c and simd_arm64.c,
   whose GOARCH filename suffix is what keeps one architecture's intrinsics out
   of the other's build. 386 is not here on purpose: SSE2 is not in its
   baseline, and Go's own GO386=sse2 default says nothing about what the C
   compiler may emit. */
/* #undef OPUS_HAVE_RTCD */
#if defined(__x86_64__) || defined(_M_X64)
# define OPUS_X86_MAY_HAVE_SSE 1
# define OPUS_X86_PRESUME_SSE 1
# define OPUS_X86_MAY_HAVE_SSE2 1
# define OPUS_X86_PRESUME_SSE2 1
#elif defined(__aarch64__) || defined(_M_ARM64)
# define OPUS_ARM_MAY_HAVE_NEON_INTR 1
# define OPUS_ARM_PRESUME_NEON_INTR 1
#endif
/* #undef OPUS_X86_MAY_HAVE_SSE4_1 */
/* #undef OPUS_X86_MAY_HAVE_AVX2 */
/* #undef OPUS_ARM_MAY_HAVE_DOTPROD */

/* DRED and OSCE off; Deep PLC arrives from the command line. */
/* #undef ENABLE_DRED */
/* #undef ENABLE_OSCE */

#define restrict __restrict
