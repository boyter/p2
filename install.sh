#!/usr/bin/env sh
set -eu

if ! command -v go >/dev/null 2>&1; then
	echo "error: go is required but was not found in PATH" >&2
	exit 1
fi

if ! command -v install >/dev/null 2>&1; then
	echo "error: install is required but was not found in PATH" >&2
	exit 1
fi

: "${HOME:?error: HOME must be set}"

PREFIX=${PREFIX:-"$HOME/.local"}
BINDIR=${BINDIR:-"$PREFIX/bin"}

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
TMPDIR_ROOT=${TMPDIR:-/tmp}
BUILD_DIR=$(mktemp -d "$TMPDIR_ROOT/p2.XXXXXX")
cleanup() {
	rm -rf "$BUILD_DIR"
}
trap cleanup EXIT INT TERM HUP

mkdir -p "$BINDIR"

cd "$SCRIPT_DIR"
go build -o "$BUILD_DIR/p2" ./cmd/p2
install -m 0755 "$BUILD_DIR/p2" "$BINDIR/p2"

echo "installed p2 to $BINDIR/p2"

case ":${PATH:-}:" in
	*:"$BINDIR":*)
		;;
	*)
		echo "add $BINDIR to PATH to run p2 from any shell" >&2
		;;
esac
