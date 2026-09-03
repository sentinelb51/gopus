/* Hand-written for the vendored build, and copied into opus-<version>/config.h
   by update-opus.sh. libopus normally generates this with autoconf; nothing
   here runs configure, so this states the same answers directly.

   The build is float, with Deep PLC and OSCE — the latter bringing LACE,
   NoLACE and the blind bandwidth extension new in 1.6. DRED is the one
   extension left out: it needs the *sender* to enable it too, so it buys
   nothing against a far end that is not also this library. */

#define OPUS_BUILD 1
#define PACKAGE_NAME "opus"
#define PACKAGE_VERSION "1.6.1"
#define PACKAGE_STRING "opus 1.6.1"

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
   SSE2 on amd64, Neon on arm64 — so nothing in this file raises a CPU floor or
   needs a -march.

   SSE4.1 and AVX2 stay out, and not for the reason it looks like. march_amd64.go
   does give the whole package AVX2, so upstream's own SSE4.1 and AVX2 sources
   *could* be compiled and presumed here — that was tried, all ten of them, and
   it is worth **nothing**: 130.7 µs to encode a 20 ms frame either way, and
   decode inside the noise. The hand-written intrinsics were written for a
   compiler that had been told nothing about the target, and once it has been,
   gcc reaches the same width from the plain C. Ten more vendored files and a
   second dispatch level for zero.

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

/* The neural extensions. ENABLE_OSCE also turns on ENABLE_DEEP_PLC upstream;
   both are stated because this file is read rather than executed. */
#define ENABLE_DEEP_PLC 1
#define ENABLE_OSCE 1
#define ENABLE_OSCE_BWE 1
/* #undef ENABLE_DRED */

/* The weights are not compiled in: they arrive as opus_data.bin through
   OPUS_SET_DNN_BLOB, which is what keeps 71 MB of C literals out of a module
   every importer downloads. See update-opus.sh. */
#define USE_WEIGHTS_FILE 1

/* Upstream's own configure default, and what halves the blob: the float
   duplicate of every int8 array is for debugging the quantisation. */
#define DISABLE_DEBUG_FLOAT 1

#define restrict __restrict
