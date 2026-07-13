# Testing

Write tests for your CLI without running the actual binary.

## Test with Custom Arguments

Use `ParseArgs()` instead of `Parse()` and provide test arguments:

```go
func TestMyCommand(t *testing.T) {
    cli.Clear()  // Reset global state
    cli.Name("myapp")
    
    output := cli.FlagString("output", "o", "Output file")
    cli.ParseArgs([]string{"--output", "test.txt"})
    
    if output.Value() != "test.txt" {
        t.Errorf("expected test.txt, got %s", output.Value())
    }
}
```

`ParseArgs()` takes a slice of arguments (without the program name) and parses them.

## Capture Output

Redirect stdout/stderr to test output:

```go
func TestHelpOutput(t *testing.T) {
    cli.Clear()
    cli.Name("myapp")
    cli.Description("My app")
    
    var buf bytes.Buffer
    cli.Stdout(&buf)
    
    cli.AutoHelp(true)
    cli.ParseArgs([]string{"--help"})
    cli.Help()  // Force help output
    
    output := buf.String()
    if !strings.Contains(output, "My app") {
        t.Error("help text missing description")
    }
}
```

## Use Panic Instead of Exit

Enable panic mode for tests to avoid actual process exit:

```go
func TestParseError(t *testing.T) {
    cli.Clear()
    cli.Name("myapp")
    cli.UsePanic(true)
    
    required := cli.FlagString("required", "r", "Required flag").AsRequired()
    
    defer func() {
        if r := recover(); r == nil {
            t.Error("expected panic, but none occurred")
        }
    }()
    
    cli.ParseArgs([]string{})  // No required flag provided
}
```

Panic mode lets you catch failures in tests without the process terminating.

## Test Subcommands

Test subcommand logic by calling it directly:

```go
func TestStartCommand(t *testing.T) {
    cli.Clear()
    
    detached := cli.FlagBool("detached", "d", "Detached mode")
    cli.UsePanic(true)
    
    cli.ParseArgs([]string{"--detached"})
    
    if !detached.IsProvided() {
        t.Error("detached flag not provided")
    }
}
```

Each test should call `cli.Clear()` at the start to reset the global state.
