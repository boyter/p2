package main

import (
	"bytes"
	"testing"

	"p2/internal/powers"
)

func TestRunNoArgs(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run(nil, &stdout, &stderr)

	if exitCode != 0 {
		t.Fatalf("run() exit code = %d, want 0", exitCode)
	}

	want := powers.FormatEntries(powers.All()) + "\n"
	if stdout.String() != want {
		t.Fatalf("stdout = %q, want %q", stdout.String(), want)
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunExponentArg(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"5"}, &stdout, &stderr)

	if exitCode != 0 {
		t.Fatalf("run() exit code = %d, want 0", exitCode)
	}

	if stdout.String() != "5 (32)\n" {
		t.Fatalf("stdout = %q, want %q", stdout.String(), "5 (32)\n")
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunClosestArg(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"30000"}, &stdout, &stderr)

	if exitCode != 0 {
		t.Fatalf("run() exit code = %d, want 0", exitCode)
	}

	if stdout.String() != "15 (32,768)\n" {
		t.Fatalf("stdout = %q, want %q", stdout.String(), "15 (32,768)\n")
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunTieArg(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"48"}, &stdout, &stderr)

	if exitCode != 0 {
		t.Fatalf("run() exit code = %d, want 0", exitCode)
	}

	if stdout.String() != "5 (32), 6 (64)\n" {
		t.Fatalf("stdout = %q, want %q", stdout.String(), "5 (32), 6 (64)\n")
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunInvalidInput(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		args []string
	}{
		{name: "negative", args: []string{"-1"}},
		{name: "decimal", args: []string{"5.5"}},
		{name: "nonnumeric", args: []string{"hello"}},
		{name: "extra args", args: []string{"1", "2"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			exitCode := run(tc.args, &stdout, &stderr)

			if exitCode != 2 {
				t.Fatalf("run(%v) exit code = %d, want 2", tc.args, exitCode)
			}

			if stdout.Len() != 0 {
				t.Fatalf("stdout = %q, want empty", stdout.String())
			}

			if stderr.Len() == 0 {
				t.Fatal("stderr is empty, want usage text")
			}
		})
	}
}
