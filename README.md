# Minimalist CLI

Simple CLI tool that works as a state machine.

## Install

```
go get github.com/renatopp/go-cli
```

## Usage

### Hello, World

```go
import "github.com/renatopp/go-cli"

func main() {
  cli.Name("hello")
  cli.Description("Prints a classical message.")
  cli.AutoHelp()
  cli.Parse()
  println("Hello, World!")
}
```

```
$ hello
Hello, World!

$ hello --help
Usage: hello

Prints a classical message.

``` 

### Flags

You can declare and use typed flags with the following functions:

- Flag (alias to FlagString)
- FlagString
- FlagInt
- FlagInt64
- FlagFloat (float64)
- FlagStringSlice
- FlagBool

All flags have long and short names, following a description, example: `Flag(long, short, description)`. You may mark it as required using `.AsRequired()` or customize the custom value as `WithDefault` chain methods.

```go
import "github.com/renatopp/go-cli"

func main() {
  cli.Name("hello")
  cli.Description("Prints a classical message.")
  cli.AutoHelp()
  name := cli.Flag("name", "n", "Your name.").WithDefault("World")
  cli.Parse()
  println("Hello, " + name.Value + "!")
}
```

```
$ hello
Hello, World!

$ hello -n Renato
Hello, Renato!

$ hello --help
Usage: hello [options]

Prints a classical message.

Options:
  -n, --name        Your name.

```

### Positional Arguments

Positional arguments works similarly to flags, you can define them with the following functions:

- Pos
- PosString
- PosInt
- PosInt64
- PosFloat

```go
import "github.com/renatopp/go-cli"

func main() {
  cli.Name("hello")
  cli.AutoHelp()
  verbose := cli.FlagBool("verbose", "v", "Verbose logs")
  first := cli.Pos(0, "input", "The config file").WithDefault("config.yaml")
  second := cli.Pos(1, "output", "The output file").WithDefault("output.yaml")
  cli.Parse()
  if verbose.Value {
    println(first.Value, "->", second.Value)
  }
}
```

```
$ hello

$ hello -v
config.yaml -> output.yaml

$ hello --verbose config.local.yaml
config.local.yaml -> output.yaml

$ hello --verbose config.local.yaml output.json
config.local.yaml -> output.json

```

You may also access positional arguments on demand as:

```
cli.Arg(0) // returning string
cli.Args() // returning all []string
```

### Subcommands

You can define sub commands recursively using the same functions for flags and positionals. Just remember to call `Parse` again inside the subcommand to parse the arguments properly.

```go
import "github.com/renatopp/go-cli"

func main() {
  cli.Name("hello")
  cli.AutoHelp()
  cli.Cmd("version", "Show the version.", cmdVersion)
  cli.Cmd("greet", "Greet someone", cmdGreet)
  cli.Parse()
  cli.ShowHelp() // If no subcommand is called
}

func cmdVersion() {
  // There is no need for arguments here
  println("1.0.0")
}

func cmdGreet() {
  name := cli.Flag("name", "n", "Your name.").AsRequired()
  msg := cli.Pos(0, "msg", "Message.").WithDefault("welcome")
  cli.Parse()
  fmt.Printf("Hello, %s, %s.\n", name.Value, msg.Value)
}
```

```
$ hello
Usage: hello <command>

Commands:
  version           Show the version.
  greet             Greet someone

$ hello version
1.0.0

$ hello greet
missing required flag: --name
exit status 1

$ hello greet --name=Renato 'is this the real life'
Hello, Renato, is this the real life.

$ hello greet --help
Usage: hello greet [options] [msg]

Options:
  -n, --name        Your name. Required.

```