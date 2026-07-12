package main

import (
	"errors"
	"fmt"

	"github.com/renatopp/go-cli"
	"github.com/renatopp/go-cli/pkg"
)

func main() {
	cli.Name("styles")
	cli.Description("A simple CLI application to demonstrate custom help and error styles.")
	cli.FlagString("size", "s", "sample")
	cli.AutoHelp(true)

	// Wrap the default help style with a custom banner.
	cli.HelpFormatter(func(cmd *pkg.Command) string {
		return "== STYLES ==\n\n" + pkg.DefaultHelpFormatter(cmd)
	})

	// Customize error messages, inspecting typed errors for special cases.
	cli.ErrorFormatter(func(err error) string {
		var unknown *pkg.UnknownFlagError
		if errors.As(err, &unknown) {
			return fmt.Sprintf("oops! I don't know the flag %q", unknown.Name)
		}
		return "oops! " + err.Error()
	})

	cli.Parse()
	cli.Help()
}
