#!/bin/sh
# Re-vendor opus-<version>/ from xiph/opus.
#
# Upstream is autotools and downloads its neural model data at ./autogen.sh
# time; this module is cgo and a fresh `go get` must build with nothing fetched.
# The gap is closed here rather than at build time: every model is written into
# one binary blob, opus_data.bin, which is committed and go:embed'd and handed
# to each decoder through OPUS_SET_DNN_BLOB. The *_data.c files that would
# otherwise carry those weights as C literals — 71 MB of them, for a module
# every importer downloads — are reduced to their init functions.
#
# Needs: git, curl, sha256sum, tar and a C compiler. Nothing here runs during a
# build.
set -e

VERSION=1.6.1
TAG=v$VERSION
# Upstream names the model archive by its own SHA-256, so the pin and the
# checksum are one string. It is the argument autogen.sh passes to
# dnn/download_model.sh; take it from there when moving version.
MODEL=a5177ec6fb7d15058e99e57029746100121f68e4890b1467d4094aa336b6013e

root=$(cd "$(dirname "$0")" && pwd)
dst=$root/opus-$VERSION
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT

echo "cloning xiph/opus $TAG"
git -C "$work" clone --quiet --depth 1 --branch "$TAG" https://github.com/xiph/opus.git src

echo "fetching model $MODEL"
curl -sSLo "$work/opus_data.tar.gz" \
	"https://media.xiph.org/opus/models/opus_data-$MODEL.tar.gz"
echo "$MODEL  $work/opus_data.tar.gz" | sha256sum -c -
tar xzf "$work/opus_data.tar.gz" -C "$work/src"

# The weights, as the bytes OPUS_SET_DNN_BLOB parses. Two changes to upstream's
# dumper, both ours:
#
#   - "wb", because it opens the blob in text mode and Windows corrupts it.
#   - bbwenet, which upstream's dumper does not write. Blind BWE is new in 1.6
#     and osce_load_models *does* read it back, so without this the blob is
#     short one model, init_bbwenet fails, and OSCE turns itself off entirely
#     rather than losing just the bandwidth extension.
#   - DRED out, which upstream's dumper writes unconditionally. This build does
#     not enable it, so those are 690 KB nothing ever reads.
#
# DISABLE_DEBUG_FLOAT is upstream's own configure default and is what drops the
# float duplicate of every int8 array: 3.5 MB rather than 13.4.
echo "building the weight blob"
sed -i 's/fopen("weights_blob.bin", "w")/fopen("weights_blob.bin", "wb")/' \
	"$work/src/dnn/write_lpcnet_weights.c"
sed -i 's|^#include "nolace_data.c"$|#include "nolace_data.c"\n#include "bbwenet_data.c"|' \
	"$work/src/dnn/write_lpcnet_weights.c"
sed -i 's|^  write_weights(nolacelayers_arrays, fout);$|  write_weights(nolacelayers_arrays, fout);\n  write_weights(bbwenetlayers_arrays, fout);|' \
	"$work/src/dnn/write_lpcnet_weights.c"
sed -i '/^#include "dred_rdovae_\(enc\|dec\)_data\.c"$/d;
	/^  write_weights(rdovae\(enc\|dec\)_arrays, fout);$/d' \
	"$work/src/dnn/write_lpcnet_weights.c"
grep -q bbwenetlayers_arrays "$work/src/dnn/write_lpcnet_weights.c" &&
	! grep -q rdovae "$work/src/dnn/write_lpcnet_weights.c" ||
	{ echo "the dumper patches did not apply"; exit 1; }

cat >"$work/src/config.h" <<'EOF'
#define OPUS_BUILD 1
#define PACKAGE_VERSION "blob"
#define VAR_ARRAYS 1
#define ENABLE_DEEP_PLC 1
#define ENABLE_OSCE 1
#define ENABLE_OSCE_BWE 1
#define DISABLE_DEBUG_FLOAT 1
EOF
cc=${CC:-$(command -v cc || command -v gcc)}
(cd "$work/src" && "$cc" -O2 -DHAVE_CONFIG_H -DDUMP_BINARY_WEIGHTS -I. -Iinclude \
	-Icelt -Isilk -Isilk/float -Idnn dnn/write_lpcnet_weights.c \
	-o dump_weights_blob -lm && ./dump_weights_blob)

echo "vendoring"
rm -rf "$dst"
mkdir -p "$dst"
cp "$work/src/AUTHORS" "$work/src/COPYING" "$dst/"

# Every header, because opus_decoder.c reaches DRED's and the fixed-point
# tree's whether or not their .c files are compiled; only the .c files this
# module builds, which is the float, no-DRED, no-lossgen configuration plus the
# per-architecture files simd_*.c and dnn_*.c name.
(cd "$work/src" && find celt silk src dnn include -name '*.h' -not -path '*/tests/*') |
	while read -r p; do
		mkdir -p "$dst/$(dirname "$p")"
		cp "$work/src/$p" "$dst/$p"
	done

sources=$(cd "$work/src" && cat celt_sources.mk silk_sources.mk opus_sources.mk \
	lpcnet_sources.mk | grep -oE '(celt|silk|src|dnn)/[A-Za-z0-9_/]+\.c' |
	grep -vE '^(celt|silk|dnn)/(x86|arm)/|^silk/fixed/|dred|lossgen' | sort -u)
arch="celt/x86/pitch_sse.c celt/x86/vq_sse2.c dnn/x86/nnet_sse2.c
	celt/arm/celt_neon_intr.c celt/arm/pitch_neon_intr.c
	silk/arm/biquad_alt_neon_intr.c dnn/arm/nnet_neon.c"
for p in $sources $arch; do
	mkdir -p "$dst/$(dirname "$p")"
	cp "$work/src/$p" "$dst/$p"
done

cp "$work/src/weights_blob.bin" "$root/opus_data.bin"

# Everything before init_* in a *_data.c is the weights as C literals, which
# USE_WEIGHTS_FILE excludes anyway. Dropping the text keeps 71 MB out of the
# module and off every compile.
for p in "$dst"/dnn/*_data.c; do
	name=$(basename "$p")
	{
		printf '/* gopus: upstream'\''s %s reduced to what USE_WEIGHTS_FILE\n' "$name"
		printf '   compiles. The weights it would otherwise carry are in\n'
		printf '   opus_data.bin. Regenerate with update-opus.sh. */\n\n'
		printf '#include "config.h"\n#include "%s"\n#include "nnet.h"\n\n' \
			"${name%.c}.h"
		sed -n '/^#ifndef DUMP_BINARY_WEIGHTS/,$p' "$p"
	} >"$p.trimmed"
	mv "$p.trimmed" "$p"
done

# The hand-written config.h: nothing here runs configure. Kept in the tree
# rather than written above, because it is the one file in opus-*/ that is ours.
cp "$root/opus_config.h" "$dst/config.h"

# What is committed stays LF, and core.autocrlf=true will have converted every
# one of these on the way out of upstream's checkout.
find "$dst" -name '*.c' -o -name '*.h' | while read -r p; do sed -i 's/\r$//' "$p"; done

echo "done — opus $TAG, model $MODEL, blob $(stat -c %s "$root/opus_data.bin") bytes"
