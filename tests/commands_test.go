package tests

import (
	"strings"
	"testing"

	"github.com/renatopp/go-cli"
)

func TestCommand(t *testing.T) {
	defer cli.Clear()
	args := make_args("test", "value1", "--flag1", "value2")
	cli.UsePanic(true)
	cli.Command("test", "a test command", func() {
		cli.Pos("arg1", "first argument").AsRequired()
		cli.Pos("arg2", "second argument").WithDefault("defaulted")
		cli.FlagString("flag1", "f", "a flag").AsRequired()
		cli.FlagString("flag2", "g", "another flag").WithDefault("defaulted")
		cli.ParseArgs(args)
	})
	expectPanicWith(t, func() {
		cli.ParseArgs(args)
	}, 0)
}

func TestHiddenSubcommandNotInHelp(t *testing.T) {
	defer cli.Clear()

	cli.Name("app")
	cli.Command("public", "public command", func() {}).AsHidden()

	help := cli.GetHelp()
	assertFalse(t, strings.Contains(help, "Commands:"))
	assertFalse(t, strings.Contains(help, "public"))
	assertFalse(t, strings.Contains(help, "<command>"))
}

func TestHiddenSubcommandStillExecutes(t *testing.T) {
	defer cli.Clear()

	executed := false
	cli.UsePanic(true)
	cli.Command("internal", "internal command", func() {
		executed = true
		cli.ParseArgs(make_args())
	}).AsHidden()

	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("internal"))
	}, 0)
	assertTrue(t, executed)
}
