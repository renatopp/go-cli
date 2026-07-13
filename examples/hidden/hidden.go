package main

import "github.com/renatopp/go-cli"

func main() {
	cli.Name("hidden")
	cli.Description("A CLI with a hidden subcommands, flags and positionals.")
	cli.AutoHelp(true)

	cli.Command("private", "an internal command", func() {}).AsHidden()
	cli.Command("public", "a public command", func() {})

	cli.FlagString("secret", "s", "a secret flag").AsHidden()
	cli.FlagString("public", "p", "a public flag")

	cli.Pos("hidden", "a hidden positional").AsHidden()
	cli.Pos("public", "public positional")
	cli.Parse()

	cli.Help()
}
