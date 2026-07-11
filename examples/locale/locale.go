package main

import (
	"github.com/renatopp/go-cli"
	"github.com/renatopp/go-cli/locales"
)

func main() {
	cli.Name("locale")
	cli.Description("A simple CLI application to demonstrate locale usage.")
	cli.Flag("size", "s", "sample")
	cli.FlagBool("other", "o", "sample")
	cli.Pos("input", "sample")
	cli.Command("command", "sample", func() {})
	cli.AutoHelp(true)
	cli.SetLocale(locales.PTBR())
	cli.Version("0.0.0")
	cli.Parse()
	cli.ShowHelp()
}
