package main

import "github.com/renatopp/go-cli"

func main() {
	cli.Name("global")
	cli.Description("Example of global flags")
	cli.FlagBool("verbose", "v", "enable verbose output").AsGlobal()
	cli.FlagInt("level", "l", "set the level").AsGlobal().WithDefault(1)
	cli.Flag("config", "c", "path to config file").WithDefault("config.yaml")
	cli.Command("cmd1", "", cmd1)
	cli.Command("cmd2", "", cmd2)
	cli.Parse()
	cli.ShowHelp()
}

func cmd1() {
	cli.Flag("another", "a", "another flag").AsGlobal()
	cli.Command("cmd3", "", cmd3)
	cli.Parse()
	cli.ShowHelp()
}

func cmd2() {
	cli.Parse()
	cli.ShowHelp()
}

func cmd3() {
	cli.Parse()
	cli.ShowHelp()
}
