package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Goalt/FileSharer/cmd/subcomands"
	_ "github.com/Goalt/FileSharer/cmd/subcomands/file_sharer_migrations"
	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/provider"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/context"
)

var (
	DebugLevel        = "DEBUG_LEVEL"
	MaxFileSize       = "MAX_FILE_SIZE"
	RootPath          = "ROOT_PATH"
	SecretKey         = "SECRET_KEY"
	MysqlDatabaseName = "MYSQL_DATABASE"
	MysqlUser         = "MYSQL_USER"
	MysqlPassword     = "MYSQL_PASSWORD"
	MysqlHost         = "MYSQL_HOST"
	MysqlPort         = "MYSQL_PORT"
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
			&cli.IntFlag{
				Name:    MaxFileSize,
				Value:   1,
				EnvVars: []string{MaxFileSize},
			},
			&cli.StringFlag{
				Name:    RootPath,
				EnvVars: []string{RootPath},
			},
			&cli.StringFlag{
				Name:    SecretKey,
				EnvVars: []string{SecretKey},
			},
			&cli.StringFlag{
				Name:    MysqlDatabaseName,
				EnvVars: []string{MysqlDatabaseName},
			},
			&cli.StringFlag{
				Name:    MysqlUser,
				EnvVars: []string{MysqlUser},
			},
			&cli.StringFlag{
				Name:    MysqlPassword,
				EnvVars: []string{MysqlPassword},
			},
			&cli.StringFlag{
				Name:    MysqlHost,
				EnvVars: []string{MysqlHost},
			},
			&cli.StringFlag{
				Name:    MysqlPort,
				EnvVars: []string{MysqlPort},
			},
		},
		Action: func(ctx *cli.Context) error {
			cfg := config.Config{
				MaxFileSize: ctx.Int(MaxFileSize),
				DebugLevel:  ctx.Int(DebugLevel),
				RootPath:    ctx.String(RootPath),
				Key:         []byte(ctx.String(SecretKey)),
				Database: config.Database{
					Host:     ctx.String(MysqlHost),
					Port:     ctx.String(MysqlPort),
					User:     ctx.String(MysqlUser),
					Password: ctx.String(MysqlPassword),
					DBName:   ctx.String(MysqlDatabaseName),
				},
				Server: config.Server{
					Port: 8080,
				},
			}

			fmt.Printf("%+v\n", cfg)

			signalCtx, cancel := context.WithCancel(context.Background())
			app, cleanup, err := provider.InitializeApp(cfg, signalCtx)
			defer cleanup()
			if err != nil {
				fmt.Println(err)
			}

			err = app.Run()
			if err != nil {
				fmt.Println(err)
			}

			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)

			<-c
			cancel()

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
