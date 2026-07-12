package main

import "github.com/renatopp/go-cli"

func main() {
	cli.Name("verbose")
	cli.Description("A simple example of using the verbose flag with repetitions")
	verbose := cli.FlagBool("", "v", "Enable verbose output. -v for verbose, -vv for more verbose, -vvv for maximum verbosity").AsRepeatable()
	cli.Parse()

	println("Vebosity Level:", min(3, len(verbose.Values())))
}
