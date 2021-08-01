package subcomands

import "github.com/urfave/cli/v2"

var subcommandSingleton []*cli.Command

func Add(command *cli.Command) {
	if subcommandSingleton == nil {
		subcommandSingleton = make([]*cli.Command, 1, 1)
		subcommandSingleton[0] = command
		return
	}

	subcommandSingleton = append(subcommandSingleton, command)
}

func Get() []*cli.Command {
	if subcommandSingleton == nil {
		return []*cli.Command{}
	}

	return subcommandSingleton
}
