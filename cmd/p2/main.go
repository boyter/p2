package main

import (
	"fmt"
	"io"
	"math/big"
	"os"

	"p2/internal/powers"
)

const usageText = `usage: p2 [integer]

With no arguments, p2 prints powers of 2 from 2^0 through 2^32.
With one integer argument:
  0..32  treat the value as an exponent
  >32    treat the value as a target and print the closest power of 2
`

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout io.Writer, stderr io.Writer) int {
	switch len(args) {
	case 0:
		_, _ = fmt.Fprintln(stdout, powers.FormatEntries(powers.All()))
		return 0
	case 1:
	default:
		return usageError(stderr, "expected zero or one argument")
	}

	input, ok := new(big.Int).SetString(args[0], 10)
	if !ok {
		return usageError(stderr, fmt.Sprintf("invalid integer %q", args[0]))
	}

	if input.Sign() < 0 {
		return usageError(stderr, "negative integers are not supported")
	}

	if input.Cmp(big.NewInt(powers.MaxExponent)) <= 0 {
		entry, found := powers.ByExponent(uint(input.Uint64()))
		if !found {
			return usageError(stderr, fmt.Sprintf("unsupported exponent %q", args[0]))
		}

		_, _ = fmt.Fprintln(stdout, powers.FormatEntries([]powers.Entry{entry}))
		return 0
	}

	_, _ = fmt.Fprintln(stdout, powers.FormatEntries(powers.ClosestTo(input)))
	return 0
}

func usageError(stderr io.Writer, message string) int {
	_, _ = fmt.Fprintf(stderr, "error: %s\n\n%s", message, usageText)
	return 2
}
