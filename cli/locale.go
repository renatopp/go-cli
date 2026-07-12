package cli

import (
	"reflect"

	"github.com/renatopp/go-cli/cli/locales"
)

// Locale is re-exported here for convenience of the internal package.
type Locale = locales.Locale

// currentLocale holds the active locale used across the library. It defaults
// to the English locale returned by locales.EN.
var currentLocale = locales.EN()

// SetLocale replaces the active locale used for help text and error messages.
// Any field left as the zero value ("") falls back to the default English
// text, so callers can override only the strings they want to translate.
func SetLocale(l Locale) {
	d := locales.EN()

	lv := reflect.ValueOf(&l).Elem()
	dv := reflect.ValueOf(d)
	for i := 0; i < lv.NumField(); i++ {
		if lv.Field(i).String() == "" {
			lv.Field(i).Set(dv.Field(i))
		}
	}

	currentLocale = l
}

// GetLocale returns the currently active locale.
func GetLocale() Locale {
	return currentLocale
}
