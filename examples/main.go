package main

import (
	"fmt"

	"github.com/renatopp/go-cli"
)

func main() {
	cli.Name("stub")
	cli.Description("This is a stub cli to showcase the library features.")
	cli.Parse()

	fmt.Printf("Arguments: %v\n", cli.Args())
	fmt.Printf("Extra Arguments: %v\n", cli.ExtraArgs())
}
