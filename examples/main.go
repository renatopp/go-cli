package main

import (
	"github.com/renatopp/go-cli"
)

func main() {
	// fmt.Printf("%+v\n", strings.Join(os.Args, "|"))
	cli.Name("stub")
	cli.Description("This is a stub cli to showcase the library features.")
	// cli.Pos("p_a", "first positional argument")
	// cli.PosInt("p_b", "second positional argument")
	// cli.Flag("f_a", "a", "first flag").AsRequired()
	// cli.Flag("f_b", "b", "second flag").WithDefault("defaulted")
	cli.Command("c1", "sample", func() {
		cli.ShowHelp()
	})
	cli.Command("c2", "sample", func() {})
	cli.Parse()
	cli.ShowHelp()

	// fmt.Printf("Arguments: %v\n", cli.Args())
	// fmt.Printf("Extra Arguments: %v\n", cli.ExtraArgs())
	// fmt.Printf("Positional Argument 1: %v\n", p1.Value())
	// fmt.Printf("Positional Argument 2: %v\n", p2.Value())
}
