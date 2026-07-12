// Package locales provides the built-in translations for go-cli's help text
// and error messages, along with the Locale type used to define custom ones.
// Pass a Locale to cli.SetLocale to change the language of the CLI output.
package locales

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

	// Error messages
	ErrInvalidFlagValue          string // %s = flag signature, %v = provided value
	ErrInvalidPositionalValue    string // %s = positional name, %v = provided value
	ErrMissingRequiredFlag       string // %s = flag signature
	ErrMissingRequiredPositional string // %s = positional name
	ErrUnknownFlag               string // %s = flag name
	ErrFlagSpecifiedMultiple     string // %s = flag name
	ErrMissingValueForFlag       string // %s = flag name
	ErrUnexpectedExtraPositional string // %s = token
	ErrExclusiveFlags            string // %s = flag signatures joined by "and"
	ErrAtLeastOneFlag            string // %s = flag signatures joined by "or"
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

		ErrInvalidFlagValue:          "invalid value for flag %s: %v",
		ErrInvalidPositionalValue:    "invalid value for positional argument %s: %v",
		ErrMissingRequiredFlag:       "missing required flag %s",
		ErrMissingRequiredPositional: "missing required positional argument: %s",
		ErrUnknownFlag:               "unknown flag %s",
		ErrFlagSpecifiedMultiple:     "flag %s was specified multiple times",
		ErrMissingValueForFlag:       "missing value for flag %s",
		ErrUnexpectedExtraPositional: "unexpected extra positional argument: %s",
		ErrExclusiveFlags:            "mutually exclusive flags provided: %s",
		ErrAtLeastOneFlag:            "at least one of the following flags must be provided: %s",
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

		ErrInvalidFlagValue:          "valor inválido para a flag %s: %v",
		ErrInvalidPositionalValue:    "valor inválido para o argumento posicional %s: %v",
		ErrMissingRequiredFlag:       "flag obrigatória ausente %s",
		ErrMissingRequiredPositional: "argumento posicional obrigatório ausente: %s",
		ErrUnknownFlag:               "flag desconhecida %s",
		ErrFlagSpecifiedMultiple:     "flag %s foi especificada múltiplas vezes",
		ErrMissingValueForFlag:       "valor ausente para a flag %s",
		ErrUnexpectedExtraPositional: "argumento posicional extra inesperado: %s",
		ErrExclusiveFlags:            "flags mutuamente exclusivas fornecidas: %s",
		ErrAtLeastOneFlag:            "pelo menos uma das seguintes flags deve ser fornecida: %s",
	}
}
