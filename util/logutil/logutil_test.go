package logutil_test

import (
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/util/logutil"
)

func TestNewLogger_production(t *testing.T) {
	logger := logutil.NewLogger("production")

	if logger == nil {
		t.Error("expected logger, got nil")
	}
}

func TestNewLogger_development(t *testing.T) {
	logger := logutil.NewLogger("development")

	if logger == nil {
		t.Error("expected logger, got nil")
	}
}
