package main

import (
	"fmt"
	"os"

	"github.com/Goalt/FileSharer/cmd/subcomands"
	_ "github.com/Goalt/FileSharer/cmd/subcomands/file_sharer_migrations"
	"github.com/Goalt/FileSharer/cmd/variables"
	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/provider"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/context"
)

func main() {
	app := &cli.App{
		Name:     "FileSharer",
		Usage:    "./filesharer",
		HelpName: "Web Server for filesharer app",
		Commands: subcomands.Get(),
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    variables.DebugLevel,
				Value:   1,
				EnvVars: []string{variables.DebugLevel},
			},
			&cli.IntFlag{
				Name:    variables.MaxFileSize,
				Value:   1,
				EnvVars: []string{variables.MaxFileSize},
			},
			&cli.StringFlag{
				Name:    variables.RootPath,
				EnvVars: []string{variables.RootPath},
			},
			&cli.StringFlag{
				Name:    variables.SecretKey,
				EnvVars: []string{variables.SecretKey},
			},
			&cli.StringFlag{
				Name:    variables.MysqlDatabaseName,
				EnvVars: []string{variables.MysqlDatabaseName},
			},
			&cli.StringFlag{
				Name:    variables.MysqlUser,
				EnvVars: []string{variables.MysqlUser},
			},
			&cli.StringFlag{
				Name:    variables.MysqlPassword,
				EnvVars: []string{variables.MysqlPassword},
			},
			&cli.StringFlag{
				Name:    variables.MysqlHost,
				EnvVars: []string{variables.MysqlHost},
			},
			&cli.StringFlag{
				Name:    variables.MysqlPort,
				EnvVars: []string{variables.MysqlPort},
			},
		},
		Action: func(ctx *cli.Context) error {
			cfg := config.Config{
				MaxFileSize: ctx.Int(variables.MaxFileSize),
				RootPath:    ctx.String(variables.RootPath),
				Key:         []byte(ctx.String(variables.SecretKey)),
				Logger: config.Logger{
					SetReportCaller: true,
					Level:           config.InfoLevel,
				},
				Database: config.Database{
					Host:     ctx.String(variables.MysqlHost),
					Port:     ctx.String(variables.MysqlPort),
					User:     ctx.String(variables.MysqlUser),
					Password: ctx.String(variables.MysqlPassword),
					DBName:   ctx.String(variables.MysqlDatabaseName),
				},
				Server: config.Server{
					Port: 33333,
				},
			}
			fmt.Printf("%+v\n", cfg)
			fmt.Println("BUILDKIT_HOST:", os.Getenv("BUILDKIT_HOST"))
			fmt.Println("OKTETO_REGISTRY_URL", os.Getenv("OKTETO_REGISTRY_URL"))

			signalCtx, _ := context.WithCancel(context.Background())
			app, cleanup, err := provider.InitializeApp(cfg, signalCtx)
			if cleanup != nil {
				defer cleanup()
			}
			if err != nil {
				fmt.Println(err)
			}

			err = app.Run()
			if err != nil {
				fmt.Println(err)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
