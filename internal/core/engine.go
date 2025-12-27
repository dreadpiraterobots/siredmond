package core

import (
	"fmt"
)

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) DownloadCVRF() error {
	// We'll implement the actual http logic next,
	// for now, let's just make sure it wires up.
	fmt.Printf("Engine: Fetching data...\n")
	return nil
}
