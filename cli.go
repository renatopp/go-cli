package cli

import (
	"fmt"
	"io"
	"time"

	"github.com/renatopp/go-cli/core"
	"github.com/renatopp/go-cli/errors"
	"github.com/renatopp/go-cli/formatters"
	"github.com/renatopp/go-cli/locales"
	"github.com/renatopp/go-cli/parsers"
)

var app *core.App

// Initialize global state
func init() {
	app = core.NewApp(formatters.DefaultHelpFormatter, formatters.DefaultErrorFormatter)
}

func Clear() { app.Clear() }

// Name sets the name for the current command. The name is used in help text
// to identify the command and its usage. Use only its immediate name (e.g.
// "version" instead of "app version") since the command hierarchy is
// automatically handled by go-cli.
func Name(n string) { app.CurrentCommand().WithName(n) }

// Description sets the description for the current command. Descriptions are
// used in help text to provide more information about the command and its
// purpose.
func Description(d string) { app.CurrentCommand().WithDescription(d) }

// Version sets the version for the CLI. This enables the --version flag for
// the root command.
func Version(v string) { app.WithVersion(v) }

// Locale replaces the strings used for help text and error messages
// across the CLI, allowing you to localize go-cli's built-in output. It
// applies globally, independent of which App instance is used. The built-in
// locales and the Locale type live in the locales package, e.g.:
//
//	cli.Locale(locales.PT_BR())
func Locale(locale locales.Locale) {
	app.WithLocale(locale)
}

// Stdout allows you to specify a custom io.Writer for handling standard
// output. This can be useful for redirecting output to a file, logging system,
// or for testing purposes. It is used to print the help text.
func Stdout(w io.Writer) {
	app.WithStdout(w)
}

// Stderr allows you to specify a custom io.Writer for handling standard error
// output. This can be useful for redirecting error messages to a file, logging
// system, or for testing purposes. It is used to print error messages.
func Stderr(w io.Writer) {
	app.WithStderr(w)
}

// HelpFormatter replaces the function used to render the help message of a
// command, allowing you to fully customize the help style. The default is
// DefaultHelpFormatter, which can also be wrapped, e.g.:
//
//	cli.HelpFormatter(func(cmd *core.Command) string {
//		return banner + core.DefaultHelpFormatter(cmd)
//	})
func HelpFormatter(f core.HelpFormatter) {
	app.WithHelpFormatter(f)
}

// ErrorFormatter replaces the function used to render error messages
// before they are written to stderr, allowing you to fully customize the
// error style. The default is DefaultErrorFormatter. Parsing errors are typed
// (e.g. *UnknownFlagError, *MissingRequiredFlagError), so the formatter can
// inspect them with errors.As.
func ErrorFormatter(f core.ErrorFormatter) {
	app.WithErrorFormatter(f)
}

// UsePanic configures the CLI to panic instead of exiting when
// an error  occurs or when a command finishes execution. This can be useful
// for testing purposes or customization of the cli behavior. The panic will
// be called with the exit code as the argument.
func UsePanic(usePanic bool) {
	app.UsePanic(usePanic)
}

// AllowExtraPos configures the CLI to allow extra positional arguments
// that are not defined in the command.  You may use variadic positional for
// this purpose as well. By default, extra positional arguments are not allowed.
//
// Extra positional arguments can be accessed using the `Arg` function or the
// `ExtraArg` function.
func AllowExtraPos(allow bool) {
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

// AutoHelp configures the CLI to automatically show the help message when the user
// provides the `-h` or `--help` flag. By default, auto help is disabled.
func AutoHelp(enabled bool) {
	app.WithAutoHelp(enabled)
}

func Args(args []string) {
	app.WithArgs(args)
}

// Command is an alias for Cmd. It creates a new command with the specified
// name, short description, and execution function. The command is added as a
// subcommand to the current command. The execute function will be called when
// the command is invoked by the user.
func Command(name string, shortDescription string, execute func()) *core.Command {
	c := core.NewCommand(app.CurrentCommand()).
		WithName(name).
		WithShortDescription(shortDescription).
		WithDescription(shortDescription).
		WithExecute(execute)

	app.CurrentCommand().
		WithSubcommand(c)

	return c
}

// Example adds an example usage for the current command. The example is used in
// the help output for this command. The usage string should show how to use the
// command, and the description should explain what the example does.
func Example(usage, description string) {
	app.CurrentCommand().WithExample(usage, description)
}

// Pos is an alias for PosString. It creates a string positional argument with
// the given name and description.
func Pos(name, description string) *core.Positional[string] {
	return PosString(name, description)
}

// PosFunc creates a positional argument of any type T, using the provided
// parser function to convert the user input string into T. Use this to define
// custom positional types, e.g.:
//
//	level := cli.PosFunc("level", "The log level.", ParseLevel)
func PosFunc[T any](name, description string, parser func(string) (T, error)) *core.Positional[T] {
	p := core.NewPositional(name, description, parser)
	app.CurrentCommand().WithPositional(p)
	return p
}

// PosString creates a string positional argument with the given name and
// description.
func PosString(name, description string) *core.Positional[string] {
	p := core.NewPositional(name, description, parsers.String)
	app.CurrentCommand().WithPositional(p)
	return p
}

// PosInt creates an integer positional argument with the given name and
// description.
func PosInt(name, description string) *core.Positional[int] {
	p := core.NewPositional(name, description, parsers.Int[int])
	app.CurrentCommand().WithPositional(p)
	return p
}

// PosUint creates an unsigned integer positional argument with the given name and
// description.
func PosUint(name, description string) *core.Positional[uint] {
	p := core.NewPositional(name, description, parsers.Uint[uint])
	app.CurrentCommand().WithPositional(p)
	return p
}

// PosFloat creates a floating-point positional argument with the given name and
// description.
func PosFloat(name, description string) *core.Positional[float64] {
	p := core.NewPositional(name, description, parsers.Float[float64])
	app.CurrentCommand().WithPositional(p)
	return p
}

// PosBool creates a boolean positional argument with the given name and
// description.
func PosBool(name, description string) *core.Positional[bool] {
	p := core.NewPositional(name, description, parsers.Bool)
	app.CurrentCommand().WithPositional(p)
	return p
}

// PosDuration creates a time.Duration positional argument with the given name and
// description.
//
// Durations are parsed using the time.ParseDuration function, which supports
// formats like "300ms", "-1.5h" or "2h45m".
func PosDuration(name, description string) *core.Positional[time.Duration] {
	p := core.NewPositional(name, description, parsers.Duration)
	app.CurrentCommand().WithPositional(p)
	return p
}

// Flag is an alias for FlagString. It creates a string flag with the given
// long name, short name, and description.
func Flag(long, short, description string) *core.Flag[string] {
	return FlagString(long, short, description)
}

// FlagFunc creates a flag of any type T, using the provided parser function to
// convert the user input string into T. Use this to define custom flag types,
// e.g.:
//
//	level := cli.FlagFunc("level", "l", "The log level.", ParseLevel)
func FlagFunc[T any](long, short, description string, parser func(string) (T, error)) *core.Flag[T] {
	f := core.NewFlag(long, short, description, parser)
	app.CurrentCommand().WithFlag(f)
	return f
}

// FlagString creates a string flag with the given long name, short name, and
// description.
func FlagString(long, short, description string) *core.Flag[string] {
	f := core.NewFlag(long, short, description, parsers.String)
	app.CurrentCommand().WithFlag(f)
	return f
}

// FlagInt creates an integer flag with the given long name, short name, and
// description.
func FlagInt(long, short, description string) *core.Flag[int] {
	f := core.NewFlag(long, short, description, parsers.Int[int])
	app.CurrentCommand().WithFlag(f)
	return f
}

// FlagUint creates an unsigned integer flag with the given long name, short
// name, and description.
func FlagUint(long, short, description string) *core.Flag[uint] {
	f := core.NewFlag(long, short, description, parsers.Uint[uint])
	app.CurrentCommand().WithFlag(f)
	return f
}

// FlagFloat creates a floating-point flag with the given long name, short
// name, and description.
func FlagFloat(long, short, description string) *core.Flag[float64] {
	f := core.NewFlag(long, short, description, parsers.Float[float64])
	app.CurrentCommand().WithFlag(f)
	return f
}

// FlagBool creates a boolean flag with the given long name, short name, and
// description.
func FlagBool(long, short, description string) *core.Flag[bool] {
	f := core.NewFlag(long, short, description, parsers.Bool)
	app.CurrentCommand().WithFlag(f)
	return f
}

// FlagDuration creates a time.Duration flag with the given long name, short
// name, and description.
//
// Durations are parsed using the time.ParseDuration function, which supports
// formats like "300ms", "-1.5h" or "2h45m".
func FlagDuration(long, short, description string) *core.Flag[time.Duration] {
	f := core.NewFlag(long, short, description, parsers.Duration)
	app.CurrentCommand().WithFlag(f)
	return f
}

// Fatal prints a formatted error message using the error formatter and the
// stderr writer, and then exits with code 1.
func Fatal(format string, v ...any) { app.Fatal(format, v...) }

// FatalIf checks if the provided error is not nil, and if so, it prints the error
// message using the error formatter and the stderr writer, and then exits with
// code 1.
func FatalIf(err error) { app.FatalIf(err) }

// Exit terminates the program with the provided exit code. An exit code of 0
// typically indicates successful execution, while a non-zero exit code
// indicates an error or abnormal termination.
func Exit(code int) {
	app.Exit(code)
}

// Help prints the help message for the current command, including its description,
// usage, and available flags and subcommands. This function uses the Stdout
// function to output the help message.
func Help() {
	app.Help()
}

// Parse processes the command-line arguments provided by the user and executes
// the appropriate command based on the defined command structure. This function
// should be called after all commands, flags, and positional arguments have
// been defined.
//
// Subcommands are executed based on the first argument that matches a defined
// name, interrupting the execution of code after the Parse() call on the
// parent commands.
func Parse() *core.Result {
	return app.Parse()
}

// ParseArgs is similar to Parse but allows you to specify a custom slice of
// arguments to parse instead of using the default os.Args. This can be useful for
// testing purposes or when you want to parse a specific set of arguments without
// relying on the command-line input.
//
// DO NOT PROVIDE the program name (i.e., os.Args[0]) in the args slice.
func ParseArgs(args []string) *core.Result {
	return app.ParseArgs(args)
}

// CheckExclusiveFlags checks that at most one of the provided flags is passed.
// This function should be called after Parse().
func CheckExclusiveFlags(flags ...core.AnyFlag) {
	parsedSignatures := []string{}
	for _, flag := range flags {
		if flag.IsProvided() {
			parsedSignatures = append(parsedSignatures, flag.Signature())
		}
	}

	if len(parsedSignatures) > 1 {
		app.FatalIf(errors.NewExclusiveFlagsError(parsedSignatures))
	}
}

// CheckAnyFlag checks that at least one of the provided flags is passed. This
// function should be called after Parse().
func CheckAnyFlag(flags ...core.AnyFlag) {
	for _, flag := range flags {
		if flag.IsProvided() {
			return
		}
	}

	signatures := make([]string, 0, len(flags))
	for _, flag := range flags {
		signatures = append(signatures, flag.Signature())
	}
	app.FatalIf(errors.NewAtLeastOneFlagError(signatures))
}

// IsParsed returns true if the arguments have been parsed
// successfully.
func IsParsed() bool {
	return app.IsParsed()
}

// GetRootCommand returns the root command from the stack.
func GetRootCommand() *core.Command { return app.RootCommand() }

// GetCurrentCommand returns the current command being executed.
func GetCurrentCommand() *core.Command { return app.CurrentCommand() }

// GetHelp returns the help message for the current command as a string.
func GetHelp() string {
	return app.GetHelp()
}

// GetFlag retrieves a flag by its long or short name and attempts to cast it to the specified type T.
// If the flag is not found or cannot be cast to the desired type, an error is returned.
func GetFlag[T core.AnyFlag](longOrShort string) (T, error) {
	f, err := app.CurrentCommand().GetFlag(longOrShort)
	if err != nil {
		var zero T
		return zero, err
	}

	typed, ok := f.(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("flag is not of the expected type")
	}

	return typed, nil
}
