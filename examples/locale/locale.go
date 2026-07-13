package main

import (
	"github.com/renatopp/go-cli"
	"github.com/renatopp/go-cli/pkg/locales"
)

func main() {
	cli.Name("locale")
	cli.Description("A simple CLI application to demonstrate locale usage.")
	cli.FlagString("size", "s", "sample")
	cli.FlagBool("other", "o", "sample")
	cli.Pos("input", "sample")
	cli.Command("command", "sample", func() {})
	cli.AutoHelp(true)
	cli.Locale(locales.PT_BR())
	cli.Version("0.0.0")
	cli.Parse()
	cli.Help()
}
