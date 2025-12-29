package core

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"
)

type CleanHandler struct {
	out io.Writer
}

func (h *CleanHandler) Enabled(_ context.Context, _ slog.Level) bool { return true }

func (h *CleanHandler) Handle(_ context.Context, r slog.Record) error {
	// Time in UTC
	// The CVRF uses UTC, so we will follow suit
	t := r.Time.UTC().Format("2006-01-02T15:04:05Z")

	// Derive component name from filename
	// We use r.PC (Program Counter) to find the file that called the logger
	component := "unknown"
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		if f.File != "" {
			component = strings.TrimSuffix(filepath.Base(f.File), ".go")
		}
	}

	// Collect attributes into a single space-separated string: key=value
	// This is useful when we want to log additional context in a way that is easy to parse
	var parts []string
	r.Attrs(func(a slog.Attr) bool {
		parts = append(parts, fmt.Sprintf("%s=%s", a.Key, a.Value.String()))
		return true
	})

	attrs := ""
	if len(parts) > 0 {
		attrs = " " + strings.Join(parts, " ")
	}

	// Thread-safe write to STDERR; should be atomic for any reasonably sized log entry
	_, err := fmt.Fprintf(h.out, "%s %s [%s] %s%s\n",
		t, component, r.Level.String(), r.Message, attrs)
	return err
}

// These are required to satisfy the slog.Handler interface
func (h *CleanHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }
func (h *CleanHandler) WithGroup(name string) slog.Handler       { return h }
