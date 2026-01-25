package core

import (
	"context"
	"reflect"
	"strings"
	"testing"
)

// Ensure NewEngine constructs a usable Engine with a logger.
func TestNewEngineInitializesLogger(t *testing.T) {
	e := NewEngine()
	if e == nil {
		t.Fatal("NewEngine returned nil")
	}
	if e.logger == nil {
		t.Fatal("NewEngine did not initialize logger")
	}
}

// Basic check that the logger's handler is the custom CleanHandler (by type name).
// This avoids depending on internals of slog while still validating we wired our handler.
func TestNewEngineUsesCleanHandler(t *testing.T) {
	e := NewEngine()
	h := e.logger.Handler()
	if h == nil {
		t.Fatal("logger.Handler() returned nil")
	}
	typ := reflect.TypeOf(h).String()
	if !strings.Contains(typ, "CleanHandler") {
		t.Fatalf("expected handler type to mention CleanHandler, got: %s", typ)
	}
}

// Smoke test: DownloadCVRF should respect the provided context.
// We supply an already-cancelled context and expect an error quickly.
// This avoids depending on network responses while asserting the method observes context.
func TestDownloadCVRFRespectsContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately. Your Uber driver will not be giving you 5 stars today.

	e := NewEngine()
	if err := e.DownloadCVRF(ctx); err == nil {
		t.Fatal("expected DownloadCVRF to return an error when context is cancelled")
	}
}
