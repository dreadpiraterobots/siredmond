package core

import (
	"log/slog"
	"os"
)

type Engine struct {
	logger *slog.Logger
}

// The fat controller
func NewEngine() *Engine {
	// Set up logging as defined in logger.go
	myHandler := &CleanHandler{out: os.Stderr}
	return &Engine{
		logger: slog.New(myHandler),
	}
}
