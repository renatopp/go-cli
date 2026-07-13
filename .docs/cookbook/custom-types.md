# Custom Types

Define flags and positionals with custom types using parser functions.

## Custom Flag Type

Create a parser function and pass it to `FlagFunc()`:

```go
type LogLevel int

const (
    Debug LogLevel = iota
    Info
    Warn
    Error
)

func ParseLogLevel(s string) (LogLevel, error) {
    switch s {
    case "debug":
        return Debug, nil
    case "info":
        return Info, nil
    case "warn":
        return Warn, nil
    case "error":
        return Error, nil
    default:
        return 0, fmt.Errorf("invalid log level: %s", s)
    }
}

func main() {
    level := cli.FlagFunc("level", "l", "Log level", ParseLogLevel)
    cli.Parse()
    
    switch level.Value() {
    case Debug:
        log.SetLevel(log.DebugLevel)
    // ...
    }
}
```

## Custom Positional Type

Same pattern with `PosFunc()`:

```go
duration := cli.PosFunc("duration", "Duration", time.ParseDuration)
cli.Parse()

timeout := duration.Value()  // time.Duration
```

## Combining with Validation

Add validation on top of custom parsing:

```go
port := cli.FlagFunc("port", "p", "Port", ParsePort).
    WithValidation(func(p int) error {
        if p < 1 || p > 65535 {
            return fmt.Errorf("port must be 1-65535")
        }
        return nil
    })
```

The parser converts the string to your type; validation checks the converted value.
