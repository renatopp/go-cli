package cli

import v2 "github.com/renatopp/go-cli/v2"

var app *v2.App

// Initialize global state
func init() {
	app = v2.NewApp()
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

// // Command registers a new subcommand within current command, allowing nested
// // subcommand structures.
// func Command(name, shortDescription string, execute func()) *internal.Command {
// 	cmd := internal.NewCommand().
// 		WithName(name).
// 		WithShortDescription(shortDescription).
// 		WithExecute(execute)
// 	internal.GetCmd().WithSubcommand(cmd)
// 	return cmd
// }

// // Cmd is an alias for Command, provided for convenience and readability.
// func Cmd(name, shortDescription string, execute func()) *internal.Command {
// 	return Command(name, shortDescription, execute)
// }

// // Flag registers a flag option for the current command. This is an alias
// // for FlagString. Flags capture user input in the form of --flag=value or -f
// // value.
// func Flag(long, short, description string) *internal.StringFlag {
// 	return FlagString(long, short, description)
// }

// // FlagInt registers an integer flag option for the current command. Flags
// // capture user input in the form of --flag=value or -f value.
// //
// // The IntFlag will attempt to parse the provided value as an integer, and if
// // parsing fails, go-cli will handle the error appropriately.
// func FlagInt(long, short, description string) *internal.IntFlag {
// 	flag := internal.NewIntFlag(long, short, description)
// 	internal.GetCmd().WithFlag(flag)
// 	return flag
// }

// // FlagInt64 registers a 64-bit integer flag option for the current command.
// // Flags capture user input in the form of --flag=value or -f value.
// //
// // The Int64Flag will attempt to parse the provided value as a 64-bit integer,
// // and if parsing fails, go-cli will handle the error appropriately.
// func FlagInt64(long, short, description string) *internal.Int64Flag {
// 	flag := internal.NewInt64Flag(long, short, description)
// 	internal.GetCmd().WithFlag(flag)
// 	return flag
// }

// // FlagUint registers an unsigned integer flag option for the current command.
// // Flags capture user input in the form of --flag=value or -f value.
// //
// // The UintFlag will attempt to parse the provided value as an unsigned integer,
// // and if parsing fails, go-cli will handle the error appropriately.
// func FlagUint(long, short, description string) *internal.UintFlag {
// 	flag := internal.NewUintFlag(long, short, description)
// 	internal.GetCmd().WithFlag(flag)
// 	return flag
// }

// // FlagUint64 registers a 64-bit unsigned integer flag option for the current command.
// // Flags capture user input in the form of --flag=value or -f value.
// //
// // The Uint64Flag will attempt to parse the provided value as a 64-bit unsigned integer,
// // and if parsing fails, go-cli will handle the error appropriately.
// func FlagUint64(long, short, description string) *internal.Uint64Flag {
// 	flag := internal.NewUint64Flag(long, short, description)
// 	internal.GetCmd().WithFlag(flag)
// 	return flag
// }

// // FlagFloat registers a floating-point flag option for the current command.
// // Flags capture user input in the form of --flag=value or -f value.
// //
// // The FloatFlag will attempt to parse the provided value as a floating-point
// // number, and if parsing fails, go-cli will handle the error appropriately.
// func FlagFloat(long, short, description string) *internal.FloatFlag {
// 	flag := internal.NewFloatFlag(long, short, description)
// 	internal.GetCmd().WithFlag(flag)
// 	return flag
// }

// // FlagBool registers a boolean flag option for the current command. Flags
// // capture user input in the form of --flag or -f. Short boolean flags DO NOT
// // ACCEPT argument but they can be  merged into a single argument (e.g., -abc
// // is equivalent to -a -b -c). Long boolean flags receive an optional argument
// // that will be parsed, possible values are "true", "false", "1", "0". If no
// // argument is provided the flag will be set to true.
// func FlagBool(long, short, description string) *internal.BoolFlag {
// 	flag := internal.NewBoolFlag(long, short, description)
// 	internal.GetCmd().WithFlag(flag)
// 	return flag
// }

// // FlagStringSlice registers a string slice flag option for the current command.
// // String slices are strings split by comma so the user can provide multiple
// // values in a single flag (e.g., --flag=value1,value2,value3). Flags capture
// // user input in the form of --flag=value or -f value.
// func FlagStringSlice(long, short, description string) *internal.StringSliceFlag {
// 	flag := internal.NewStringSliceFlag(long, short, description)
// 	internal.GetCmd().WithFlag(flag)
// 	return flag
// }

// // FlagString registers a string flag option for the current command. Flags
// // capture user input in the form of --flag=value or -f value.
// func FlagString(long, short, description string) *internal.StringFlag {
// 	flag := internal.NewStringFlag(long, short, description)
// 	internal.GetCmd().WithFlag(flag)
// 	return flag
// }

// // Pos registers a string positional argument for the current command. Positional
// // arguments capture user input based on their position in the command line. For
// // example, in the command "app command arg1 arg2", "arg1" and "arg2" are
// // positional arguments.
// func Pos(index int, name, description string) *internal.StringPositional {
// 	arg := internal.NewStringPositional(index, name, description)
// 	internal.GetCmd().WithPositional(arg)
// 	return arg
// }

// // PosInt registers an integer positional argument for the current command.
// // Positional arguments capture user input based on their position in the
// // command line. For example, in the command "app command arg1 arg2", "arg1"
// // and "arg2" are positional arguments.
// // The PosInt positional will attempt to parse the provided value as an integer,
// // and if parsing fails, go-cli will handle the error appropriately.
// func PosInt(index int, name, description string) *internal.IntPositional {
// 	arg := internal.NewIntPositional(index, name, description)
// 	internal.GetCmd().WithPositional(arg)
// 	return arg
// }

// // PosInt64 registers an integer positional argument for the current command.
// // Positional arguments capture user input based on their position in the
// // command line. For example, in the command "app command arg1 arg2", "arg1"
// // and "arg2" are positional arguments.
// // The PosInt64 positional will attempt to parse the provided value as an integer,
// // and if parsing fails, go-cli will handle the error appropriately.
// func PosInt64(index int, name, description string) *internal.Int64Positional {
// 	arg := internal.NewInt64Positional(index, name, description)
// 	internal.GetCmd().WithPositional(arg)
// 	return arg
// }

// // PosFloat registers a floating-point positional argument for the current command.
// // Positional arguments capture user input based on their position in the
// // command line. For example, in the command "app command arg1 arg2", "arg1"
// // and "arg2" are positional arguments.
// // The PosFloat positional will attempt to parse the provided value as a
// // floating-point number, and if parsing fails, go-cli will handle the error
// // appropriately.
// func PosFloat(index int, name, description string) *internal.FloatPositional {
// 	arg := internal.NewFloatPositional(index, name, description)
// 	internal.GetCmd().WithPositional(arg)
// 	return arg
// }

// // PosString registers a string positional argument for the current command.
// // Positional arguments capture user input based on their position in the
// // command line. For example, in the command "app command arg1 arg2", "arg1"
// // and "arg2" are positional arguments.
// func PosString(index int, name, description string) *internal.StringPositional {
// 	arg := internal.NewStringPositional(index, name, description)
// 	internal.GetCmd().WithPositional(arg)
// 	return arg
// }

// // NArgs returns the number of positional arguments provided by the user.
// // Should be used only after Parse() is called, otherwise it will return 0.
// func NArgs() int {
// 	return internal.GetApp().NArgs()
// }

// // Arg retrieves the value of a positional argument by its index.
// // Should be used only after Parse() is called, otherwise it will return an
// // empty string.
// func Arg(index int) string {
// 	return internal.GetApp().Arg(index)
// }

// // Args retrieves all positional arguments provided by the user.
// // Should be used only after Parse() is called, otherwise it will return an
// // empty slice.
// func Args() []string {
// 	return internal.GetApp().Args()
// }
