# p2

[![CI](https://github.com/petesmithofficial/p2/actions/workflows/ci.yml/badge.svg)](https://github.com/petesmithofficial/p2/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

`p2` is a tiny Go CLI for printing powers of 2 through `2^32`.

## What it does

- `p2` prints the full list from `2^0` to `2^32`
- `p2 5` prints `5 (32)`
- `p2 30000` finds the closest supported power of 2 and prints `15 (32,768)`
- Exact midpoint ties return both matches, for example `p2 48` prints `5 (32), 6 (64)`

## Install

Install into `~/.local/bin`:

```sh
./install.sh
```

Install into a custom location:

```sh
BINDIR="$HOME/bin" ./install.sh
```

Build manually:

```sh
go build ./cmd/p2
```

## Usage

```sh
p2
p2 5
p2 30000
p2 48
```

## Development

```sh
go test ./...
go build ./cmd/p2
```

## License

This project is licensed under the [MIT License](LICENSE).
