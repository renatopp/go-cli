package cli

import (
	"errors"
	"io"
	"time"

	"github.com/renatopp/go-cli/pkg"
	cerrors "github.com/renatopp/go-cli/pkg/errors"
	"github.com/renatopp/go-cli/pkg/formatters"
	"github.com/renatopp/go-cli/pkg/locales"
	"github.com/renatopp/go-cli/pkg/parsers"
)

var app *pkg.App

// Initialize global state
func init() {
	app = pkg.NewApp(formatters.DefaultHelpFormatter, formatters.DefaultErrorFormatter)
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
//	cli.HelpFormatter(func(cmd *pkg.Command) string {
//		return banner + pkg.DefaultHelpFormatter(cmd)
//	})
func HelpFormatter(f pkg.HelpFormatter) {
	app.WithHelpFormatter(f)
}

// ErrorFormatter replaces the function used to render error messages
// before they are written to stderr, allowing you to fully customize the
// error style. The default is DefaultErrorFormatter. Parsing errors are typed
// (e.g. *UnknownFlagError, *MissingRequiredFlagError), so the formatter can
// inspect them with errors.As.
func ErrorFormatter(f pkg.ErrorFormatter) {
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
func Command(name string, shortDescription string, execute func()) *pkg.Command {
	c := pkg.NewCommand(app.CurrentCommand()).
		WithName(name).
		WithShortDescription(shortDescription).
		WithDescription(shortDescription).
		WithExecute(execute)

	app.CurrentCommand().
		WithSubcommand(c)

	return c
}

func Pos(name, description string) *pkg.GenericPositional[string] {
	return PosString(name, description)
}

// PosFunc creates a positional argument of any type T, using the provided
// parser function to convert the user input string into T. Use this to define
// custom positional types, e.g.:
//
//	level := cli.PosFunc("level", "The log level.", ParseLevel)
func PosFunc[T any](name, description string, parser func(string) (T, error)) *pkg.GenericPositional[T] {
	p := pkg.NewGenericPositional(name, description, parser)
	app.CurrentCommand().WithPositional(p)
	return p
}

func PosString(name, description string) *pkg.GenericPositional[string] {
	p := pkg.NewGenericPositional(name, description, parsers.String)
	app.CurrentCommand().WithPositional(p)
	return p
}

func PosInt(name, description string) *pkg.GenericPositional[int] {
	p := pkg.NewGenericPositional(name, description, parsers.Int[int])
	app.CurrentCommand().WithPositional(p)
	return p
}

func PosUint(name, description string) *pkg.GenericPositional[uint] {
	p := pkg.NewGenericPositional(name, description, parsers.Uint[uint])
	app.CurrentCommand().WithPositional(p)
	return p
}

func PosFloat(name, description string) *pkg.GenericPositional[float64] {
	p := pkg.NewGenericPositional(name, description, parsers.Float[float64])
	app.CurrentCommand().WithPositional(p)
	return p
}

func PosBool(name, description string) *pkg.GenericPositional[bool] {
	p := pkg.NewGenericPositional(name, description, parsers.Bool)
	app.CurrentCommand().WithPositional(p)
	return p
}

func PosDuration(name, description string) *pkg.GenericPositional[time.Duration] {
	p := pkg.NewGenericPositional(name, description, parsers.Duration)
	app.CurrentCommand().WithPositional(p)
	return p
}

// Flag is an alias for FlagString. It creates a string flag with the given
// long name, short name, and description.
func Flag(long, short, description string) *pkg.GenericFlag[string] {
	return FlagString(long, short, description)
}

// FlagFunc creates a flag of any type T, using the provided parser function to
// convert the user input string into T. Use this to define custom flag types,
// e.g.:
//
//	level := cli.FlagFunc("level", "l", "The log level.", ParseLevel)
func FlagFunc[T any](long, short, description string, parser func(string) (T, error)) *pkg.GenericFlag[T] {
	f := pkg.NewGenericFlag(long, short, description, parser)
	app.CurrentCommand().WithFlag(f)
	return f
}

func FlagString(long, short, description string) *pkg.GenericFlag[string] {
	f := pkg.NewGenericFlag(long, short, description, parsers.String)
	app.CurrentCommand().WithFlag(f)
	return f
}

func FlagInt(long, short, description string) *pkg.GenericFlag[int] {
	f := pkg.NewGenericFlag(long, short, description, parsers.Int[int])
	app.CurrentCommand().WithFlag(f)
	return f
}

func FlagUint(long, short, description string) *pkg.GenericFlag[uint] {
	f := pkg.NewGenericFlag(long, short, description, parsers.Uint[uint])
	app.CurrentCommand().WithFlag(f)
	return f
}

func FlagFloat(long, short, description string) *pkg.GenericFlag[float64] {
	f := pkg.NewGenericFlag(long, short, description, parsers.Float[float64])
	app.CurrentCommand().WithFlag(f)
	return f
}

func FlagBool(long, short, description string) *pkg.GenericFlag[bool] {
	f := pkg.NewGenericFlag(long, short, description, parsers.Bool)
	app.CurrentCommand().WithFlag(f)
	return f
}

func FlagDuration(long, short, description string) *pkg.GenericFlag[time.Duration] {
	f := pkg.NewGenericFlag(long, short, description, parsers.Duration)
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
func Parse() *pkg.Arguments {
	return app.Parse()
}

// ParseArgs is similar to Parse but allows you to specify a custom slice of
// arguments to parse instead of using the default os.Args. This can be useful for
// testing purposes or when you want to parse a specific set of arguments without
// relying on the command-line input.
//
// DO NOT PROVIDE the program name (i.e., os.Args[0]) in the args slice.
func ParseArgs(args []string) *pkg.Arguments {
	return app.ParseArgs(args)
}

// CheckExclusiveFlags checks that at most one of the provided flags is passed.
// This function should be called after Parse().
func CheckExclusiveFlags(flags ...pkg.Flag) {
	parsedSignatures := []string{}
	for _, flag := range flags {
		if flag.IsParsed() {
			parsedSignatures = append(parsedSignatures, flag.Signature())
		}
	}

	if len(parsedSignatures) > 1 {
		app.FatalIf(cerrors.NewExclusiveFlagsError(parsedSignatures))
	}
}

// CheckAnyFlag checks that at least one of the provided flags is passed. This
// function should be called after Parse().
func CheckAnyFlag(flags ...pkg.Flag) {
	for _, flag := range flags {
		if flag.IsParsed() {
			return
		}
	}

	signatures := make([]string, 0, len(flags))
	for _, flag := range flags {
		signatures = append(signatures, flag.Signature())
	}
	app.FatalIf(cerrors.NewAtLeastOneFlagError(signatures))
}

// IsParsed returns true if the arguments have been parsed
// successfully.
func IsParsed() bool {
	return app.IsParsed()
}

// GetRootCommand returns the root command from the stack.
func GetRootCommand() *pkg.Command { return app.RootCommand() }

// GetCurrentCommand returns the current command being executed.
func GetCurrentCommand() *pkg.Command { return app.CurrentCommand() }

// GetHelp returns the help message for the current command as a string.
func GetHelp() string {
	return app.GetHelp()
}

// GetFlag retrieves a flag by its long or short name and attempts to cast it to the specified type T.
// If the flag is not found or cannot be cast to the desired type, an error is returned.
func GetFlag[T pkg.Flag](longOrShort string) (T, error) {
	f, err := app.CurrentCommand().GetFlag(longOrShort)
	if err != nil {
		var zero T
		return zero, err
	}

	typed, ok := f.(T)
	if !ok {
		var zero T
		return zero, errors.New("flag is not of the expected type")
	}

	return typed, nil
}
