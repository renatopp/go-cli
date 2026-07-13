# Positionals

Positional arguments are ordered, unnamed parameters passed without a flag prefix. They're matched in the order they're defined.

```go
func main() {
    source := cli.Pos("source", "Source file")
    dest := cli.Pos("dest", "Destination file")
    
    cli.Parse()
    
    println(source.Value(), dest.Value())  // myapp file.txt out.txt
}
```

## Value Access

After `Parse()`, use these methods to access positional values:

```go
func main() {
    dur := cli.PosDuration("duration", "Duration").AsRequired()
    cli.Parse()
    
    // Get the parsed value (or default if not provided), as time.Duration
    println(dur.Value())
    
    // Get the raw string from user input
    println(dur.RawValue())
    
    // Check if user provided the positional
    if dur.IsProvided() {
        println("Duration was provided, seconds:", dur.Value().Seconds())
    }
}
```

- **`Value()`** — Parsed value; returns default if not provided
- **`RawValue()`** — Raw string from user input
- **`IsProvided()`** — True if user passed the argument

## Modifiers

### WithDefault

Set a default value if the user doesn't provide the positional:

```go
cli.Pos("format", "Output format").WithDefault("json")
```

### WithEnv

Bind to an environment variable as a fallback:

```go
cli.Pos("config", "Config file").WithEnv("APP_CONFIG")
```

Order of precedence: user input → env var → default value.

### WithValidation

Add custom validation after parsing:

```go
cli.Pos("count", "Count").WithValidation(func(s string) error {
    n, err := strconv.Atoi(s)
    if err != nil || n < 0 {
        return fmt.Errorf("count must be a non-negative number")
    }
    return nil
})
```

### AsRequired

Require the user to provide the positional:

```go
cli.Pos("action", "Action to perform").AsRequired()
```

Missing required positionals cause a parse error.

### AsHidden

Exclude the positional from help output:

```go
cli.Pos("internal", "Internal argument").AsHidden()
```

Hidden positionals still work when provided; they're just not listed in help.

### AsVariadic

Capture all remaining arguments into a single positional:

```go
func main() {
    action := cli.Pos("action", "Action").AsRequired()
    args := cli.Pos("args", "Arguments").AsVariadic()
    
    cli.Parse()
    
    // User: myapp exec arg1 arg2 arg3
    println(action.Value())      // "exec"
    println(args.Value())        // "arg3" (last value)
    println(len(args.Values()))  // 3
    println(args.Count())        // 3
}
```

- **`Value()`** — Last argument captured
- **`Values()`** — All arguments as a slice
- **`Count()`** — Number of arguments captured

**Rules for variadic:**
- Only the last positional can be variadic
- A variadic positional captures all remaining positionals
- Flags after `--` are treated as positional arguments (and will be captured by variadic)

## Common Pattern: Variable Arguments

Forward all remaining arguments to another program:

```go
func main() {
    program := cli.Pos("program", "Program to run").AsRequired()
    args := cli.Pos("args", "Program arguments").AsVariadic()
    
    cli.Parse()
    
    // Execute with: myapp myprogram arg1 arg2 arg3
    cmd := exec.Command(program.Value(), args.Values()...)
    cmd.Run()
}
```
