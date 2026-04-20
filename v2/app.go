package v2

import "os"

type App struct {
	queue          []string // the queue of arguments to be parsed
	rootCommand    *Command
	currentCommand *Command
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
}

func (a *App) RootCommand() *Command    { return a.rootCommand }
func (a *App) CurrentCommand() *Command { return a.currentCommand }

// func (a *App) Arguments() *Arguments    { return a.arguments }
// func (a *App) IsParsed() bool           { return a.arguments != nil }
func (a *App) Exit(code int) { os.Exit(code) }

// Parse is called for every command in the path.
func (a *App) Parse() {
	// Check if already parsed
	// if a.IsParsed() {
	// 	return
	// }

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
				os.Exit(0)
			}
		}
	}

	// Parse the flags and positionals of the stack
	// args, err := parseArguments(a.queue, a)
	parseArguments(a)
}
