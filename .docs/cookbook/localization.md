# Localization

Change the language of help text and error messages.

## Using Built-in Locales

Pass a built-in locale to `cli.Locale()`:

```go
import "github.com/renatopp/go-cli/locales"

func main() {
    cli.Name("myapp")
    cli.Locale(locales.PT_BR())  // Portuguese (Brazil)
    // ... rest of your CLI
}
```

Built-in locales: `EN()`, `PT_BR()`, etc. (check the `locales` package for full list).

## Custom Locale

Create a custom locale by filling the `Locale` struct:

```go
import "github.com/renatopp/go-cli/locales"

customLocale := locales.EN()  // Start from English
customLocale.CommandsLabel = "Subcomandos"
customLocale.OptionsLabel = "Opções"
customLocale.Errors[errors.ErrMissingRequiredFlag] = "falta a opção obrigatória %s"

cli.Locale(customLocale)
```

Fields left empty fall back to English defaults. Locales apply globally, regardless of which `App` instance you use.
