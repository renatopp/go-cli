# Getting Started

## Why go-cli?

go-cli is a minimal CLI library for Go. It's designed to be intuitive and require minimal boilerplate — no scaffolding, no configuration objects, no dependencies. You define arguments directly in your code as you go, and the library handles parsing and help text automatically.

The library works as a **state machine**: commands flow sequentially, and when a subcommand is invoked, `Parse()` interrupts the parent and hands execution to the child.

## Core Pattern

Every command follows this pattern:

1. **Define metadata** — name, description
2. **Define arguments** — flags, positionals
3. **Call `Parse()`** — this parses arguments and may interrupt (if a subcommand was invoked)
4. **Run logic** — code after `Parse()` only executes if no subcommand was invoked

```go
func main() {
    cli.Name("myapp")
    cli.Description("My CLI application.")
    // Define arguments here
    cli.Parse()
    // This code runs only if no subcommand was invoked
}
```

## The Parse Flow

When you call `cli.Parse()`:

1. If a subcommand name is found, that subcommand function is called — code after `Parse()` in the parent is **skipped**.
2. Arguments are matched against defined flags and positionals.
3. If an unknown flag or extra positional is encountered (and not allowed), an error is printed.
4. Code after `Parse()` in the current function continues.

This is the **interruption** mechanism: `Parse()` can transfer control to a subcommand.

```go
func main() {
    cli.Command("start", "Start server", start)
    cli.Parse()
    
    // This line only executes if 'start' was NOT invoked
    println("No subcommand provided")
    cli.Help()
}

func start() {
    cli.Parse()
    // This line only executes if start has no subcommands that were invoked
    println("Starting server...")
}
```


## Commands: Root and Subcommands

The root command is created automatically — in the examples, everything you define in `main()` belongs to it.

Add subcommands with `cli.Command()`:

```go
func main() {
    cli.Name("git")
    cli.Command("clone", "Clone a repository", clone)
    cli.Command("commit", "Record changes", commit)
    cli.Parse()
}

func clone() {
    // This is the clone subcommand
    cli.Parse()
}

func commit() {
    // This is the commit subcommand
    cli.Parse()
}
```

Subcommands can have their own subcommands (nesting):

```go
func clone() {
    cli.Command("mirror", "Mirror clone", mirrorClone)
    cli.Parse()
}

func mirrorClone() {
    cli.Parse()
}
```

Check out the [commands guide](.docs/guide/commands.md) for more advanced patterns.

## Flags: Named Arguments

Flags are named arguments passed with `--long` or `-s` syntax. Define them before `Parse()`:

```go
func main() {
    message := cli.Flag("message", "m", "Commit message").AsRequired()
    verbose := cli.FlagBool("verbose", "v", "Verbose output")
    
    cli.Parse()
    
    println("Message:", message.Value())
    if verbose.IsProvided() {
        println("Verbose mode enabled")
    }
}
```

Use specialized flag creators for common types:

```go
cli.FlagString("name", "n", "Your name")
cli.FlagInt("count", "c", "Count")
cli.FlagBool("force", "f", "Force operation")
cli.FlagFloat("ratio", "r", "Ratio")
cli.FlagDuration("timeout", "t", "Timeout (e.g., '30s', '5m')")
```

Check out the [flags guide](.docs/guide/flags.md) for more information.


## Positionals: Positional Arguments

Positional arguments are ordered, not named. Define them before `Parse()`:

```go
func main() {
    source := cli.Pos("source", "Source file").AsRequired()
    dest := cli.Pos("dest", "Destination file")  // optional
    
    cli.Parse()
    
    println("Copying:", source.Value())
    if dest.IsProvided() {
        println("To:", dest.Value())
    }
}
```

Like flags, positionals have typed variants:

```go
cli.PosString("name", "Name")
cli.PosInt("count", "Count")
cli.PosBool("enabled", "Enabled")
cli.PosFloat("ratio", "Ratio")
cli.PosDuration("timeout", "Timeout")
```

Check out the [positionals guide](.docs/guide/positionals.md) for more information.

## Help and Errors

### Automatic Help

Enable automatic help with `-h` or `--help`:

```go
cli.AutoHelp(true)
cli.Parse()
```

If enabled, help is shown and the app exits when `--help` or `-h` is provided.

### Manual Help

Show help explicitly:

```go
cli.Help()
```

### Examples

Add examples to help text:

```go
cli.Example("myapp --verbose file.txt", "Run with verbose output")
cli.Example("myapp file.txt --output out.txt", "Specify output")
```

Examples appear in the help output.

### Error Handling

By default, parsing errors cause the app to exit with code 1. You can:

- **Print error and exit manually**: use `cli.Fatal()` or `cli.FatalIf()`
- **Switch to panic**: use `cli.UsePanic(true)` to panic instead of exit (useful for testing)

```go
cli.FatalIf(someError)
cli.Fatal("Something went wrong: %v", value)
```

### Customization

Check out the [formatting recipe](.docs/cookbook/formatting.md) for more information on how customize help and errors.

## Global Configuration

These apply to the entire app and should be called before `Parse()`:

```go
cli.Name("myapp")                              // App name
cli.Description("My application")               // Description
cli.Version("1.0.0")                            // Enable --version
cli.AutoHelp(true)                              // Enable automatic help
cli.AllowExtraFlags(true)                       // Allow unknown flags
cli.AllowExtraPos(true)                         // Allow extra positionals
cli.AllowRepeatedFlags(true)                    // Allow repeated flags
cli.Locale(locales.PT_BR())                     // Localization
cli.HelpFormatter(customHelpFormatter)          // Custom help format
cli.ErrorFormatter(customErrorFormatter)        // Custom error format
cli.Stdout(customWriter)                        // Redirect output
cli.Stderr(customWriter)                        // Redirect errors
```

## Limitations

- **Global instance**: By default, go-cli uses a global instance (like the `log` package). This simplifies the API but makes global state unavoidable. You can create your own `core.App` instance if needed, though the API is less ergonomic.
- **Manual Parse calls**: You must call `cli.Parse()` in every command and subcommand. Forgetting it can cause unexpected behavior.
- **Interruption model**: The CLI uses `os.Exit()` by default to stop execution when moving between commands. This is not always suitable; use `cli.UsePanic(true)` for testing.

## Next Steps

- Explore the [Commands guide](.docs/guide/commands.md) for advanced command patterns.
- Check [Flags guide](.docs/guide/flags.md) for detailed flag configuration.
- See [Positionals guide](.docs/guide/positionals.md) for advanced positional patterns.
- Check out the [Cookbook folder](.docs/cookbook/) for recipes on formatting, validation, and more.
- Look at examples in [Samples folder](.docs/samples/) for working code.
