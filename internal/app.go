package internal

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func printf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

type App struct {
	path                    []string // the path of commands leading to the current command, e.g., ["git", "commit"]
	queue                   []string // the queue of arguments to be parsed
	rootCommand             *Command
	currentCommand          *Command
	arguments               *Arguments // parsed arguments
	Stdout                  func(format string, a ...any)
	Stderr                  func(format string, a ...any)
	PanicInsteadOfExit      bool
	ExtraFlagsAllowed       bool
	ExtraPositionalsAllowed bool
	RepeatedFlagsAllowed    bool
}

func NewApp() *App {
	s := &App{}
	s.Clear()
	return s
}

func (a *App) Clear() {
	a.path = []string{}
	a.queue = os.Args[1:]
	a.rootCommand = NewCommand()
	a.currentCommand = a.rootCommand
	a.arguments = nil
	a.Stdout = printf
	a.Stderr = printf
	a.PanicInsteadOfExit = false
	a.ExtraFlagsAllowed = false
	a.ExtraPositionalsAllowed = false
	a.RepeatedFlagsAllowed = false
}

func (a *App) RootCommand() *Command    { return a.rootCommand }
func (a *App) CurrentCommand() *Command { return a.currentCommand }
func (a *App) Arguments() *Arguments    { return a.arguments }
func (a *App) IsParsed() bool           { return a.arguments != nil }
func (a *App) Exit(code int) {
	if a.PanicInsteadOfExit {
		panic(code)
	}
	os.Exit(code)
}

// Parse is called for every command in the path.
func (a *App) Parse() {
	// Check if already parsed
	if a.IsParsed() {
		return
	}

	a.initialize()

	// Check new subcommand
	if len(a.queue) > 0 {
		next := a.queue[0]

		// There is a subcommand, so we execute it
		for _, cmd := range a.currentCommand.subcommands {
			if cmd.name == next {
				// Prepare the state for the subcommand
				a.currentCommand = cmd
				a.queue = a.queue[1:]
				a.path = append(a.path, cmd.name)

				// Pass the execution to the subcommand
				cmd.execute()

				// Exit as the first command fully executes, interrupting the flow of
				// the parent command, i.e., if there is a subcommand, the parent
				// command will not execute after Parse
				a.Exit(0)
			}
		}
	}

	// Parse the flags and positionals of the stack
	args, err := parseArguments(a)
	if err != nil {
		a.Stderr(err.Error())
		a.Exit(1)
	}
	a.arguments = args
}

// ParseArgs parse the given arguments instead of os.Args. This is useful for
// testing and edge cases. The arguments should not include the program name.
func (a *App) ParseArgs(args []string) {
	a.queue = args
	a.Parse()
}

// ShowHelp prints the help message for the current command, including its description,
// usage, and available flags and subcommands.
func (a *App) ShowHelp() {
	s := a.GetHelpString()
	a.Stdout(s)
}

func (a *App) GetHelpString() string {
	a.initialize()
	name := strings.Join(a.path, " ")

	cmds := ""
	cmd := a.CurrentCommand()
	if len(cmd.subcommands) > 0 {
		cmds = " <command>"
	}

	opts := ""
	if len(cmd.flags) > 0 {
		opts = " [options]"
	}

	positionals := ""
	for _, p := range cmd.positionals {
		if p.IsRequired() {
			positionals += " <" + p.Name() + ">"
			continue
		}
		positionals += " [<" + p.Name() + ">]"
	}

	writer := NewDefaultTypewriter()
	writer.WriteLine("Usage: %s%s%s%s", name, cmds, opts, positionals)
	if cmd.description != "" {
		writer.WriteLine("\n%s", cmd.description)
	}

	if len(cmd.flags) > 0 {
		writer.WriteLine("")
		writer.WriteLine("Options:")
		for _, f := range cmd.flags {
			opts := f.Signature()
			desc := f.Description()
			req := ""
			if f.IsRequired() {
				req = "(required) "
			} else if f.HasDefault() {
				req = fmt.Sprintf("(default=%v) ", f.RawDefault())
			}

			writer.WriteLine("  %s\t%s%s", opts, req, desc)
		}
	}

	if len(cmd.subcommands) > 0 {
		writer.WriteLine("")
		writer.WriteLine("Commands:")
		for _, cmd := range cmd.subcommands {
			writer.WriteLine("  %s\t%s", cmd.name, cmd.shortDescription)
		}
	}

	if len(cmd.positionals) > 0 {
		writer.WriteLine("")
		writer.WriteLine("Positionals:")
		for _, p := range cmd.positionals {
			desc := p.Description()
			req := ""
			if p.IsRequired() {
				req = "(required) "
			}
			writer.WriteLine("  %s\t%s%s", p.Name(), req, desc)
		}
	}

	return writer.Flush()
}

func (a *App) initialize() {
	if len(a.path) == 0 {
		if a.rootCommand.name != "" {
			a.path = append(a.path, a.rootCommand.name)
		} else {
			exec := os.Args[0]
			name := path.Base(exec)
			ext := path.Ext(name)
			if ext != "" {
				name = name[:len(name)-len(ext)]
			}
			a.path = append(a.path, name)
		}
	}
}
