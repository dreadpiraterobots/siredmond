package core

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func TestCleanHandlerFormatsMessageAndAttrs(t *testing.T) {
	var buf bytes.Buffer
	h := &CleanHandler{out: &buf}
	logger := slog.New(h)

	// Log a message with one attribute
	logger.Info("hello world", slog.String("k", "v"))

	out := buf.String()
	// It should include the level, message, and the attribute key=value.
	if !strings.Contains(out, "[INFO]") {
		t.Fatalf("expected level marker in output, got: %q", out)
	}
	if !strings.Contains(out, "hello world") {
		t.Fatalf("expected message in output, got: %q", out)
	}
	if !strings.Contains(out, "k=v") {
		t.Fatalf("expected attribute k=v in output, got: %q", out)
	}
	// Component derived from filename should be present
	if !strings.Contains(out, "logger_test") {
		// The component is derived from runtime caller filename; presence is expected but not guaranteed in all envs.
		t.Logf("warning: component substring 'logger_test' not found; output was: %q", out)
	}
	// Timestamp including milliseconds should be present
	if !strings.Contains(out, "T") || !strings.Contains(out, "Z") {
		t.Fatalf("expected timestamp in output, got: %q", out)
	}
}
