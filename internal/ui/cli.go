/* Define the CLI application and its commands.
We also initialize bits of the core Engine as we need them. */

package ui

import (
	"github.com/dreadpiraterobots/siredmond/internal/core"
	"github.com/urfave/cli/v2"
)

func NewCLI(engine *core.Engine) *cli.App {
	return &cli.App{
		Name:  "siredmond",
		Usage: "Microsoft security advisory analysis and exploration",
		Commands: []*cli.Command{
			{
				Name:  "download",
				Usage: "Download security advisory resources",
				Subcommands: []*cli.Command{
					{
						Name:  "cvrf",
						Usage: "Download CVRF data",
						Action: func(c *cli.Context) error {
							// 1. Initialize the Engine
							engine := core.NewEngine()

							// 2. Call the method in downloader.go
							return engine.DownloadCVRF()
						},
					},
				},
			},
		},
	}
}
