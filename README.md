# gopus

gopus is a Go binding for the [Opus](http://www.opus-codec.org/) audio codec.

## Documentation

- [API Reference](https://godoc.org/layeh.com/gopus)

## Requirements

- cgo
- [opus](http://www.opus-codec.org/) development library (only on platforms where the shared library is used)

## Build tags

By default libopus 1.5.2 is compiled from the vendored tree, so a fresh clone
needs no system library and no pkg-config.

| tag | effect |
| --- | --- |
| `opus_shared` | link the system libopus instead of the vendored one. Needs the development library, and a system libopus older than 1.5 has no Deep PLC. |
| `opus_baseline` | drop the `-march=x86-64-v3` floor on amd64. |

`-march=x86-64-v3` means AVX2, FMA and BMI: Haswell and Excavator (2013) and
newer. It is on by default because libopus 1.5's Deep PLC chooses its vector
width from the compiler's own `__AVX2__`, and without it a concealed frame costs
about three times as much. On a machine older than that the failure is SIGILL
rather than a slow decode, and `-tags opus_baseline` is the answer.

`-fno-math-errno -fno-trapping-math` are worth a further ~6% and are safe here,
but cgo's flag allowlist will not let a package carry them. Pass them through
`CGO_CFLAGS` if you want them. `-Ofast` is refused by libopus itself, and `-flto`
gains nothing — the vendored build is already a single translation unit.

## License

Public domain

## Author

Tim Cooper (<tim.cooper@layeh.com>)
