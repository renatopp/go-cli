package formatters

import (
	"fmt"

	"github.com/renatopp/go-cli/errors"
	"github.com/renatopp/go-cli/locales"
)

// DefaultErrorFormatter is the built-in error style. It prefixes the error
// message with the localized error label, e.g. "Error: unknown flag x", and
// appends the "did you mean" hint when the error carries a suggestion.
func DefaultErrorFormatter(err error, loc locales.Locale) string {
	msg := fmt.Sprintf("%s: %s", loc.ErrorLabel, loc.LocalizeError(err))

	var cliErr *errors.CliError
	if errors.As(err, &cliErr) && cliErr.Suggestion != "" && loc.DidYouMeanLabel != "" {
		msg += ", " + fmt.Sprintf(loc.DidYouMeanLabel, cliErr.Suggestion)
	}

	return errorStyle(msg)
}
