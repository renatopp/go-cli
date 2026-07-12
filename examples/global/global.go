package main

import "github.com/renatopp/go-cli/cli"

func main() {
	cli.Name("global")
	cli.Description("Example of global flags")
	cli.FlagBool("verbose", "v", "enable verbose output").AsGlobal()
	cli.FlagInt("level", "l", "set the level").AsGlobal().WithDefault(1)
	cli.FlagString("config", "c", "path to config file").WithDefault("config.yaml")
	cli.Cmd("cmd1", "", cmd1)
	cli.Cmd("cmd2", "", cmd2)
	cli.Parse()
	cli.ShowHelp()
}

func cmd1() {
	cli.FlagString("another", "a", "another flag").AsGlobal()
	cli.Cmd("cmd3", "", cmd3)
	cli.Parse()
	cli.ShowHelp()

	verbose, _ := cli.GetFlag[*cli.GenericFlag[bool]]("verbose")
	println("verbose:", verbose.Value())
}

func cmd2() {
	cli.Parse()
	cli.ShowHelp()
}

func cmd3() {
	cli.Parse()
	cli.ShowHelp()
}
