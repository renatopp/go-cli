package internal

// Locale holds all the user-facing strings used by go-cli, such as help text
// labels and error messages. Every field that accepts arguments follows the
// standard fmt formatting verbs (e.g. %s, %v) and is rendered with
// fmt.Sprintf/fmt.Errorf, so you can reorder or drop placeholders as needed
// when translating.
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
}

// DefaultLocale returns the built-in English locale used by go-cli.
func DefaultLocale() Locale {
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
	}
}

func PTBRLocale() Locale {
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
	}
}

// currentLocale holds the active locale used across the library. It defaults
// to the English locale returned by DefaultLocale.
var currentLocale = DefaultLocale()

// SetLocale replaces the active locale used for help text and error messages.
// Any field left as the zero value ("") falls back to the default English
// text, so callers can override only the strings they want to translate.
func SetLocale(l Locale) {
	d := DefaultLocale()

	if l.UsageLabel == "" {
		l.UsageLabel = d.UsageLabel
	}
	if l.CommandsLabel == "" {
		l.CommandsLabel = d.CommandsLabel
	}
	if l.OptionsLabel == "" {
		l.OptionsLabel = d.OptionsLabel
	}
	if l.ArgumentsLabel == "" {
		l.ArgumentsLabel = d.ArgumentsLabel
	}
	if l.FlagRequiredLabel == "" {
		l.FlagRequiredLabel = d.FlagRequiredLabel
	}
	if l.FlagDefaultLabel == "" {
		l.FlagDefaultLabel = d.FlagDefaultLabel
	}
	if l.HelpFlagDescription == "" {
		l.HelpFlagDescription = d.HelpFlagDescription
	}
	if l.VersionFlagDescription == "" {
		l.VersionFlagDescription = d.VersionFlagDescription
	}
	if l.ErrInvalidFlagValue == "" {
		l.ErrInvalidFlagValue = d.ErrInvalidFlagValue
	}
	if l.ErrInvalidPositionalValue == "" {
		l.ErrInvalidPositionalValue = d.ErrInvalidPositionalValue
	}
	if l.ErrMissingRequiredFlag == "" {
		l.ErrMissingRequiredFlag = d.ErrMissingRequiredFlag
	}
	if l.ErrMissingRequiredPositional == "" {
		l.ErrMissingRequiredPositional = d.ErrMissingRequiredPositional
	}
	if l.ErrUnknownFlag == "" {
		l.ErrUnknownFlag = d.ErrUnknownFlag
	}
	if l.ErrFlagSpecifiedMultiple == "" {
		l.ErrFlagSpecifiedMultiple = d.ErrFlagSpecifiedMultiple
	}
	if l.ErrMissingValueForFlag == "" {
		l.ErrMissingValueForFlag = d.ErrMissingValueForFlag
	}
	if l.ErrUnexpectedExtraPositional == "" {
		l.ErrUnexpectedExtraPositional = d.ErrUnexpectedExtraPositional
	}

	currentLocale = l
}

// GetLocale returns the currently active locale.
func GetLocale() Locale {
	return currentLocale
}
