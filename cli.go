package cli

import (
	"time"

	internal "github.com/renatopp/go-cli/internal"
)

var app *internal.App

// Initialize global state
func init() {
	app = internal.NewApp()
}

// Name sets the name for the current command. The name is used in help text
// to identify the command and its usage. Use only its immediate name (e.g.
// "version" instead of "app version") since the command hierarchy is
// automatically handled by go-cli.
func Name(n string) {
	app.CurrentCommand().WithName(n)
}

// Description sets the description for the current command. Descriptions are
// used in help text to provide more information about the command and its
// purpose.
func Description(d string) {
	app.CurrentCommand().WithDescription(d)
}

// ShowHelp prints the help message for the current command, including its description,
// usage, and available flags and subcommands.
func ShowHelp() {
	app.ShowHelp()
}

func Help() string {
	return app.GetHelpString()
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

func ParseArgs(args []string) {
	app.ParseArgs(args)
}

// Clear resets the state of the CLI, allowing users to define a new command
// structure and configuration. This can be useful in scenarios where you want to
// reuse the same CLI instance for different commands or when you want to reset
// the CLI state after executing a command.
func Clear() {
	app.Clear()
}

// Exit terminates the program with the provided exit code. An exit code of 0
// typically indicates successful execution, while a non-zero exit code
// indicates an error or abnormal termination.
func Exit(code int) {
	app.Exit(code)
}

func StdoutWith(fn func(msg string, args ...any)) {
	app.Stdout = fn
}

func StderrWith(fn func(msg string, args ...any)) {
	app.Stderr = fn
}

func UsePanicInsteadOfExit(usePanic bool) {
	app.PanicInsteadOfExit = usePanic
}

func AllowExtraPositionals(allow bool) {
	app.ExtraPositionalsAllowed = allow
}

func AllowExtraFlags(allow bool) {
	app.ExtraFlagsAllowed = allow
}

func AllowRepeatedFlags(allow bool) {
	app.RepeatedFlagsAllowed = allow
}

// NArgs returns the number of positional arguments provided by the user.
// Should be used only after Parse() is called, otherwise it will return 0.
func NArgs() int {
	if !app.IsParsed() {
		return 0
	}
	return len(app.Arguments().Args)
}

// Arg retrieves the value of a positional argument by its index.
// Should be used only after Parse() is called, otherwise it will return an
// empty string.
func Arg(index int) string {
	if !app.IsParsed() {
		return ""
	}
	args := app.Arguments().Args
	if index < 0 || index >= len(args) {
		return ""
	}
	return args[index]
}

// Args retrieves all positional arguments provided by the user.
// Should be used only after Parse() is called, otherwise it will return an
// empty slice.
func Args() []string {
	if !app.IsParsed() {
		return []string{}
	}
	return app.Arguments().Args
}

func NExtraArgs() int {
	if !app.IsParsed() {
		return 0
	}
	return len(app.Arguments().ExtraArgs)
}

func ExtraArg(index int) string {
	if !app.IsParsed() {
		return ""
	}
	extraArgs := app.Arguments().ExtraArgs
	if index < 0 || index >= len(extraArgs) {
		return ""
	}
	return extraArgs[index]
}

// ExtraArgs retrieves all extra positional arguments provided by the user, i.e.,
// those that are not defined in the command. Should be used only after Parse() is
// called, otherwise it will return an empty slice.
func ExtraArgs() []string {
	if !app.IsParsed() {
		return []string{}
	}
	return app.Arguments().ExtraArgs
}

func Command(name string, shortDescription string, execute func()) *internal.Command {
	cmd := internal.NewCommand().WithName(name).WithShortDescription(shortDescription).WithExecute(execute)
	app.CurrentCommand().WithSubcommand(cmd)
	return cmd
}

func Cmd(name string, shortDescription string, execute func()) *internal.Command {
	return Command(name, shortDescription, execute)
}

func _addpos[T internal.Positional](a T) T {
	app.CurrentCommand().WithPositional(a)
	return a
}

func Pos(name, description string) *internal.GenericPositional[string] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseString))
}
func PosInt(name, description string) *internal.GenericPositional[int] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseInt[int]))
}
func PosInt8(name, description string) *internal.GenericPositional[int8] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseInt[int8]))
}
func PosInt16(name, description string) *internal.GenericPositional[int16] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseInt[int16]))
}
func PosInt32(name, description string) *internal.GenericPositional[int32] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseInt[int32]))
}
func PosInt64(name, description string) *internal.GenericPositional[int64] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseInt[int64]))
}
func PosUint(name, description string) *internal.GenericPositional[uint] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseUint[uint]))
}
func PosUint8(name, description string) *internal.GenericPositional[uint8] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseUint[uint8]))
}
func PosUint16(name, description string) *internal.GenericPositional[uint16] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseUint[uint16]))
}
func PosUint32(name, description string) *internal.GenericPositional[uint32] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseUint[uint32]))
}
func PosUint64(name, description string) *internal.GenericPositional[uint64] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseUint[uint64]))
}
func PosFloat32(name, description string) *internal.GenericPositional[float32] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseFloat[float32]))
}
func PosFloat64(name, description string) *internal.GenericPositional[float64] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseFloat[float64]))
}
func PosBool(name, description string) *internal.GenericPositional[bool] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseBool))
}
func PosDuration(name, description string) *internal.GenericPositional[time.Duration] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseDuration))
}
func PosIntSlice(name, description string) *internal.GenericPositional[[]int] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseIntSlice[int]))
}
func PosInt8Slice(name, description string) *internal.GenericPositional[[]int8] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseIntSlice[int8]))
}
func PosInt16Slice(name, description string) *internal.GenericPositional[[]int16] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseIntSlice[int16]))
}
func PosInt32Slice(name, description string) *internal.GenericPositional[[]int32] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseIntSlice[int32]))
}
func PosInt64Slice(name, description string) *internal.GenericPositional[[]int64] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseIntSlice[int64]))
}
func PosUintSlice(name, description string) *internal.GenericPositional[[]uint] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseUintSlice[uint]))
}
func PosUint8Slice(name, description string) *internal.GenericPositional[[]uint8] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseUintSlice[uint8]))
}
func PosUint16Slice(name, description string) *internal.GenericPositional[[]uint16] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseUintSlice[uint16]))
}
func PosUint32Slice(name, description string) *internal.GenericPositional[[]uint32] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseUintSlice[uint32]))
}
func PosUint64Slice(name, description string) *internal.GenericPositional[[]uint64] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseUintSlice[uint64]))
}
func PosFloat32Slice(name, description string) *internal.GenericPositional[[]float32] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseFloatSlice[float32]))
}
func PosFloat64Slice(name, description string) *internal.GenericPositional[[]float64] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseFloatSlice[float64]))
}
func PosBoolSlice(name, description string) *internal.GenericPositional[[]bool] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseBoolSlice))
}
func PosDurationSlice(name, description string) *internal.GenericPositional[[]time.Duration] {
	return _addpos(internal.NewGenericPositional(name, description, internal.ParseDurationSlice))
}

func _addflag[T internal.Flag](a T) T {
	app.CurrentCommand().WithFlag(a)
	return a
}

func Flag(long, short, description string) *internal.GenericFlag[string] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseString))
}
func FlagInt(long, short, description string) *internal.GenericFlag[int] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseInt[int]))
}
func FlagInt8(long, short, description string) *internal.GenericFlag[int8] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseInt[int8]))
}
func FlagInt16(long, short, description string) *internal.GenericFlag[int16] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseInt[int16]))
}
func FlagInt32(long, short, description string) *internal.GenericFlag[int32] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseInt[int32]))
}
func FlagInt64(long, short, description string) *internal.GenericFlag[int64] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseInt[int64]))
}
func FlagUint(long, short, description string) *internal.GenericFlag[uint] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseUint[uint]))
}
func FlagUint8(long, short, description string) *internal.GenericFlag[uint8] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseUint[uint8]))
}
func FlagUint16(long, short, description string) *internal.GenericFlag[uint16] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseUint[uint16]))
}
func FlagUint32(long, short, description string) *internal.GenericFlag[uint32] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseUint[uint32]))
}
func FlagUint64(long, short, description string) *internal.GenericFlag[uint64] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseUint[uint64]))
}
func FlagFloat32(long, short, description string) *internal.GenericFlag[float32] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseFloat[float32]))
}
func FlagFloat64(long, short, description string) *internal.GenericFlag[float64] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseFloat[float64]))
}
func FlagBool(long, short, description string) *internal.GenericFlag[bool] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseBool))
}
func FlagDuration(long, short, description string) *internal.GenericFlag[time.Duration] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseDuration))
}
func FlagIntSlice(long, short, description string) *internal.GenericFlag[[]int] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseIntSlice[int]))
}
func FlagInt8Slice(long, short, description string) *internal.GenericFlag[[]int8] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseIntSlice[int8]))
}
func FlagInt16Slice(long, short, description string) *internal.GenericFlag[[]int16] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseIntSlice[int16]))
}
func FlagInt32Slice(long, short, description string) *internal.GenericFlag[[]int32] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseIntSlice[int32]))
}
func FlagInt64Slice(long, short, description string) *internal.GenericFlag[[]int64] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseIntSlice[int64]))
}
func FlagUintSlice(long, short, description string) *internal.GenericFlag[[]uint] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseUintSlice[uint]))
}
func FlagUint8Slice(long, short, description string) *internal.GenericFlag[[]uint8] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseUintSlice[uint8]))
}
func FlagUint16Slice(long, short, description string) *internal.GenericFlag[[]uint16] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseUintSlice[uint16]))
}
func FlagUint32Slice(long, short, description string) *internal.GenericFlag[[]uint32] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseUintSlice[uint32]))
}
func FlagUint64Slice(long, short, description string) *internal.GenericFlag[[]uint64] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseUintSlice[uint64]))
}
func FlagFloat32Slice(long, short, description string) *internal.GenericFlag[[]float32] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseFloatSlice[float32]))
}
func FlagFloat64Slice(long, short, description string) *internal.GenericFlag[[]float64] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseFloatSlice[float64]))
}
func FlagBoolSlice(long, short, description string) *internal.GenericFlag[[]bool] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseBoolSlice))
}
func FlagDurationSlice(long, short, description string) *internal.GenericFlag[[]time.Duration] {
	return _addflag(internal.NewGenericFlag(long, short, description, internal.ParseDurationSlice))
}
