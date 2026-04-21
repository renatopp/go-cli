package tests

import (
	"testing"

	"github.com/renatopp/go-cli"
)

func TestCommand(t *testing.T) {
	defer cli.Clear()
	args := make_args("test", "value1", "--flag1", "value2")
	cli.UsePanicInsteadOfExit(true)
	cli.Command("test", "a test command", func() {
		cli.Pos("arg1", "first argument").AsRequired()
		cli.Pos("arg2", "second argument").WithDefault("defaulted")
		cli.Flag("flag1", "f", "a flag").AsRequired()
		cli.Flag("flag2", "g", "another flag").WithDefault("defaulted")
		cli.ParseArgs(args)
	})
	expectPanicWith(t, func() {
		cli.ParseArgs(args)
	}, 0)
}
