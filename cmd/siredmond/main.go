package main

import (
	"github.com/dreadpiraterobots/siredmond/internal/core"
	"github.com/dreadpiraterobots/siredmond/internal/ui"
	"log"
	"os"
)

func main() {
	engine := core.NewEngine()
	app := ui.NewCLI(engine)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
