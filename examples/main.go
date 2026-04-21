package main

import (
	"fmt"

	"github.com/renatopp/go-cli"
)

func main() {
	cli.Name("stub")
	cli.Description("This is a stub cli to showcase the library features.")
	p1 := cli.Pos("p_a", "first positional argument")
	p2 := cli.PosInt("p_b", "second positional argument")
	cli.Parse()

	fmt.Printf("Arguments: %v\n", cli.Args())
	fmt.Printf("Extra Arguments: %v\n", cli.ExtraArgs())
	fmt.Printf("Positional Argument 1: %v\n", p1.Value())
	fmt.Printf("Positional Argument 2: %v\n", p2.Value())
}
