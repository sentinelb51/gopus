# gopus

gopus is a Go binding for the [Opus](http://www.opus-codec.org/) audio codec.

A fork of Tim Cooper's [layeh.com/gopus](https://github.com/layeh/gopus), which
vendors Opus 1.1.2 and was last touched in 2021. This one vendors 1.6.1, adds the
encoder CTLs an interactive client needs (in-band FEC, packet-loss estimate, DTX),
every neural extension bar DRED — Deep PLC, LACE, NoLACE and the blind bandwidth
extension — plus `Decoder.DecodeIn` and per-architecture vector builds. The module
path is this repository's own, so nothing importing it needs a `replace`.

The weights do not ship as C. Upstream emits 71 MB of float literals and fetches
them at `autogen.sh` time; here they are one 3.5 MB `opus_data.bin`, embedded and
handed to each decoder through `OPUS_SET_DNN_BLOB`. `./update-opus.sh` is what
regenerates both the tree and the blob, and pins the model by its own SHA-256.

## Documentation

- [API Reference](https://pkg.go.dev/github.com/sentinelb51/gopus)

## Requirements

- cgo
- [opus](http://www.opus-codec.org/) development library (only on platforms where the shared library is used)

## Build tags

By default libopus 1.6.1 is compiled from the vendored tree, so a fresh clone
needs no system library and no pkg-config.

| tag | effect |
| --- | --- |
| `opus_shared` | link the system libopus instead of the vendored one. Needs the development library; one older than 1.5 has no Deep PLC and one older than 1.6, or built without `--enable-osce`, has no LACE, NoLACE or bandwidth extension. |
| `opus_baseline` | drop the `-march=x86-64-v3` floor on amd64. |

`-march=x86-64-v3` means AVX2, FMA and BMI: Haswell and Excavator (2013) and
newer. It is on by default because the neural models choose their vector width
from the compiler's own `__AVX2__`, and without it a concealed frame costs about
three times as much. On a machine older than that the failure is SIGILL rather
than a slow decode, and `-tags opus_baseline` is the answer.

Compiling upstream's own SSE4.1 and AVX2 sources for celt and silk on top of that
is worth nothing and is not done — measured at 130.7 µs to encode a 20 ms frame
either way. `opus-1.6.1/config.h` has the detail.

`-fno-math-errno -fno-trapping-math` are worth a further ~6% and are safe here,
but cgo's flag allowlist will not let a package carry them. Pass them through
`CGO_CFLAGS` if you want them. `-Ofast` is refused by libopus itself, and `-flto`
gains nothing — the vendored build is already a single translation unit.

## License

Public domain

## Author

Tim Cooper (<tim.cooper@layeh.com>)
