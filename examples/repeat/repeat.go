package main

import (
	"strings"

	"github.com/renatopp/go-cli"
)

func main() {
	cli.Name("repeat")
	cli.Description("Repeat a string a specified number of times")
	cli.AutoHelp(true)

	// you can use `--times 3` or `-t 3`
	t := cli.FlagInt("time", "t", "Number of times to repeat the string").AsRequired()

	// message is a variadic positional argument, so you can provide multiple
	// values for it, and they will be joined together with spaces to form the
	// final message to repeat.
	m := cli.Pos("message", "Message to repeat").AsRequired().AsVariadic()

	cli.Parse()

	// Variadic positionals or repeated flags retrieve all their `Values()`
	msg := strings.Join(m.Values(), " ")

	// Single value positionals or flags can be retrieved with `Value()`
	for range t.Value() {
		println(msg)
	}
}
