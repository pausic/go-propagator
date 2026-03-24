package config

import (
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Setenv("CONFIG_PATH", "testdata/valid.yaml")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected config, got nil")
	}
}

func TestLoadErrors(t *testing.T) {
	tests := []struct {
		name       string
		configPath string
		wantErr    string
	}{
		{"default path not found", "", "reading config file"},
		{"file not found", "/nonexistent/config.yaml", "reading config file"},
		{"invalid yaml", "testdata/invalid.yaml", "parsing config file"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("CONFIG_PATH", tt.configPath)

			_, err := Load()
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("expected %q in error, got: %v", tt.wantErr, err)
			}
		})
	}
}
