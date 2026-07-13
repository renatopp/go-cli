# Commands

A command is a unit of work in your CLI. Every CLI has a root command (implicit in `main()`), and you can add subcommands to organize functionality.

## Command Anatomy

Every command has four parts:

1. **Metadata** — Name, description, examples
2. **Arguments** — Flags and positionals defined before `Parse()`
3. **Configuration** — Via method chains (e.g., `AsHidden()`, `AsRequired()`)
4. **Logic** — Code after `Parse()` that runs if no subcommand was invoked

```go
func main() {
    cli.Name("myapp")
    cli.Description("My application")
    
    // Define arguments
    cli.Flag("config", "c", "Config file").WithDefault("config.yaml")
    cli.Pos("action", "Action to perform").AsRequired()
    
    // Define subcommands
    cli.Command("start", "Start the app", startCmd)
    cli.Command("stop", "Stop the app", stopCmd)

		// Setup configuration for the current command and its subcommands
		cli.AllowExtraFlags()
    
    cli.Parse()
    
    // Logic runs only if no subcommand was invoked
    println("Root command logic")
}
```

## Subcommands and Nesting

Subcommands global flags, and configuration from their parent. Create nested subcommands by defining commands within subcommands:

```go
func startCmd() {
    cli.Name("start")
    
    // These define arguments for this subcommand only
    cli.FlagBool("detached", "d", "Run in background")
    
    // Nested subcommand
    cli.Command("primary", "Start primary service", startPrimary)
    
    cli.Parse()
    println("Starting...")
}

func startPrimary() {
    cli.Parse()
    println("Starting primary service...")
}
```

## Propagation

When you define a flag or configuration at a command level:

- **Global flags** (`.AsGlobal()`) are accessible in all nested subcommands.
- **Configuration** like `AllowExtraFlags()` applies to that command and inherits downward.

## Modifiers

### AsHidden

Mark a command as hidden to exclude it from help output (but still executable):

```go
cli.Command("debug", "Debug utilities", debugCmd).AsHidden()
```

Hidden commands don't appear in help but work when invoked directly:

```bash
myapp debug  # Works
myapp help   # "debug" not listed
```
