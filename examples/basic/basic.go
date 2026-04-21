package main

import "github.com/renatopp/go-cli"

func main() {
  cli.Name("hello")
  cli.Description("Prints a classical message.")
  cli.AutoHelp(true)
  cli.Parse()

  println("Hello, World!")
}