package main

import (
	"github.com/renatopp/go-cli/cli"
	"github.com/renatopp/go-cli/cli/locales"
)

func main() {
	cli.Name("locale")
	cli.Description("A simple CLI application to demonstrate locale usage.")
	cli.FlagString("size", "s", "sample")
	cli.FlagBool("other", "o", "sample")
	cli.Pos("input", "sample")
	cli.Cmd("command", "sample", func() {})
	cli.AutoHelp(true)
	cli.SetLocale(locales.PTBR())
	cli.Version("0.0.0")
	cli.Parse()
	cli.ShowHelp()
}
