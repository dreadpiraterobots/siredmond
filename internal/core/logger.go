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

func (h *CleanHandler) Handle(_ context.Context, logRecord slog.Record) error {
	// Time in UTC
	// The CVRF uses UTC (very sensible!) so we will follow suit
	ts := logRecord.Time.UTC().Format("2006-01-02T15:04:05.000Z")

	// Derive component name from filename
	// logRecord.PC (Program Counter) to find the file that called the logger
	component := "unknown"
	if logRecord.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{logRecord.PC})
		f, _ := fs.Next()
		if f.File != "" {
			component = strings.TrimSuffix(filepath.Base(f.File), ".go")
		}
	}

	// Collect attributes into a single space-separated string: key=value
	// This is useful when we want to log additional context in a way that is easy to parse
	var parts []string
	logRecord.Attrs(func(a slog.Attr) bool {
		parts = append(parts, fmt.Sprintf("%s=%s", a.Key, a.Value.String()))
		return true
	})

	attrs := ""
	if len(parts) > 0 {
		attrs = " " + strings.Join(parts, " ")
	}

	// Thread-safe write to STDERR; should be atomic for any reasonably sized log entry
	_, err := fmt.Fprintf(h.out, "%s %s [%s] %s%s\n",
		ts, component, logRecord.Level.String(), logRecord.Message, attrs)
	return err
}

// These are required to satisfy the slog.Handler interface
func (h *CleanHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }
func (h *CleanHandler) WithGroup(name string) slog.Handler       { return h }
