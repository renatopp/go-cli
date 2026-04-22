# go-cli

This is a simple, minimalist, intuitive CLI library that works as a state machine.

**Why?**

I created this library because other existing CLI libs for Golang are complex or require too much boilerplate that requires checking documentation everytime I use it. This library is intuitive, configuration are explicit and does not require any additional structure but functions for subcommands.

## Install

```bash
go get github.com/renatopp/go-cli
```

After that, just import the package and use the `cli` name:

```go
import "github.com/renatopp/go-cli"

func main() {
  cli.Parse()
}
```

## Overview

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

Additionally, there are other interesting features:

- **Strict by default**.
- Optional **auto help** and **auto version**.
- Optional **positional variadic**, collecting all positional arguments in the list.
- Support for nested **subcommands**.
- Optional **repeated flags**, collecting all values.
- **Required, default value** and **custom validation** for flags and positional arguments.
- Support for **extra flags** and **extra positional arguments**.

## Usage and Cookbook

You can check all examples in `examples/` directory:

| Example                                  | What does it shows?                             |
|------------------------------------------|-------------------------------------------------|
| [basic](./examples/basic/basic.go)       | Minimal example                                 |
| [repeat](./examples/repeat/repeat.go)    | Basic flags and positionals                     |
| [list](./examples/list/list.go)          | Simple sub command usage                        |
| [verbose](./examples/verbose/verbose.go) | Flag repetition and counter `-vvv`              |
| [format](./examples/format/format.go)    | Mutually exclusive options `--json` or `--yaml` |


### Hello, World

```go
package main

import "github.com/renatopp/go-cli"

func main() {
  cli.Name("hello")
  cli.Description("Prints a classical message.")
  cli.AutoHelp(true)
  cli.Parse()

  println("Hello, World!")
}
```

```
$ hello
Hello, World!

$ hello --help
Usage: hello [options]

Prints a classical message.

Options:
  -h, --help        Show help message
``` 

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
