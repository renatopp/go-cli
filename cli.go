package cli

import (
	"io"
	"strings"
	"time"

	internal "github.com/renatopp/go-cli/internal"
)

var app *internal.App

type T_PosString = *internal.GenericPositional[string]
type T_PosInt = *internal.GenericPositional[int]
type T_PosInt8 = *internal.GenericPositional[int8]
type T_PosInt16 = *internal.GenericPositional[int16]
type T_PosInt32 = *internal.GenericPositional[int32]
type T_PosInt64 = *internal.GenericPositional[int64]
type T_PosUint = *internal.GenericPositional[uint]
type T_PosUint8 = *internal.GenericPositional[uint8]
type T_PosUint16 = *internal.GenericPositional[uint16]
type T_PosUint32 = *internal.GenericPositional[uint32]
type T_PosUint64 = *internal.GenericPositional[uint64]
type T_PosFloat = *internal.GenericPositional[float64]
type T_PosFloat32 = *internal.GenericPositional[float32]
type T_PosFloat64 = *internal.GenericPositional[float64]
type T_PosBool = *internal.GenericPositional[bool]
type T_PosDuration = *internal.GenericPositional[time.Duration]
type T_PosIntSlice = *internal.GenericPositional[[]int]
type T_PosInt8Slice = *internal.GenericPositional[[]int8]
type T_PosInt16Slice = *internal.GenericPositional[[]int16]
type T_PosInt32Slice = *internal.GenericPositional[[]int32]
type T_PosInt64Slice = *internal.GenericPositional[[]int64]
type T_PosUintSlice = *internal.GenericPositional[[]uint]
type T_PosUint8Slice = *internal.GenericPositional[[]uint8]
type T_PosUint16Slice = *internal.GenericPositional[[]uint16]
type T_PosUint32Slice = *internal.GenericPositional[[]uint32]
type T_PosUint64Slice = *internal.GenericPositional[[]uint64]
type T_PosFloatSlice = *internal.GenericPositional[[]float64]
type T_PosFloat32Slice = *internal.GenericPositional[[]float32]
type T_PosFloat64Slice = *internal.GenericPositional[[]float64]
type T_PosBoolSlice = *internal.GenericPositional[[]bool]
type T_PosDurationSlice = *internal.GenericPositional[[]time.Duration]
type T_FlagString = *internal.GenericFlag[string]
type T_FlagInt = *internal.GenericFlag[int]
type T_FlagInt8 = *internal.GenericFlag[int8]
type T_FlagInt16 = *internal.GenericFlag[int16]
type T_FlagInt32 = *internal.GenericFlag[int32]
type T_FlagInt64 = *internal.GenericFlag[int64]
type T_FlagUint = *internal.GenericFlag[uint]
type T_FlagUint8 = *internal.GenericFlag[uint8]
type T_FlagUint16 = *internal.GenericFlag[uint16]
type T_FlagUint32 = *internal.GenericFlag[uint32]
type T_FlagUint64 = *internal.GenericFlag[uint64]
type T_FlagFloat = *internal.GenericFlag[float64]
type T_FlagFloat32 = *internal.GenericFlag[float32]
type T_FlagFloat64 = *internal.GenericFlag[float64]
type T_FlagBool = *internal.GenericFlag[bool]
type T_FlagDuration = *internal.GenericFlag[time.Duration]
type T_FlagIntSlice = *internal.GenericFlag[[]int]
type T_FlagInt8Slice = *internal.GenericFlag[[]int8]
type T_FlagInt16Slice = *internal.GenericFlag[[]int16]
type T_FlagInt32Slice = *internal.GenericFlag[[]int32]
type T_FlagInt64Slice = *internal.GenericFlag[[]int64]
type T_FlagUintSlice = *internal.GenericFlag[[]uint]
type T_FlagUint8Slice = *internal.GenericFlag[[]uint8]
type T_FlagUint16Slice = *internal.GenericFlag[[]uint16]
type T_FlagUint32Slice = *internal.GenericFlag[[]uint32]
type T_FlagUint64Slice = *internal.GenericFlag[[]uint64]
type T_FlagFloatSlice = *internal.GenericFlag[[]float64]
type T_FlagFloat32Slice = *internal.GenericFlag[[]float32]
type T_FlagFloat64Slice = *internal.GenericFlag[[]float64]
type T_FlagBoolSlice = *internal.GenericFlag[[]bool]
type T_FlagDurationSlice = *internal.GenericFlag[[]time.Duration]

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

// Locale is a struct containing the user-facing strings used by go-cli, such
// as help text labels ("Usage", "Commands", "Options", ...) and error
// messages (unknown flag, missing required flag, etc). Any field left as the
// zero value falls back to the built-in English text, so you only need to
// override the strings you want to translate.
type Locale = internal.Locale

// SetLocale replaces the strings used for help text and error messages
// across the CLI, allowing you to localize go-cli's built-in output. It
// applies globally, independent of which App instance is used.
func SetLocale(locale Locale) {
	internal.SetLocale(locale)
}

var EN = internal.DefaultLocale()
var PTBR = internal.PTBRLocale()

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

// Shell creates a new shell command to be executed.
func Shell(name string, args ...string) *internal.Shell {
	return app.Shell(name, args...)
}

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

func Pos(name, description string) T_PosString {
	return app.Pos(name, description)
}
func PosString(name, description string) T_PosString {
	return app.PosString(name, description)
}
func PosInt(name, description string) T_PosInt {
	return app.PosInt(name, description)
}
func PosInt8(name, description string) T_PosInt8 {
	return app.PosInt8(name, description)
}
func PosInt16(name, description string) T_PosInt16 {
	return app.PosInt16(name, description)
}
func PosInt32(name, description string) T_PosInt32 {
	return app.PosInt32(name, description)
}
func PosInt64(name, description string) T_PosInt64 {
	return app.PosInt64(name, description)
}
func PosUint(name, description string) T_PosUint {
	return app.PosUint(name, description)
}
func PosUint8(name, description string) T_PosUint8 {
	return app.PosUint8(name, description)
}
func PosUint16(name, description string) T_PosUint16 {
	return app.PosUint16(name, description)
}
func PosUint32(name, description string) T_PosUint32 {
	return app.PosUint32(name, description)
}
func PosUint64(name, description string) T_PosUint64 {
	return app.PosUint64(name, description)
}
func PosFloat(name, description string) T_PosFloat {
	return app.PosFloat(name, description)
}
func PosFloat32(name, description string) T_PosFloat32 {
	return app.PosFloat32(name, description)
}
func PosFloat64(name, description string) T_PosFloat64 {
	return app.PosFloat64(name, description)
}
func PosBool(name, description string) T_PosBool {
	return app.PosBool(name, description)
}
func PosDuration(name, description string) T_PosDuration {
	return app.PosDuration(name, description)
}
func PosIntSlice(name, description string) T_PosIntSlice {
	return app.PosIntSlice(name, description)
}
func PosInt8Slice(name, description string) T_PosInt8Slice {
	return app.PosInt8Slice(name, description)
}
func PosInt16Slice(name, description string) T_PosInt16Slice {
	return app.PosInt16Slice(name, description)
}
func PosInt32Slice(name, description string) T_PosInt32Slice {
	return app.PosInt32Slice(name, description)
}
func PosInt64Slice(name, description string) T_PosInt64Slice {
	return app.PosInt64Slice(name, description)
}
func PosUintSlice(name, description string) T_PosUintSlice {
	return app.PosUintSlice(name, description)
}
func PosUint8Slice(name, description string) T_PosUint8Slice {
	return app.PosUint8Slice(name, description)
}
func PosUint16Slice(name, description string) T_PosUint16Slice {
	return app.PosUint16Slice(name, description)
}
func PosUint32Slice(name, description string) T_PosUint32Slice {
	return app.PosUint32Slice(name, description)
}
func PosUint64Slice(name, description string) T_PosUint64Slice {
	return app.PosUint64Slice(name, description)
}
func PosFloatSlice(name, description string) T_PosFloatSlice {
	return app.PosFloatSlice(name, description)
}
func PosFloat32Slice(name, description string) T_PosFloat32Slice {
	return app.PosFloat32Slice(name, description)
}
func PosFloat64Slice(name, description string) T_PosFloat64Slice {
	return app.PosFloat64Slice(name, description)
}
func PosBoolSlice(name, description string) T_PosBoolSlice {
	return app.PosBoolSlice(name, description)
}
func PosDurationSlice(name, description string) T_PosDurationSlice {
	return app.PosDurationSlice(name, description)
}

func Flag(long, short, description string) T_FlagString {
	return app.Flag(long, short, description)
}
func FlagString(long, short, description string) T_FlagString {
	return app.FlagString(long, short, description)
}
func FlagInt(long, short, description string) T_FlagInt {
	return app.FlagInt(long, short, description)
}
func FlagInt8(long, short, description string) T_FlagInt8 {
	return app.FlagInt8(long, short, description)
}
func FlagInt16(long, short, description string) T_FlagInt16 {
	return app.FlagInt16(long, short, description)
}
func FlagInt32(long, short, description string) T_FlagInt32 {
	return app.FlagInt32(long, short, description)
}
func FlagInt64(long, short, description string) T_FlagInt64 {
	return app.FlagInt64(long, short, description)
}
func FlagUint(long, short, description string) T_FlagUint {
	return app.FlagUint(long, short, description)
}
func FlagUint8(long, short, description string) T_FlagUint8 {
	return app.FlagUint8(long, short, description)
}
func FlagUint16(long, short, description string) T_FlagUint16 {
	return app.FlagUint16(long, short, description)
}
func FlagUint32(long, short, description string) T_FlagUint32 {
	return app.FlagUint32(long, short, description)
}
func FlagUint64(long, short, description string) T_FlagUint64 {
	return app.FlagUint64(long, short, description)
}
func FlagFloat(long, short, description string) T_FlagFloat {
	return app.FlagFloat(long, short, description)
}
func FlagFloat32(long, short, description string) T_FlagFloat32 {
	return app.FlagFloat32(long, short, description)
}
func FlagFloat64(long, short, description string) T_FlagFloat64 {
	return app.FlagFloat64(long, short, description)
}
func FlagBool(long, short, description string) T_FlagBool {
	return app.FlagBool(long, short, description)
}
func FlagDuration(long, short, description string) T_FlagDuration {
	return app.FlagDuration(long, short, description)
}
func FlagIntSlice(long, short, description string) T_FlagIntSlice {
	return app.FlagIntSlice(long, short, description)
}
func FlagInt8Slice(long, short, description string) T_FlagInt8Slice {
	return app.FlagInt8Slice(long, short, description)
}
func FlagInt16Slice(long, short, description string) T_FlagInt16Slice {
	return app.FlagInt16Slice(long, short, description)
}
func FlagInt32Slice(long, short, description string) T_FlagInt32Slice {
	return app.FlagInt32Slice(long, short, description)
}
func FlagInt64Slice(long, short, description string) T_FlagInt64Slice {
	return app.FlagInt64Slice(long, short, description)
}
func FlagUintSlice(long, short, description string) T_FlagUintSlice {
	return app.FlagUintSlice(long, short, description)
}
func FlagUint8Slice(long, short, description string) T_FlagUint8Slice {
	return app.FlagUint8Slice(long, short, description)
}
func FlagUint16Slice(long, short, description string) T_FlagUint16Slice {
	return app.FlagUint16Slice(long, short, description)
}
func FlagUint32Slice(long, short, description string) T_FlagUint32Slice {
	return app.FlagUint32Slice(long, short, description)
}
func FlagUint64Slice(long, short, description string) T_FlagUint64Slice {
	return app.FlagUint64Slice(long, short, description)
}
func FlagFloatSlice(long, short, description string) T_FlagFloatSlice {
	return app.FlagFloatSlice(long, short, description)
}
func FlagFloat32Slice(long, short, description string) T_FlagFloat32Slice {
	return app.FlagFloat32Slice(long, short, description)
}
func FlagFloat64Slice(long, short, description string) T_FlagFloat64Slice {
	return app.FlagFloat64Slice(long, short, description)
}
func FlagBoolSlice(long, short, description string) T_FlagBoolSlice {
	return app.FlagBoolSlice(long, short, description)
}
func FlagDurationSlice(long, short, description string) T_FlagDurationSlice {
	return app.FlagDurationSlice(long, short, description)
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
