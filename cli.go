package cli

import (
	"io"
	"strings"
	"time"

	internal "github.com/renatopp/go-cli/internal"
	"github.com/renatopp/go-cli/locales"
)

var app *internal.App

type TFlag[T any] = *internal.GenericFlag[T]
type TPositional[T any] = *internal.GenericPositional[T]

// Initialize global state
func init() {
	app = internal.NewApp()
}

// New creates a new App instance with default settings.
func New() *internal.App {
	return internal.NewApp()
}

// Clear resets the state of the CLI, allowing users to define a new command
// structure and configuration. This can be useful in scenarios where you want to
// reuse the same CLI instance for different commands or when you want to reset
// the CLI state after executing a command.
func Clear() {
	app.Clear()
}

// SetLocale replaces the strings used for help text and error messages
// across the CLI, allowing you to localize go-cli's built-in output. It
// applies globally, independent of which App instance is used. The built-in
// locales and the Locale type live in the locales package, e.g.:
//
//	cli.SetLocale(locales.PTBR())
func SetLocale(locale locales.Locale) {
	internal.SetLocale(locale)
}

// Name sets the name for the current command. The name is used in help text
// to identify the command and its usage. Use only its immediate name (e.g.
// "version" instead of "app version") since the command hierarchy is
// automatically handled by go-cli.
func Name(n string) { app.Name(n) }

// Description sets the description for the current command. Descriptions are
// used in help text to provide more information about the command and its
// purpose.
func Description(d string) { app.Description(d) }

// Version sets the version for the CLI. This enables the --version flag for
// the root command.
func Version(v string) { app.Version(v) }

// StdoutWith allows you to specify a custom io.Writer for handling standard
// output. This can be useful for redirecting output to a file, logging system,
// or for testing purposes. It is used to print the help text.
func StdoutWith(w io.Writer) {
	app.StdoutWith(w)
}

// StderrWith allows you to specify a custom io.Writer for handling standard error
// output. This can be useful for redirecting error messages to a file, logging
// system, or for testing purposes. It is used to print error messages.
func StderrWith(w io.Writer) {
	app.StderrWith(w)
}

// Print prints a formatted message using the stdout function.
func Print(format string, v ...any) { app.Print(format, v...) }

// Error prints a formatted error message using the stderr function.
func Error(format string, v ...any) { app.Error(format, v...) }

// Fatal prints a formatted error message using the stderr function and then
// exits with code 1.
func Fatal(format string, v ...any) { app.Fatal(format, v...) }

// FatalIf checks if the provided error is not nil, and if so, it prints the error
// message using the stderr function and then exits with code 1.
func FatalIf(err error) { app.FatalIf(err) }

// CurrentCommand returns the current command being executed.
func CurrentCommand() *internal.Command { return app.CurrentCommand() }

// Arguments returns the parsed arguments for the current command. It will be
// nil if the arguments have not been parsed yet.
func Arguments() *internal.Arguments { return app.Arguments() }

// NArgs returns the number of positional arguments provided by the user.
// Should be used only after Parse() is called, otherwise it will return 0.
func NArgs() int {
	return app.NArgs()
}

// Arg retrieves the value of a positional argument by its index.
// Should be used only after Parse() is called, otherwise it will return an
// empty string.
func Arg(index int) string {
	return app.Arg(index)
}

// Args retrieves all positional arguments provided by the user.
// Should be used only after Parse() is called, otherwise it will return an
// empty slice.
func Args() []string {
	return app.Args()
}

// NExtraArgs returns the number of extra positional arguments provided by the user,
// i.e., those that are not defined in the command. Should be used only after
// Parse() is called, otherwise it will return 0.
func NExtraArgs() int {
	return app.NExtraArgs()
}

// ExtraArg retrieves the value of an extra positional argument by its index, i.e.,
// those that are not defined in the command. Should be used only after Parse() is
// called, otherwise it will return an empty string.
func ExtraArg(index int) string {
	return app.ExtraArg(index)
}

// ExtraArgs retrieves all extra positional arguments provided by the user, i.e.,
// those that are not defined in the command. Should be used only after Parse() is
// called, otherwise it will return an empty slice.
func ExtraArgs() []string {
	return app.ExtraArgs()
}

// SetArgs sets the arguments for the app.
func SetArgs(args []string) { app.SetArgs(args) }

// UsePanicInsteadOfExit configures the CLI to panic instead of exiting when
// an error  occurs or when a command finishes execution. This can be useful
// for testing purposes or customization of the cli behavior. The panic will
// be called with the exit code as the argument.
func UsePanicInsteadOfExit(usePanic bool) {
	app.UsePanicInsteadOfExit(usePanic)
}

// AllowExtraPositionals configures the CLI to allow extra positional arguments
// that are not defined in the command.  You may use variadic positional for
// this purpose as well. By default, extra positional arguments are not allowed.
//
// Extra positional arguments can be accessed using the `Arg` function or the
// `ExtraArg` function.
func AllowExtraPositionals(allow bool) {
	app.AllowExtraPositionals(allow)
}

// AllowExtraFlags configures the CLI to allow extra flags that are not defined
// in the command. By default, extra flags are not allowed.
func AllowExtraFlags(allow bool) {
	app.AllowExtraFlags(allow)
}

// AllowRepeatedFlags configures the CLI to allow repeated flags. If set to true,
// the CLI will not return an error if a flag is provided multiple times. Instead,
// the last value provided for the flag will be used. If set to false, the CLI
// will return an error if a flag is provided multiple times. By default, repeated
// flags are not allowed.
func AllowRepeatedFlags(allow bool) {
	app.AllowRepeatedFlags(allow)
}

// Exit terminates the program with the provided exit code. An exit code of 0
// typically indicates successful execution, while a non-zero exit code
// indicates an error or abnormal termination.
func Exit(code int) {
	app.Exit(code)
}

// ShowHelp prints the help message for the current command, including its description,
// usage, and available flags and subcommands. This function uses the Stdout
// function to output the help message.
func ShowHelp() {
	app.ShowHelp()
}

// HelpString returns the help message for the current command as a string.
func HelpString() string {
	return app.HelpString()
}

// AutoHelp configures the CLI to automatically show the help message when the user
// provides the `-h` or `--help` flag. By default, auto help is disabled.
func AutoHelp(enabled bool) {
	app.AutoHelp(enabled)
}

// IsParsed returns true if the arguments have been parsed
// successfully.
func IsParsed() bool {
	return app.IsParsed()
}

// Parse processes the command-line arguments provided by the user and executes
// the appropriate command based on the defined command structure. This function
// should be called after all commands, flags, and positional arguments have
// been defined.
//
// Subcommands are executed based on the first argument that matches a defined
// name, interrupting the execution of code after the Parse() call on the
// parent commands.
func Parse() {
	app.Parse()
}

// ParseArgs is similar to Parse but allows you to specify a custom slice of
// arguments to parse instead of using the default os.Args. This can be useful for
// testing purposes or when you want to parse a specific set of arguments without
// relying on the command-line input.
//
// DO NOT PROVIDE the program name (i.e., os.Args[0]) in the args slice.
func ParseArgs(args []string) {
	app.ParseArgs(args)
}

// Command creates a new command with the specified name, short description, and
// execution function. The command is added as a subcommand to the current command.
// The execute function will be called when the command is invoked by the user.
func Command(name string, shortDescription string, execute func()) *internal.Command {
	return app.Command(name, shortDescription, execute)
}

// Cmd is an alias for Command, providing a shorter name for creating commands.
// It creates a new command with the specified name, short description, and
// execution function. The command is added as a subcommand to the current
// command.
func Cmd(name string, shortDescription string, execute func()) *internal.Command {
	return app.Cmd(name, shortDescription, execute)
}

// PosFunc creates a positional argument of any type T, using the provided
// parser function to convert the user input string into T. Use this to define
// custom positional types, e.g.:
//
//	level := cli.PosFunc("level", "The log level.", ParseLevel)
func PosFunc[T any](name, description string, parser func(string) (T, error)) TPositional[T] {
	return internal.PosFunc(app, name, description, parser)
}

func Pos(name, description string) TPositional[string] {
	return app.Pos(name, description)
}
func PosString(name, description string) TPositional[string] {
	return app.PosString(name, description)
}
func PosInt(name, description string) TPositional[int] {
	return app.PosInt(name, description)
}
func PosUint(name, description string) TPositional[uint] {
	return app.PosUint(name, description)
}
func PosFloat(name, description string) TPositional[float64] {
	return app.PosFloat(name, description)
}
func PosBool(name, description string) TPositional[bool] {
	return app.PosBool(name, description)
}
func PosDuration(name, description string) TPositional[time.Duration] {
	return app.PosDuration(name, description)
}

// FlagFunc creates a flag of any type T, using the provided parser function to
// convert the user input string into T. Use this to define custom flag types,
// e.g.:
//
//	level := cli.FlagFunc("level", "l", "The log level.", ParseLevel)
func FlagFunc[T any](long, short, description string, parser func(string) (T, error)) TFlag[T] {
	return internal.FlagFunc(app, long, short, description, parser)
}

func Flag(long, short, description string) TFlag[string] {
	return app.Flag(long, short, description)
}
func FlagString(long, short, description string) TFlag[string] {
	return app.FlagString(long, short, description)
}
func FlagInt(long, short, description string) TFlag[int] {
	return app.FlagInt(long, short, description)
}
func FlagUint(long, short, description string) TFlag[uint] {
	return app.FlagUint(long, short, description)
}
func FlagFloat(long, short, description string) TFlag[float64] {
	return app.FlagFloat(long, short, description)
}
func FlagBool(long, short, description string) TFlag[bool] {
	return app.FlagBool(long, short, description)
}
func FlagDuration(long, short, description string) TFlag[time.Duration] {
	return app.FlagDuration(long, short, description)
}

// GetFlag retrieves a flag by its long or short name and attempts to cast it to the specified type T.
// If the flag is not found or cannot be cast to the desired type, an error is returned.
func GetFlag[T internal.Flag](longOrShort string) (T, error) {
	f, err := app.GetFlag(longOrShort)
	if err != nil {
		var zero T
		return zero, err
	}

	typed, ok := f.(T)
	if !ok {
		var zero T
		return zero, internal.ErrFlagNotType
	}

	return typed, nil
}

// CheckExclusiveFlags checks that at most one of the provided flags is passed.
// This function should be called after Parse().
func CheckExclusiveFlags(flags ...internal.Flag) {
	parsedFlags := []internal.Flag{}
	for _, flag := range flags {
		if flag.IsParsed() {
			parsedFlags = append(parsedFlags, flag)
		}
	}

	if len(parsedFlags) > 1 {
		flagNames := []string{}
		for _, flag := range parsedFlags {
			flagNames = append(flagNames, flag.Signature())
		}
		app.Error("mutually exclusive flags provided: %s", strings.Join(flagNames, " and "))
		app.Exit(1)
	}
}

// CheckAnyFlag checks that at least one of the provided flags is passed. This
// function should be called after Parse().
func CheckAnyFlag(flags ...internal.Flag) {
	for _, flag := range flags {
		if flag.IsParsed() {
			return
		}
	}

	flagNames := []string{}
	for _, flag := range flags {
		flagNames = append(flagNames, flag.Signature())
	}
	app.Error("at least one of the following flags must be provided: %s", strings.Join(flagNames, " or "))
	app.Exit(1)
}
