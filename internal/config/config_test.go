package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPathFromUserConfigDir(t *testing.T) {
	t.Parallel()

	got := PathFromUserConfigDir("/tmp/config-root")
	want := filepath.Join("/tmp/config-root", "p2", "config.json")
	if got != want {
		t.Fatalf("PathFromUserConfigDir() = %q, want %q", got, want)
	}
}

func TestLoadFromPathMissingFileUsesDefaults(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "missing.json")

	got, err := LoadFromPath(path)
	if err != nil {
		t.Fatalf("LoadFromPath() error = %v, want nil", err)
	}

	if got != Default() {
		t.Fatalf("LoadFromPath() = %#v, want %#v", got, Default())
	}
}

func TestLoadFromPathPartialConfigMergesDefaults(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "config.json")
	data := []byte(`{"lower_bound":5,"upper_bound":8,"use_commas":false}`)
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	got, err := LoadFromPath(path)
	if err != nil {
		t.Fatalf("LoadFromPath() error = %v, want nil", err)
	}

	want := Default()
	want.LowerBound = 5
	want.UpperBound = 8
	want.UseCommas = false

	if got != want {
		t.Fatalf("LoadFromPath() = %#v, want %#v", got, want)
	}
}

func TestLoadFromPathUnknownFieldsAreIgnored(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "config.json")
	data := []byte(`{"mystery":"value","use_commas":false}`)
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	got, err := LoadFromPath(path)
	if err != nil {
		t.Fatalf("LoadFromPath() error = %v, want nil", err)
	}

	want := Default()
	want.UseCommas = false

	if got != want {
		t.Fatalf("LoadFromPath() = %#v, want %#v", got, want)
	}
}

func TestLoadFromPathInvalidJSON(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(path, []byte(`{"lower_bound":`), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	_, err := LoadFromPath(path)
	if err == nil {
		t.Fatal("LoadFromPath() error = nil, want non-nil")
	}
}

func TestLoadFromPathInvalidBounds(t *testing.T) {
	t.Parallel()

	testCases := []string{
		`{"lower_bound":-1}`,
		`{"upper_bound":33}`,
		`{"lower_bound":8,"upper_bound":5}`,
	}

	for _, input := range testCases {
		input := input
		t.Run(input, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.json")
			if err := os.WriteFile(path, []byte(input), 0o644); err != nil {
				t.Fatalf("WriteFile() error = %v", err)
			}

			_, err := LoadFromPath(path)
			if err == nil {
				t.Fatal("LoadFromPath() error = nil, want non-nil")
			}
		})
	}
}

func TestSaveToPath(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "nested", "config.json")
	cfg := Default()
	cfg.LowerBound = 5
	cfg.UpperBound = 8
	cfg.UseCommas = false

	if err := SaveToPath(path, cfg); err != nil {
		t.Fatalf("SaveToPath() error = %v, want nil", err)
	}

	got, err := LoadFromPath(path)
	if err != nil {
		t.Fatalf("LoadFromPath() error = %v, want nil", err)
	}

	if got != cfg {
		t.Fatalf("LoadFromPath() = %#v, want %#v", got, cfg)
	}
}

func TestSaveToPathInvalidConfig(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "config.json")
	cfg := Default()
	cfg.LowerBound = 9
	cfg.UpperBound = 2

	if err := SaveToPath(path, cfg); err == nil {
		t.Fatal("SaveToPath() error = nil, want non-nil")
	}
}
