package main

import (
	"fmt"
	"os"

	"github.com/Goalt/FileSharer/cmd/subcomands"
	_ "github.com/Goalt/FileSharer/cmd/subcomands/file_sharer_migrations"
	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/provider"
	"github.com/sethvargo/go-signalcontext"
	"github.com/urfave/cli/v2"
)

var (
	DebugLevel = "DEBUG_LEVEL"
)

func main() {
	app := &cli.App{
		Name:     "FileSharer",
		Usage:    "./filesharer",
		HelpName: "Web Server for filesharer app",
		Commands: subcomands.Get(),
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    DebugLevel,
				Value:   1,
				EnvVars: []string{DebugLevel},
			},
		},
		Action: func(ctx *cli.Context) error {
			cfg := config.Config{
				DebugLevel: ctx.Int(DebugLevel),
				Server: config.Server{
					Port: 8080,
				},
			}

			signalCtx, cancel := signalcontext.OnInterrupt()
			defer cancel()

			app := provider.InitializeApp(cfg, signalCtx)

			_ = app.Run()

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
