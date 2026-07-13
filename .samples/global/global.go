package main

import (
	"github.com/renatopp/go-cli"
	"github.com/renatopp/go-cli/core"
)

func main() {
	cli.Name("global")
	cli.Description("Example of global flags")
	cli.FlagBool("verbose", "v", "enable verbose output").AsGlobal()
	cli.FlagInt("level", "l", "set the level").AsGlobal().WithDefault(1)
	cli.FlagString("config", "c", "path to config file").WithDefault("config.yaml")
	cli.Command("cmd1", "", cmd1)
	cli.Command("cmd2", "", cmd2)
	cli.Parse()
	cli.Help()
}

func cmd1() {
	cli.FlagString("another", "a", "another flag").AsGlobal()
	cli.Command("cmd3", "", cmd3)
	cli.Parse()
	cli.Help()

	verbose, _ := cli.GetFlag[*core.GenericFlag[bool]]("verbose")
	println("verbose:", verbose.Value())
}

func cmd2() {
	cli.Parse()
	cli.Help()
}

func cmd3() {
	cli.Parse()
	cli.Help()
}
