package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "FileSharer_migrations",
		Usage: "./filesharer_migrations [version]",
		Action: func(c *cli.Context) error {
			switch c.Args().Len() {
			case 0:

				return nil
			case 1:
				return nil
			default:
				return errors.New("wrong number of args. See usage")
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
