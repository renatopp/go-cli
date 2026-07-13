package pkg

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/renatopp/go-cli/pkg/locales"
	"github.com/renatopp/go-cli/pkg/parsers"
)

// HelpFormatter converts a command into its help text. Set a custom one with
// App.HelpFormatter to fully control how help is rendered.
type HelpFormatter func(cmd *Command, loc Locale) string

// ErrorFormatter converts an error into the message printed to stderr. Set a
// custom one with App.ErrorFormatter to fully control how errors are rendered.
// Parsing errors are of type *errors.CliError (see the errors package), so
// formatters can inspect them with errors.As and use Locale.LocalizedError to
// render locale-specific messages.
type ErrorFormatter func(err error, loc Locale) string

// Locale is re-exported here for convenience of the internal package.
type Locale = locales.Locale

type App struct {
	path                    []string // the path of commands leading to the current command, e.g., ["git", "commit"]
	queue                   []string // the queue of arguments to be parsed
	rootCommand             *Command
	currentCommand          *Command
	arguments               *Arguments // parsed arguments
	stdout                  io.Writer
	stderr                  io.Writer
	helpFormatter           HelpFormatter
	errorFormatter          ErrorFormatter
	locale                  Locale
	panicInsteadOfExit      bool
	extraFlagsAllowed       bool
	extraPositionalsAllowed bool
	repeatedFlagsAllowed    bool
	autoHelp                bool
	version                 string
}

// NewApp creates a new App instance with default settings.
func NewApp(hf HelpFormatter, ef ErrorFormatter) *App {
	s := &App{
		helpFormatter:  hf,
		errorFormatter: ef,
	}
	s.Clear()
	return s
}

// WithVersion sets the version for the CLI. This enables the --version flag for
// the root command.
func (a *App) WithVersion(v string) *App { a.version = v; return a }

// WithStdout allows you to specify a custom io.Writer for handling standard
// output. This can be useful for redirecting output to a file, logging system,
// or for testing purposes. It is used to print the help text.
func (a *App) WithStdout(w io.Writer) *App { a.stdout = w; return a }

// WithStderr allows you to specify a custom io.Writer for handling standard error
// output. This can be useful for redirecting error messages to a file, logging
// system, or for testing purposes. It is used to print error messages.
func (a *App) WithStderr(w io.Writer) *App { a.stderr = w; return a }

// HelpFormatter replaces the function used to render the help message for
// a command. The default is DefaultHelpFormatter.
func (a *App) WithHelpFormatter(f HelpFormatter) *App { a.helpFormatter = f; return a }

// ErrorFormatter replaces the function used to render error messages
// before they are written to stderr. The default is DefaultErrorFormatter.
// Parsing errors are typed (e.g. *UnknownFlagError), so the formatter can
// inspect them with errors.As.
func (a *App) WithErrorFormatter(f ErrorFormatter) *App { a.errorFormatter = f; return a }

// WithLocale replaces the active locale used for this App's help text and error
// messages. Any field left as the zero value ("") falls back to the default
// English text, so callers can override only the strings they want to
// translate.
func (a *App) WithLocale(l Locale) *App { a.locale = l; return a }

// WithArgs allows you to specify a custom set of arguments for the CLI to parse,
// instead of using os.Args. This is useful for testing and edge cases. The
// arguments should NOT INCLUDE the program name.
func (a *App) WithArgs(args []string) *App {
	a.queue = args
	return a
}

// WithAutoHelp configures the CLI to automatically show the help message when the user
// provides the `-h` or `--help` flag. By default, auto help is disabled.
func (a *App) WithAutoHelp(enabled bool) {
	a.autoHelp = enabled
}

// UsePanic configures the CLI to panic instead of exiting when
// an error occurs or when a command finishes execution. This can be useful
// for testing purposes or customization of the cli behavior. The panic will
// be called with the exit code as the argument.
func (a *App) UsePanic(usePanic bool) {
	a.panicInsteadOfExit = usePanic
}

// AllowExtraPositionals configures the CLI to allow extra positional arguments
// that are not defined in the command.  You may use variadic positional for
// this purpose as well. By default, extra positional arguments are not allowed.
//
// Extra positional arguments can be accessed using the `Arg` function or the
// `ExtraArg` function.
func (a *App) AllowExtraPositionals(allow bool) {
	a.extraPositionalsAllowed = allow
}

// AllowExtraFlags configures the CLI to allow extra flags that are not defined
// in the command. By default, extra flags are not allowed.
func (a *App) AllowExtraFlags(allow bool) {
	a.extraFlagsAllowed = allow
}

// AllowRepeatedFlags configures the CLI to allow repeated flags. If set to true,
// the CLI will not return an error if a flag is provided multiple times. Instead,
// the last value provided for the flag will be used. If set to false, the CLI
// will return an error if a flag is provided multiple times. By default, repeated
// flags are not allowed.
func (a *App) AllowRepeatedFlags(allow bool) {
	a.repeatedFlagsAllowed = allow
}

// RootCommand returns the root command of the CLI, which is the top-level
// command that all other subcommands are attached to. You can use this to define your commands and flags.
func (a *App) RootCommand() *Command { return a.rootCommand }

// CurrentCommand returns the current command being executed, which is the last
// command in the path. It will be the root command if no subcommand has been
// executed yet.
func (a *App) CurrentCommand() *Command { return a.currentCommand }

// Locale returns this App's currently active locale.
func (a *App) Locale() Locale { return a.locale }

// IsParsed returns true if the arguments have been parsed
// successfully.
func (a *App) IsParsed() bool { return a.arguments != nil }

// GetHelp generates and returns the help message string for the current
// command using the help formatter.
func (a *App) GetHelp() string {
	a.initialize()
	return a.helpFormatter(a.CurrentCommand(), a.locale)
}

// Clear resets the state of the CLI, allowing you to define a new set of
// commands and flags. This is useful for testing or if you want to reuse the
// same App instance for different command configurations. It clears the command
// path, argument queue, root command, current command, and parsed arguments,
// and resets the output functions and configuration options to their default values.
func (a *App) Clear() {
	a.path = []string{}
	a.queue = os.Args[1:]
	a.rootCommand = NewCommand(nil)
	a.currentCommand = a.rootCommand
	a.arguments = nil
	a.stdout = os.Stdout
	a.stderr = os.Stderr
	a.locale = locales.EN()
	a.panicInsteadOfExit = false
	a.extraFlagsAllowed = false
	a.extraPositionalsAllowed = false
	a.repeatedFlagsAllowed = false
	a.autoHelp = false
	a.version = ""
}

// Fatal formats an error message, renders it using the error formatter,
// writes it to the stderr writer and then exits with code 1.
func (a *App) Fatal(format string, v ...any) {
	err := fmt.Errorf(format, v...)
	fmt.Fprintf(a.stderr, "%s\n", a.errorFormatter(err, a.locale))
	a.Exit(1)
}

// FatalIf checks if the provided error is not nil, and if so, it renders the
// error using the error formatter, writes it to the stderr writer and then
// exits with code 1. Unlike Fatal, it passes the error through unchanged, so
// the error formatter can inspect its concrete type (e.g. with errors.As).
func (a *App) FatalIf(err error) {
	if err != nil {
		fmt.Fprintf(a.stderr, "%s\n", a.errorFormatter(err, a.locale))
		a.Exit(1)
	}
}

// Exit terminates the program with the given exit code. If
// PanicInsteadOfExit is true, it panics with the exit code instead of exiting,
// which can be useful for testing.
func (a *App) Exit(code int) {
	if a.panicInsteadOfExit {
		panic(code)
	}
	os.Exit(code)
}

// Help prints the help message for the current command, including its description,
// usage, and available flags and subcommands, using the help formatter.
func (a *App) Help() {
	fmt.Fprintf(a.stdout, "%s\n", a.GetHelp())
}

// Parse is called for every command in the path.
//
// This function is responsible for traversing the command hierarchy recursively.
// If a subcommand is found, it will be executed and the parent command will not
// continue execution. If no subcommand is found, it will parse the flags and
// positionals for the current command.
func (a *App) Parse() *Arguments {
	// Check if already parsed
	if a.IsParsed() {
		return a.arguments
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
				cmd.inheritFlags()
				cmd.execute()

				// Exit as the first command fully executes, interrupting the flow of
				// the parent command, i.e., if there is a subcommand, the parent
				// command will not execute after Parse
				a.Exit(0)
			}
		}
	}

	// There is no match with any subcommand, so this command will execute.
	// Parse the flags and positionals of the stack
	args, err := parseArguments(a)
	a.FatalIf(err)
	a.arguments = args
	return args
}

// ParseArgs parse the given arguments instead of os.Args. This is useful for
// testing and edge cases. The arguments should NOT INCLUDE the program name.
func (a *App) ParseArgs(args []string) *Arguments {
	a.queue = args
	return a.Parse()
}

func (a *App) initialize() {
	rootCmd := a.rootCommand
	curCmd := a.currentCommand

	if a.autoHelp && (!curCmd.HasFlag("help") || !curCmd.HasFlag("h")) {
		helpFlag := NewGenericFlag("help", "h", a.locale.HelpFlagDescription, parsers.Bool)
		curCmd.WithFlag(helpFlag)
	}

	if a.version != "" && curCmd == rootCmd && (!rootCmd.HasFlag("version") || !rootCmd.HasFlag("v")) {
		versionFlag := NewGenericFlag("version", "v", a.locale.VersionFlagDescription, parsers.Bool)
		rootCmd.WithFlag(versionFlag)
	}

	if len(a.path) == 0 {
		if rootCmd.name == "" {
			exec := os.Args[0]
			name := path.Base(exec)
			ext := path.Ext(name)
			if ext != "" {
				name = name[:len(name)-len(ext)]
			}
			// Set the resolved name on the root command so that helpers like
			// Command.Path() report the executable name.
			rootCmd.name = name
		}
		a.path = append(a.path, rootCmd.name)
	}
}
