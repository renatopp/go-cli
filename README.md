This is a simple, minimalist, intuitive CLI library that works as a state machine.

**Why?**

I created go-cli because I feel that other existing Go's CLI libraries are too complex and require too much boilerplate, demanding documentation lookup every time you use them. This library was made to be intuitive, configurations are explicit, it does not require any additional structure but functions (and only if you're using subcommands) -- you can discover the API as you go.

- No boilerplate code
- No need for scaffolding CLI
- No dependencies

> **Stability Notice!** This library is still in v0.x.x and its API may change in future releases. Please check the [CHANGELOG](CHANGELOG.md) check its evolution.

## Table of Contents

- [Features](#features)
- [The Bad Parts](#the-bad-parts)
- [What does it Look Like?](#what-does-it-look-like)
- [Installation](#installation)
- [User Guide](#user-guide)
- [Cookbook](#cookbook)
- [Examples](#examples)

## Features

General:
- **Strict rules by default**, unless explicitly allows.
- **Nested _lazy_ subcommands**.
- **Localization**.
- **Custom help and error formatting**.
- **Examples on help**.

Arguments (flags and positionals):
- **Global flags** which can be used in all nested subcommands.
- **Auto help** and **auto version** options (using --help, -h and --version).
- **Repeated flags**, collecting all values.
- **Required, default values** and **custom validation** for flags and positionals.
- **Positional variadic**, collecting all positional arguments in the end of the list.
- Optional **unknown flags** and **unknown positional arguments**.
- **Custom flags and positional types**.
- **Bind values to environment variables**.

This is the complete list of supported flags and positional arguments syntax:

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


## The Bad Parts

As any piece of technology, this library has its trade-offs:

- It uses a global instance by default, similar to what logging libraries do (eg. builtin `log`). You can create your own instance of `App` if you want to avoid the global state but its interface is not as ergonomic.
- It requires that you call manually `cli.Parse()` in all commands and subcommands. Missing this call my result in some unexpected behaviors. 
- It uses interruptions (`os.Exit`) as default to stop the flow of the CLI, which is not always desired. You may use `cli.UsePanic` to change how the interruption occours, but again, not always desired.

For most use cases, these trade-offs are acceptable.

## What does it Look Like?

```go
// Example of subcommands
func main() {
  cli.Name("git")
  cli.Description("These are common Git commands used in various situations.")
  cli.Command("clone", "Clone a repository into a new directory.", clone)
  cli.Command("commit", "Record changes to the repository.", commit)
  cli.Command("push", "Update remote refs along with associated objects.", push)
  cli.AutoHelp(true)
  cli.Parse()

  // Show help if no command is provided
  cli.Help()
}

// Example of positional usage
func clone() {
  repo := cli.Pos("repository", "The url to clone.").AsRequired()
  name := cli.Pos("name", "The name of the directory to clone into.") // optional
  cli.Parse()
	
  println("Cloning repository:", repo.Value())
  if name.IsProvided() {
    println("Into directory:", name.Value())
  }
}

// Example of flag usage
func commit() {
  message := cli.Flag("m", "Commit message.").AsRequired()
  cli.Parse()

  println("Committing with message:", message.Value())	
}

func push() {
  cli.Parse()

  println("Pushing changes...")
}
```

## Installation

```bash
go get github.com/renatopp/go-cli
```

After that, just import the package and use the `cli` name:

```go
import "github.com/renatopp/go-cli"

func main() {
  cli.Name("hello")
  cli.Description("Prints a classical message.")
  cli.AutoHelp(true)
  cli.Parse()
  println("Hello, World!")
}
```

## Overview

- [Getting Started](.docs/guide/getting-started.md)
- [Commands](.docs/guide/commands.md)
- [Flags](.docs/guide/flags.md)
- [Positionals](.docs/guide/positionals.md)

## Cookbook

- [Changing localization](.docs/cookbook/localization.md)
- [Customizing help and error formatting](.docs/cookbook/formatting.md)
- [Custom Flags and Positionals](.docs/cookbook/custom-types.md)
- [Mutually exclusive flags](.docs/cookbook/mutually-exclusive-flags.md)
- [Testing CLI applications](.docs/cookbook/testing.md)

## Examples

Check the [samples folder](.docs/samples) folder for more usage examples.
