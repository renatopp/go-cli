# Formatting

Customize the appearance of help text and error messages.

## Custom Help Format

Wrap or replace the default help formatter:

```go
import (
    "github.com/renatopp/go-cli"
    "github.com/renatopp/go-cli/core"
    "github.com/renatopp/go-cli/formatters"
)

func main() {
    cli.Name("myapp")
    
    // Wrap default formatter with a banner
    cli.HelpFormatter(func(cmd *core.Command, loc core.Locale) string {
        header := "== MY APP ==\n\n"
        return header + formatters.DefaultHelpFormatter(cmd, loc)
    })
    
    cli.Parse()
}
```

The formatter receives the current command and locale, and returns the help text as a string.

## Custom Error Format

Inspect typed errors and customize messages:

```go
import (
    "github.com/renatopp/go-cli"
    "github.com/renatopp/go-cli/errors"
)

func main() {
    cli.Name("myapp")
    
    cli.ErrorFormatter(func(err error, loc core.Locale) string {
        var cliErr *errors.CliError
        if errors.As(err, &cliErr) {
            // Handle specific error types
            switch cliErr.Code {
            case errors.ErrUnknownFlag:
                return fmt.Sprintf("Unknown flag: %q", cliErr.Parameters[0])
            case errors.ErrMissingRequiredFlag:
                return fmt.Sprintf("You forgot: %s", cliErr.Parameters[0])
            }
        }
        // Fallback to localized message
        return loc.LocalizeError(err)
    })
    
    cli.Parse()
}
```

Error types: `ErrUnknownFlag`, `ErrMissingRequiredFlag`, `ErrMissingRequiredPos`, `ErrInvalidFlagValue`, etc.

## Redirect Output

Send help and error output to custom writers:

```go
cli.Stdout(customWriter)  // Help text
cli.Stderr(customWriter)  // Error messages
```

Useful for testing, logging, or file output.
