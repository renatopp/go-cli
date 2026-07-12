package main

import (
	"errors"
	"fmt"

	"github.com/renatopp/go-cli"
	"github.com/renatopp/go-cli/pkg"
	cerrors "github.com/renatopp/go-cli/pkg/errors"
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
		var cliErr *cerrors.CliError
		if errors.As(err, &cliErr) {
			if cliErr.Code == cerrors.ErrUnknownFlag {
				// Parameters: [0] = flag name
				return fmt.Sprintf("oops! I don't know the flag %q", cliErr.Parameters[0])
			}
		}
		return "oops! " + pkg.GetLocale().LocalizeError(err)
	})

	cli.Parse()
	cli.Help()
}
