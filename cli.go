package cli

import (
	"os"

	"github.com/renatopp/go-cli/internal"
)

func Parse()               { internal.GetApp().Parse() }
func Exit(code int)        { os.Exit(code) }
func Clear()               { internal.Clear() }
func ShowHelp()            { internal.GetApp().Help() }
func Name(n string)        { internal.GetCmd().WithName(n) }
func Description(d string) { internal.GetCmd().WithDescription(d) }
func Restricted()          { internal.GetApp().WithRestrict(true) }
func AutoHelp()            { internal.GetApp().WithAutoHelp(true) }
func PrintWith(printf func(format string, a ...any)) {
	internal.GetApp().WithPrintFunc(printf)
}

//

func Command(name, shortDescription string, execute func()) *internal.Command {
	cmd := internal.NewCommand().
		WithName(name).
		WithShortDescription(shortDescription).
		WithExecute(execute)
	internal.GetCmd().WithSubcommand(cmd)
	return cmd
}
func Cmd(name, shortDescription string, execute func()) *internal.Command {
	return Command(name, shortDescription, execute)
}

//

func FlagInt(long, short, description string) *internal.IntFlag {
	flag := internal.NewIntFlag(long, short, description)
	internal.GetCmd().WithFlag(flag)
	return flag
}

func FlagInt64(long, short, description string) *internal.Int64Flag {
	flag := internal.NewInt64Flag(long, short, description)
	internal.GetCmd().WithFlag(flag)
	return flag
}

func FlagFloat(long, short, description string) *internal.FloatFlag {
	flag := internal.NewFloatFlag(long, short, description)
	internal.GetCmd().WithFlag(flag)
	return flag
}

func FlagBool(long, short, description string) *internal.BoolFlag {
	flag := internal.NewBoolFlag(long, short, description)
	internal.GetCmd().WithFlag(flag)
	return flag
}

func FlagStringSlice(long, short, description string) *internal.StringSliceFlag {
	flag := internal.NewStringSliceFlag(long, short, description)
	internal.GetCmd().WithFlag(flag)
	return flag
}

func FlagString(long, short, description string) *internal.StringFlag {
	flag := internal.NewStringFlag(long, short, description)
	internal.GetCmd().WithFlag(flag)
	return flag
}

func Flag(long, short, description string) *internal.StringFlag {
	return FlagString(long, short, description)
}

//

func PosInt(index int, name, description string) *internal.StringPositional {
	arg := internal.NewStringPositional(index, name, description)
	internal.GetCmd().WithPositional(arg)
	return arg
}

func PosInt64(index int, name, description string) *internal.Int64Positional {
	arg := internal.NewInt64Positional(index, name, description)
	internal.GetCmd().WithPositional(arg)
	return arg
}

func PosFloat(index int, name, description string) *internal.FloatPositional {
	arg := internal.NewFloatPositional(index, name, description)
	internal.GetCmd().WithPositional(arg)
	return arg
}

func PosString(index int, name, description string) *internal.StringPositional {
	arg := internal.NewStringPositional(index, name, description)
	internal.GetCmd().WithPositional(arg)
	return arg
}

func Pos(index int, name, description string) *internal.StringPositional {
	arg := internal.NewStringPositional(index, name, description)
	internal.GetCmd().WithPositional(arg)
	return arg
}

//

func NArgs() int {
	return internal.GetApp().NArgs()
}

func Arg(index int) string {
	return internal.GetApp().Arg(index)
}

func Args() []string {
	return internal.GetApp().Args()
}
