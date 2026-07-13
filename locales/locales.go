// Package locales provides the built-in translations for go-cli's help text
// and error messages, along with the Locale type used to define custom ones.
// Pass a Locale to cli.SetLocale to change the language of the CLI output.
package locales

import (
	"fmt"

	"github.com/renatopp/go-cli/errors"
)

// Locale holds all the user-facing strings used by go-cli, such as help text
// labels and error messages. Every field that accepts arguments follows the
// standard fmt formatting verbs (e.g. %s, %v) and is rendered with
// fmt.Sprintf/fmt.Errorf, so you can reorder placeholders with indexed verbs
// (e.g. %[1]s) or drop them as needed when translating.
type Locale struct {
	// Help text labels
	UsageLabel        string // e.g. "Usage"
	UsageCommandLabel string // e.g. "as in <command>"
	UsageOptionsLabel string // e.g. "as in [options]"
	CommandsLabel     string // e.g. "Commands"
	OptionsLabel      string // e.g. "Options"
	ArgumentsLabel    string // e.g. "Arguments"
	FlagGlobalLabel   string // e.g. "as in (global)"
	FlagRequiredLabel string // e.g. "as in (required)"
	FlagDefaultLabel  string // e.g. "as in (default=%v)", receives the default value
	ErrorLabel        string // e.g. "as in Error: <message>"

	// Auto-generated flag descriptions
	HelpFlagDescription    string // description for the automatic --help/-h flag
	VersionFlagDescription string // description for the automatic --version/-v flag

	// Error messages mapped by error code
	Errors map[errors.ErrorCode]string
}

// LocalizedError renders err using this locale's messages. If err is not (or
// does not wrap) an *errors.CliError, or this locale has no message
// registered for its code, the error's own Error() string is returned
// instead.
func (l Locale) LocalizeError(err error) string {
	var cliErr *errors.CliError
	if !errors.As(err, &cliErr) {
		return err.Error()
	}

	tmpl, ok := l.Errors[cliErr.Code]
	if !ok {
		return cliErr.Error()
	}
	return fmt.Sprintf(tmpl, cliErr.Parameters...)
}

// EN returns the built-in English locale, which is the default used by go-cli.
func EN() Locale {
	return Locale{
		UsageLabel:        "Usage",
		UsageCommandLabel: "command",
		UsageOptionsLabel: "options",
		CommandsLabel:     "Commands",
		OptionsLabel:      "Options",
		ArgumentsLabel:    "Arguments",
		FlagGlobalLabel:   "global",
		FlagRequiredLabel: "required",
		FlagDefaultLabel:  "default=%v",
		ErrorLabel:        "Error",

		HelpFlagDescription:    "Show help message",
		VersionFlagDescription: "Show version information",

		Errors: map[errors.ErrorCode]string{
			errors.ErrUnknownFlag:         "unknown flag %s",
			errors.ErrMissingRequiredFlag: "missing required flag %s",
			errors.ErrMissingRequiredPos:  "missing required positional argument: %s",
			errors.ErrInvalidFlagValue:    "invalid value for flag %s: %v",
			errors.ErrInvalidPosValue:     "invalid value for positional argument %s: %v",
			errors.ErrRepeatedFlag:        "flag %s was specified multiple times",
			errors.ErrMissingFlagValue:    "missing value for flag %s",
			errors.ErrUnexpectedPos:       "unexpected extra positional argument: %s",
			errors.ErrExclusiveFlags:      "mutually exclusive flags provided: %s",
			errors.ErrAtLeastOneFlag:      "at least one of the following flags must be provided: %s",
		},
	}
}

// PT_BR returns the built-in Brazilian Portuguese locale.
func PT_BR() Locale {
	return Locale{
		UsageLabel:        "Uso",
		UsageCommandLabel: "comando",
		UsageOptionsLabel: "opções",
		CommandsLabel:     "Comandos",
		OptionsLabel:      "Opções",
		ArgumentsLabel:    "Argumentos",
		FlagGlobalLabel:   "global",
		FlagRequiredLabel: "obrigatório",
		FlagDefaultLabel:  "padrão=%v",
		ErrorLabel:        "Erro",

		HelpFlagDescription:    "Exibir mensagem de ajuda",
		VersionFlagDescription: "Exibir informações da versão",

		Errors: map[errors.ErrorCode]string{
			errors.ErrUnknownFlag:         "flag desconhecida %s",
			errors.ErrMissingRequiredFlag: "flag obrigatória ausente %s",
			errors.ErrMissingRequiredPos:  "argumento posicional obrigatório ausente: %s",
			errors.ErrInvalidFlagValue:    "valor inválido para a flag %s: %v",
			errors.ErrInvalidPosValue:     "valor inválido para o argumento posicional %s: %v",
			errors.ErrRepeatedFlag:        "flag %s foi especificada múltiplas vezes",
			errors.ErrMissingFlagValue:    "valor ausente para a flag %s",
			errors.ErrUnexpectedPos:       "argumento posicional extra inesperado: %s",
			errors.ErrExclusiveFlags:      "flags mutuamente exclusivas fornecidas: %s",
			errors.ErrAtLeastOneFlag:      "pelo menos uma das seguintes flags deve ser fornecida: %s",
		},
	}
}
