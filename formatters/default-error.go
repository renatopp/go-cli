package formatters

import (
	"fmt"

	"github.com/renatopp/go-cli/locales"
)

// DefaultErrorFormatter is the built-in error style. It prefixes the error
// message with the localized error label, e.g. "Error: unknown flag x".
func DefaultErrorFormatter(err error, loc locales.Locale) string {
	return errorStyle(fmt.Sprintf("%s: %s", loc.ErrorLabel, loc.LocalizeError(err)))
}
