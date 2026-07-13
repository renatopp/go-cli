package main

import (
	"fmt"

	"github.com/renatopp/go-cli"
	"github.com/renatopp/go-cli/core"
	"github.com/renatopp/go-cli/errors"
	"github.com/renatopp/go-cli/formatters"
)

func main() {
	cli.Name("styles")
	cli.Description("A simple CLI application to demonstrate custom help and error styles.")
	cli.FlagString("size", "s", "sample")
	cli.AutoHelp(true)

	// Wrap the default help style with a custom banner.
	cli.HelpFormatter(func(cmd *core.Command, loc core.Locale) string {
		return "== STYLES ==\n\n" + formatters.DefaultHelpFormatter(cmd, loc)
	})

	// Customize error messages, inspecting typed errors for special cases.
	cli.ErrorFormatter(func(err error, loc core.Locale) string {
		var cliErr *errors.CliError
		if errors.As(err, &cliErr) {
			if cliErr.Code == errors.ErrUnknownFlag {
				// Parameters: [0] = flag name
				return fmt.Sprintf("oops! I don't know the flag %q", cliErr.Parameters[0])
			}
		}
		return "oops! " + loc.LocalizeError(err)
	})

	cli.Parse()
	cli.Help()
}
