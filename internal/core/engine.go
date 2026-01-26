// https://go.dev/blog/package-names
// "Good package names are short and clear.
// They are lower case, with no under_scores or mixedCaps. They are often simple nouns.
package core

import (
	"log/slog"
	"net/http"
	"os"
)

type Engine struct {
	logger     *slog.Logger
	httpClient *http.Client
}

// The fat controller
func NewEngine() *Engine {
	// Set up logging as defined in logger.go
	myHandler := &CleanHandler{out: os.Stderr}
	return &Engine{
		logger:     slog.New(myHandler),
		httpClient: &http.Client{Timeout: httpClientTimeout},
	}
}
