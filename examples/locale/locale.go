package main

import "github.com/renatopp/go-cli"

func main() {
	cli.Name("locale")
	cli.Description("A simple CLI application to demonstrate locale usage.")
	cli.Flag("size", "s", "sample")
	cli.FlagBool("other", "o", "sample")
	cli.Pos("input", "sample")
	cli.Command("command", "sample", func() {})
	cli.AutoHelp(true)
	cli.SetLocale(cli.PTBR)
	cli.Version("0.0.0")
	cli.Parse()
	cli.ShowHelp()
}
