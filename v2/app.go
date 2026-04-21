package v2

import (
	"fmt"
	"os"
)

func printf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

type App struct {
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

	// Check new subcommand
	if len(a.queue) > 0 {
		next := a.queue[0]

		// There is a subcommand, so we execute it
		for _, cmd := range a.currentCommand.subcommands {
			if cmd.name == next {
				// Prepare the state for the subcommand
				a.currentCommand = cmd
				a.queue = a.queue[1:]

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
		// TODO: print the error in a better way, e.g., with colors and formatting
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
