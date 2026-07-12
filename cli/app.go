package cli

import (
	"fmt"
	"github.com/renatopp/go-cli/cli/parsers"
	"io"
	"os"
	"path"
	"time"
)

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
	panicInsteadOfExit      bool
	extraFlagsAllowed       bool
	extraPositionalsAllowed bool
	repeatedFlagsAllowed    bool
	autoHelp                bool
	version                 string
}

// NewApp creates a new App instance with default settings.
func NewApp() *App {
	s := &App{}
	s.Clear()
	return s
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
	a.helpFormatter = DefaultHelpFormatter
	a.errorFormatter = DefaultErrorFormatter
	a.panicInsteadOfExit = false
	a.extraFlagsAllowed = false
	a.extraPositionalsAllowed = false
	a.repeatedFlagsAllowed = false
	a.autoHelp = false
	a.version = ""
}

// RootCommand returns the root command of the CLI, which is the top-level
// command that all other subcommands are attached to. You can use this to define your commands and flags.
func (a *App) RootCommand() *Command { return a.rootCommand }

// CurrentCommand returns the current command being executed, which is the last
// command in the path. It will be the root command if no subcommand has been
// executed yet.
func (a *App) CurrentCommand() *Command { return a.currentCommand }

// Name sets the name for the current command. The name is used in help text
// to identify the command and its usage. Use only its immediate name (e.g.
// "version" instead of "app version") since the command hierarchy is
// automatically handled by go-cli.
func (a *App) Name(n string) { a.CurrentCommand().WithName(n) }

// Description sets the description for the current command. Descriptions are
// used in help text to provide more information about the command and its
// purpose.
func (a *App) Description(d string) { a.CurrentCommand().WithDescription(d) }

// Version sets the version for the CLI. This enables the --version flag for
// the root command.
func (a *App) Version(v string) { a.version = v }

// SetStdout allows you to specify a custom io.Writer for handling standard
// output. This can be useful for redirecting output to a file, logging system,
// or for testing purposes. It is used to print the help text.
func (a *App) SetStdout(w io.Writer) {
	a.stdout = w
}

// SetStderr allows you to specify a custom io.Writer for handling standard error
// output. This can be useful for redirecting error messages to a file, logging
// system, or for testing purposes. It is used to print error messages.
func (a *App) SetStderr(w io.Writer) {
	a.stderr = w
}

// SetHelpFormatter replaces the function used to render the help message for
// a command. The default is DefaultHelpFormatter.
func (a *App) SetHelpFormatter(f HelpFormatter) {
	a.helpFormatter = f
}

// SetErrorFormatter replaces the function used to render error messages
// before they are written to stderr. The default is DefaultErrorFormatter.
// Parsing errors are typed (e.g. *UnknownFlagError), so the formatter can
// inspect them with errors.As.
func (a *App) SetErrorFormatter(f ErrorFormatter) {
	a.errorFormatter = f
}

// Fail renders the given error using the error formatter, writes it to the
// stderr writer and exits with code 1.
func (a *App) Fail(err error) {
	fmt.Fprintf(a.stderr, "%s\n", a.errorFormatter(err))
	a.Exit(1)
}

// Fatal formats an error message, renders it using the error formatter,
// writes it to the stderr writer and then exits with code 1.
func (a *App) Fatal(format string, v ...any) {
	a.Fail(fmt.Errorf(format, v...))
}

// FatalIf checks if the provided error is not nil, and if so, it renders the
// error using the error formatter, writes it to the stderr writer and then
// exits with code 1.
func (a *App) FatalIf(err error) {
	if err != nil {
		a.Fail(err)
	}
}

// Arguments returns the parsed arguments for the current command. It will be
// nil if the arguments have not been parsed yet.
func (a *App) Arguments() *Arguments { return a.arguments }

// NArgs returns the number of positional arguments provided by the user.
// Should be used only after Parse() is called, otherwise it will return 0.
func (a *App) NArgs() int {
	if !a.IsParsed() {
		return 0
	}
	return len(a.Arguments().Args)
}

// Arg retrieves the value of a positional argument by its index.
// Should be used only after Parse() is called, otherwise it will return an
// empty string.
func (a *App) Arg(index int) string {
	if !a.IsParsed() {
		return ""
	}
	args := a.Arguments().Args
	if index < 0 || index >= len(args) {
		return ""
	}
	return args[index]
}

// Args retrieves all positional arguments provided by the user.
// Should be used only after Parse() is called, otherwise it will return an
// empty slice.
func (a *App) Args() []string {
	if !a.IsParsed() {
		return []string{}
	}
	return a.Arguments().Args
}

// NExtraArgs returns the number of extra positional arguments provided by the user,
// i.e., those that are not defined in the command. Should be used only after
// Parse() is called, otherwise it will return 0.
func (a *App) NExtraArgs() int {
	if !a.IsParsed() {
		return 0
	}
	return len(a.Arguments().ExtraArgs)
}

// ExtraArg retrieves the value of an extra positional argument by its index, i.e.,
// those that are not defined in the command. Should be used only after Parse() is
// called, otherwise it will return an empty string.
func (a *App) ExtraArg(index int) string {
	if !a.IsParsed() {
		return ""
	}
	extraArgs := a.Arguments().ExtraArgs
	if index < 0 || index >= len(extraArgs) {
		return ""
	}
	return extraArgs[index]
}

// ExtraArgs retrieves all extra positional arguments provided by the user, i.e.,
// those that are not defined in the command. Should be used only after Parse() is
// called, otherwise it will return an empty slice.
func (a *App) ExtraArgs() []string {
	if !a.IsParsed() {
		return []string{}
	}
	return a.Arguments().ExtraArgs
}

// SetArgs sets the arguments for the app.
func (a *App) SetArgs(args []string) {
	a.queue = args
}

// UsePanicInsteadOfExit configures the CLI to panic instead of exiting when
// an error  occurs or when a command finishes execution. This can be useful
// for testing purposes or customization of the cli behavior. The panic will
// be called with the exit code as the argument.
func (a *App) UsePanicInsteadOfExit(usePanic bool) {
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

// Exit terminates the program with the given exit code. If
// PanicInsteadOfExit is true, it panics with the exit code instead of exiting,
// which can be useful for testing.
func (a *App) Exit(code int) {
	if a.panicInsteadOfExit {
		panic(code)
	}
	os.Exit(code)
}

// ShowHelp prints the help message for the current command, including its description,
// usage, and available flags and subcommands, using the help formatter.
func (a *App) ShowHelp() {
	fmt.Fprintf(a.stdout, "%s\n", a.HelpString())
}

// HelpString generates and returns the help message string for the current
// command using the help formatter.
func (a *App) HelpString() string {
	a.initialize()
	return a.helpFormatter(a.CurrentCommand())
}

// AutoHelp configures the CLI to automatically show the help message when the user
// provides the `-h` or `--help` flag. By default, auto help is disabled.
func (a *App) AutoHelp(enabled bool) {
	a.autoHelp = enabled
}

// IsParsed returns true if the arguments have been parsed
// successfully.
func (a *App) IsParsed() bool { return a.arguments != nil }

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
				cmd.inheritFlags()
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
		a.Fail(err)
	}
	a.arguments = args
}

// ParseArgs parse the given arguments instead of os.Args. This is useful for
// testing and edge cases. The arguments should not include the program name.
func (a *App) ParseArgs(args []string) {
	a.queue = args
	a.Parse()
}

// Command creates a new command with the specified name, short description, and
// execution function. The command is added as a subcommand to the current command.
// The execute function will be called when the command is invoked by the user.
func (a *App) Command(name string, shortDescription string, execute func()) *Command {
	cmd := NewCommand(a.CurrentCommand()).
		WithName(name).
		WithShortDescription(shortDescription).
		WithDescription(shortDescription).
		WithExecute(execute)

	a.CurrentCommand().
		WithSubcommand(cmd)

	return cmd
}

// Cmd is an alias for Command, providing a shorter name for creating commands.
// It creates a new command with the specified name, short description, and
// execution function. The command is added as a subcommand to the current
// command.
func (a *App) Cmd(name string, shortDescription string, execute func()) *Command {
	return a.Command(name, shortDescription, execute)
}

func (a *App) Pos(name, description string) *GenericPositional[string] {
	return _addpos(a, NewGenericPositional(name, description, parsers.String))
}
func (a *App) PosString(name, description string) *GenericPositional[string] {
	return _addpos(a, NewGenericPositional(name, description, parsers.String))
}
func (a *App) PosInt(name, description string) *GenericPositional[int] {
	return _addpos(a, NewGenericPositional(name, description, parsers.Int[int]))
}
func (a *App) PosUint(name, description string) *GenericPositional[uint] {
	return _addpos(a, NewGenericPositional(name, description, parsers.Uint[uint]))
}
func (a *App) PosFloat(name, description string) *GenericPositional[float64] {
	return _addpos(a, NewGenericPositional(name, description, parsers.Float[float64]))
}
func (a *App) PosBool(name, description string) *GenericPositional[bool] {
	return _addpos(a, NewGenericPositional(name, description, parsers.Bool))
}
func (a *App) PosDuration(name, description string) *GenericPositional[time.Duration] {
	return _addpos(a, NewGenericPositional(name, description, parsers.Duration))
}

func (a *App) Flag(long, short, description string) *GenericFlag[string] {
	return _addflag(a, NewGenericFlag(long, short, description, parsers.String))
}
func (a *App) FlagString(long, short, description string) *GenericFlag[string] {
	return _addflag(a, NewGenericFlag(long, short, description, parsers.String))
}
func (a *App) FlagInt(long, short, description string) *GenericFlag[int] {
	return _addflag(a, NewGenericFlag(long, short, description, parsers.Int[int]))
}
func (a *App) FlagUint(long, short, description string) *GenericFlag[uint] {
	return _addflag(a, NewGenericFlag(long, short, description, parsers.Uint[uint]))
}
func (a *App) FlagFloat(long, short, description string) *GenericFlag[float64] {
	return _addflag(a, NewGenericFlag(long, short, description, parsers.Float[float64]))
}
func (a *App) FlagBool(long, short, description string) *GenericFlag[bool] {
	return _addflag(a, NewGenericFlag(long, short, description, parsers.Bool))
}
func (a *App) FlagDuration(long, short, description string) *GenericFlag[time.Duration] {
	return _addflag(a, NewGenericFlag(long, short, description, parsers.Duration))
}

func (a *App) GetFlag(longOrShort string) (Flag, error) {
	return a.CurrentCommand().GetFlag(longOrShort)
}

func (a *App) initialize() {
	rootCmd := a.rootCommand
	curCmd := a.currentCommand

	if a.autoHelp && (!curCmd.HasFlag("help") || !curCmd.HasFlag("h")) {
		helpFlag := NewGenericFlag("help", "h", GetLocale().HelpFlagDescription, parsers.Bool)
		curCmd.WithFlag(helpFlag)
	}

	if a.version != "" && curCmd == rootCmd && (!rootCmd.HasFlag("version") || !rootCmd.HasFlag("v")) {
		versionFlag := NewGenericFlag("version", "v", GetLocale().VersionFlagDescription, parsers.Bool)
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

func _addpos[T Positional](a *App, p T) T {
	a.CurrentCommand().WithPositional(p)
	return p
}

func _addflag[T Flag](a *App, f T) T {
	a.CurrentCommand().WithFlag(f)
	return f
}
