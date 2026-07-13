This is a simple, minimalist, intuitive CLI library that works as a state machine.

**Why?**

I created go-cli because I feel that other existing Go's CLI libraries are too complex and require too much boilerplate, demanding documentation lookup every time you use them. This library was made to be intuitive, configurations are explicit, it does not require any additional structure but functions (and only if you're using subcommands) -- you can discover the API as you go.

- No boilerplate code
- No need for scaffolding CLI
- No dependencies

## Table of Contents

- [Features](#features)
- [The Bads](#the-bads)
- [Installation](#installation)
- [Overview](#overview)
  - [Commands and Subcommands](#overview)
  - [Flags](#overview)
  - [Positional Arguments](#overview)
- [Cookbook](#cookbook)
  - [Flags and Positional Arguments](#flags-and-positional-arguments)
  - [Sub Commands](#sub-commands)

## Features

- **Strict by default**.
- Nested **subcommands**.
- Optional **global flags**, applying to all nested subcommands.
- Optional **auto help** and **auto version**.
- Optional **repeated flags**, collecting all values.
- **Required, default values** and **custom validation** for flags and positional arguments.
- Optional **positional variadic**, collecting all positional arguments in the list.
- Support for **extra flags** and **extra positional arguments**.
- **Locale** support.
- **Custom help and error formatting**.
- **Custom flags**

## The Bads

As any piece of technology, this library has its trade-offs:

- It uses a global instance by default, similar to what logging library do (eg. `log`). You can create your own instance of `App` if you want to avoid global state but it is not as ergonomic.
- It requires that you call `cli.Parse()` at all commands and subcommands. Missing this call will result in some unexpected behaviors. 
- It uses interruptions (os.Exit) as default to stop the flow of the CLI, which is not always desired. You may use `cli.UsePanic` to change how the interruption occours, but again, not always desired.

For most use cases, these trade-offs are acceptable.

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

Go-cli uses some traditional conventions for flags and positional arguments, supporting POSIX/Unix/GNU syntax style. The table below shows the syntax used for flags: 

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


## Usage and Cookbook

You can check all examples in `examples/` directory:

| Example                                  | What does it shows?                             |
|------------------------------------------|-------------------------------------------------|
| [hello](./examples/hello/hello.go)       | Minimal example                                 |
| [repeat](./examples/repeat/repeat.go)    | Basic flags and positionals                     |
| [list](./examples/list/list.go)          | Simple sub command usage                        |
| [verbose](./examples/verbose/verbose.go) | Flag repetition and counter `-vvv`              |
| [format](./examples/format/format.go)    | Mutually exclusive options `--json` or `--yaml` |
| [hidden](./examples/hidden/hidden.go)    | Hidden subcommands, flags and positionals       |


### Flags and Positional Arguments

Flags and positional arguments can be collected easily, just define them **before** the `cli.Parse` call.

```go
package main

import (
	"strings"

	"github.com/renatopp/go-cli"
)

func main() {
	cli.Name("repeat")
	cli.Description("Repeat a string a specified number of times")
	cli.AutoHelp(true)

	// you can use `--times 3` or `-t 3`
	t := cli.FlagInt("times", "t", "Number of times to repeat the string").AsRequired()

	// message is a variadic positional argument, so you can provide multiple
	// values for it, and they will be joined together with spaces to form the
	// final message to repeat.
	m := cli.Pos("message", "Message to repeat").AsRequired().AsVariadic()

	cli.Parse()

	// Variadic positionals or repeated flags retrieve all their `Values()`
	msg := strings.Join(m.Values(), " ")

	// Single value positionals or flags can be retrieved with `Value()`
	for range t.Value() {
		println(msg)
	}
}
```

```
$ repeat -t 3 Hello, world!
Hello, world!
Hello, world!
Hello, world!

$ repeat --help
Usage: repeat [options] <message>

Repeat a string a specified number of times

Options:
  -t, --times       (required) Number of times to repeat the string
  -h, --help        Show help message

Arguments:
  message           (required) Message to repeat
```

By default, `cli.Flag` and `cli.Pos` are strings, but you can use other typed versions:

| Flag Version        | Positional Version | Value Example    |
|---------------------|--------------------|------------------|
| `FlagString`        | `PosString`        | `'name'`         |
| `FlagInt`           | `PosInt`           | `42`             |
| `FlagInt8`          | `PosInt8`          | `-1`             |
| `FlagInt16`         | `PosInt16`         | `-100`           |
| `FlagInt32`         | `PosInt32`         | `-1000`          |
| `FlagInt64`         | `PosInt64`         | `-9999`          |
| `FlagUint`          | `PosUint`          | `42`             |
| `FlagUint8`         | `PosUint8`         | `1`              |
| `FlagUint16`        | `PosUint16`        | `100`            |
| `FlagUint32`        | `PosUint32`        | `1000`           |
| `FlagUint64`        | `PosUint64`        | `9999`           |
| `FlagFloat`         | `PosFloat`         | `3.14`           |
| `FlagFloat32`       | `PosFloat32`       | `3.14`           |
| `FlagFloat64`       | `PosFloat64`       | `3.14159`        |
| `FlagBool`          | `PosBool`          | `true`           |
| `FlagDuration`      | `PosDuration`      | `5s`             |
| `FlagIntSlice`      | `PosIntSlice`      | `1,2,3`          |
| `FlagInt8Slice`     | `PosInt8Slice`     | `1,2,3`          |
| `FlagInt16Slice`    | `PosInt16Slice`    | `1,2,3`          |
| `FlagInt32Slice`    | `PosInt32Slice`    | `1,2,3`          |
| `FlagInt64Slice`    | `PosInt64Slice`    | `1,2,3`          |
| `FlagUintSlice`     | `PosUintSlice`     | `1,2,3`          |
| `FlagUint8Slice`    | `PosUint8Slice`    | `1,2,3`          |
| `FlagUint16Slice`   | `PosUint16Slice`   | `1,2,3`          |
| `FlagUint32Slice`   | `PosUint32Slice`   | `1,2,3`          |
| `FlagUint64Slice`   | `PosUint64Slice`   | `1,2,3`          |
| `FlagFloatSlice`    | `PosFloatSlice`    | `1.1,2.2`        |
| `FlagFloat32Slice`  | `PosFloat32Slice`  | `1.1,2.2`        |
| `FlagFloat64Slice`  | `PosFloat64Slice`  | `1.1,2.2`        |
| `FlagBoolSlice`     | `PosBoolSlice`     | `1,0,true,false` |
| `FlagDurationSlice` | `PosDurationSlice` | `1m,4s`          |

All flags can be marked `AsRequired`, `AsRepeatable`, `WithDefault` or `WithValidation`.

### Sub Commands

```go
package main

import (
	"fmt"

	"github.com/renatopp/go-cli"
)

func main() {
	cli.Name("list")
	cli.Description("List folders and files")
	cli.Command("folders", "List folders", cmdFolders)
	cli.Command("files", "List files", cmdFiles)

	// if a command is provided, parse will exit after executing it, so the code
	// after this won't be executed.
	cli.Parse()

	// will only execute this if no subcommand is provided.
	cli.ShowHelp()
}

func cmdFolders() {
	cli.Description("list all folders in the path")
	filter := cli.FlagString("filter", "f", "Filter folders by name").WithDefault("*")
	path := cli.Pos("path", "Path to list folders from").WithDefault(".")
	cli.Parse()

	fmt.Printf("I should list folders in %s with filter %s\n", path.Value(), filter.Value())
}

func cmdFiles() {
	cli.Description("list all files in the path")
	filter := cli.FlagString("filter", "f", "Filter files by name").WithDefault("*")
	path := cli.Pos("path", "Path to list files from").WithDefault(".")
	cli.Parse()

	fmt.Printf("I should list files in %s with filter %s\n", path.Value(), filter.Value())
}

```

```
$ list
Usage: list <command>

List folders and files

Commands:
  folders           List folders
  files             List files

$ list folders
I should list folders in . with filter *

$ list files -f *.png images/
I should list files in images/ with filter *.png
```
