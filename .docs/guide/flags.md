# Flags

Flags are named arguments starting with `--` (long) or `-` (short). They're optional by default, though you can make them required with `.AsRequired()`.

## Flag Syntax

| Style                       | Description                                                     |
|-----------------------------|-----------------------------------------------------------------|
| `--long=VALUE`              | Long assigned flags                                             |
| `--long VALUE`              | Long unassigned flags                                           |
| `--long`                    | Long unvalued flags for *booleans*                              |
| `-s VALUE`                  | Short uncombined valued flags                                   |
| `-sVALUE`                   | Short combined valued flags                                     |
| `-s`                        | Short unvalued flags for *booleans*                             |
| `-abc`                      | Short combined unvalued flags for *booleans*                    |
| `-absVALUE`                 | Short combined valued flags for *booleans* and last non-boolean |
| `-abs VALUE`                | Short mixed valued flags for *booleans* and last non-boolean    |
| `--long VALUE --long VALUE` | Repeated long flags (when enabled)                              |
| `-sss`                      | Repeated short flags (when enabled)                             |
| `--`                        | End-of-options followed by forced positional arguments          |
| `-`                         | Single dash as positional                                       |

## Value Access

After `Parse()`, use these methods to access the flag value:

```go
func main() {
    age := cli.FlagInt("age", "a", "User age")
    cli.Parse()
    
    // Get the parsed value (or default if not provided), as Int
    println(age.Value())
    
    // Get the raw string from user input
    println(age.RawValue())
    
    // Check if user provided the flag
    if age.IsProvided() {
        println("User provided age", age.Value())
    }
}
```

- **`Value()`** — Parsed value; returns default if not provided
- **`RawValue()`** — Raw string from user input
- **`IsProvided()`** — True if user passed the flag

## Modifiers

### WithDefault

Set a default value if the user doesn't provide the flag:

```go
cli.FlagInt("port", "p", "Port number").WithDefault(8080)
```

### WithEnv

Bind to an environment variable as a fallback:

```go
cli.FlagString("token", "t", "API token").WithEnv("API_TOKEN")
```

Order of precedence: user input → env var → default value.

### WithValidation

Add custom validation after parsing:

```go
cli.FlagInt("port", "p", "Port").WithValidation(func(port int) error {
    if port < 1 || port > 65535 {
        return fmt.Errorf("port must be 1-65535")
    }
    return nil
})
```

### OnParsed

Run a callback after parsing (useful for global flags with side effects):

```go
debug := cli.FlagBool("debug", "d", "Debug mode").AsGlobal()
debug.OnParsed(func(f *core.Flag[bool]) {
    if f.Value() {
        log.SetLevel(LogDebug)
    }
})
```

The callback runs after all arguments are parsed, before subcommand execution.

### AsGlobal

Make the flag accessible in all nested subcommands:

```go
cli.FlagBool("verbose", "v", "Verbose").AsGlobal()
cli.Command("start", "Start", startCmd)
```

In `startCmd()`, you can access the global verbose flag.

### AsRequired

Require the user to provide the flag:

```go
cli.FlagString("config", "c", "Config file").AsRequired()
```

Missing required flags cause a parse error.

### AsRepeatable

Allow the flag to be passed multiple times:

```go
tags := cli.Flag("tag", "t", "Tag").AsRepeatable()
cli.Parse()

// User: myapp -t tag1 -t tag2 -t tag3
println(tags.Value())      // "tag3" (last value)
println(len(tags.Values())) // 3
println(tags.Count())       // 3
```

- **`Value()`** — Last value provided
- **`Values()`** — All values as a slice
- **`Count()`** — Number of times the flag was provided

### AsHidden

Exclude the flag from help output:

```go
cli.FlagBool("internal", "i", "Internal use").AsHidden()
```

Hidden flags still work when invoked directly; they're just not listed in help.
