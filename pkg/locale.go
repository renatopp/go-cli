package pkg

import (
	"reflect"

	cerrors "github.com/renatopp/go-cli/pkg/errors"
	"github.com/renatopp/go-cli/pkg/locales"
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
		fv := lv.Field(i)
		// Skip the Messages map - handle it separately below
		if lv.Type().Field(i).Name == "Messages" {
			continue
		}
		if fv.String() == "" {
			fv.Set(dv.Field(i))
		}
	}

	// Fill in missing error messages from the default locale
	if l.Errors == nil {
		l.Errors = make(map[cerrors.ErrorCode]string)
	}
	for code, message := range d.Errors {
		if _, exists := l.Errors[code]; !exists {
			l.Errors[code] = message
		}
	}

	currentLocale = l
}

// GetLocale returns the currently active locale.
func GetLocale() Locale {
	return currentLocale
}
