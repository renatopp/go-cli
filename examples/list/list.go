package main

import (
	"fmt"

	"github.com/renatopp/go-cli/cli"
)

func main() {
	cli.Name("list")
	cli.Description("List folders and files")
	cli.Cmd("folders", "List folders", cmdFolders)
	cli.Cmd("files", "List files", cmdFiles)

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
