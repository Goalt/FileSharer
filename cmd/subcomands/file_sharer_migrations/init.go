package file_sharer_migrations

import (
	"errors"
	"github.com/Goalt/FileSharer/cmd/subcomands"
	"github.com/urfave/cli/v2"
)

func init() {
	subcommand := &cli.Command{
		Name:  "FileSharer migrations",
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

	subcomands.Add(subcommand)
}
