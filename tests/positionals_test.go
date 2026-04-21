package tests

import (
	"fmt"
	"testing"

	"github.com/renatopp/go-cli"
)

func TestPositionalBasic(t *testing.T) {
	defer cli.Clear()

	a := cli.Pos("a", "desc")
	b := cli.PosInt("b", "desc")
	c := cli.PosUint("c", "desc")
	cli.ParseArgs(make_args("1", "2"))

	assertEqual(t, "1", a.Value())
	assertEqual(t, 2, b.Value())
	assertEqual(t, uint(0), c.Value())
	assertEqual(t, cli.NArgs(), 2)
	assertEqual(t, cli.Arg(0), "1")
	assertEqual(t, cli.Arg(1), "2")
	assertEqual(t, cli.Arg(2), "")
	assertEqual(t, len(cli.Args()), 2)
	assertEqual(t, cli.NExtraArgs(), 0)
}

func TestPositionalWithExtra(t *testing.T) {
	defer cli.Clear()

	cli.AllowExtraPositionals(true)
	cli.Pos("a", "desc")
	cli.ParseArgs(make_args("1", "2", "3", "4"))

	assertEqual(t, cli.NExtraArgs(), 3)
	assertEqual(t, cli.ExtraArg(0), "2")
	assertEqual(t, cli.ExtraArg(1), "3")
	assertEqual(t, cli.ExtraArg(2), "4")
	assertEqual(t, len(cli.ExtraArgs()), 3)
}

func TestPositionalExtraNotAllowed(t *testing.T) {
	defer cli.Clear()

	cli.UsePanicInsteadOfExit(true)
	cli.AllowExtraPositionals(false) // default is false, but just to be explicit
	cli.Pos("a", "desc")
	cli.StderrWith(printfContains(t, "extra"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("1", "2", "3", "4"))
	}, 1)
}

func TestPositionalCustomValidation(t *testing.T) {
	defer cli.Clear()

	validFn := func(s string) error {
		if s == "1" {
			return nil
		}
		return fmt.Errorf("invalid value")
	}

	cli.UsePanicInsteadOfExit(true)
	a := cli.Pos("a", "desc").WithValidation(validFn)
	cli.Pos("b", "desc").WithValidation(validFn)
	cli.StderrWith(printfContains(t, "invalid value"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("1", "2"))
	}, 1)
	assertEqual(t, a.Value(), "1")
}

func TestPositionalRequired(t *testing.T) {
	defer cli.Clear()

	cli.UsePanicInsteadOfExit(true)
	cli.Pos("a", "desc").AsRequired()
	cli.Pos("b", "desc").AsRequired()
	cli.StderrWith(printfContains(t, ": b"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("1"))
	}, 1)
}

func TestPositionalOptional(t *testing.T) {
	defer cli.Clear()

	a := cli.Pos("a", "desc").AsRequired()
	b := cli.Pos("b", "desc").WithDefault("defaulted")
	cli.ParseArgs(make_args("1"))
	assertEqual(t, a.Value(), "1")
	assertEqual(t, b.Value(), "defaulted")
}

func TestPositionalVariadicWithMultipleValues(t *testing.T) {
	defer cli.Clear()

	a := cli.Pos("a", "desc").AsRequired()
	b := cli.Pos("b", "desc").AsRequired()
	c := cli.Pos("c", "desc").AsVariadic()
	cli.ParseArgs(make_args("1", "2", "3", "4"))

	assertEqual(t, cli.NArgs(), 4)
	assertEqual(t, cli.NExtraArgs(), 0)
	assertEqual(t, a.Value(), "1")
	assertEqual(t, b.Value(), "2")
	assertEqual(t, c.Values()[0], "3")
	assertEqual(t, c.Values()[1], "4")
}

func TestPositionalEmptyVariadic(t *testing.T) {
	defer cli.Clear()

	a := cli.Pos("a", "desc").AsRequired()
	b := cli.Pos("b", "desc").AsRequired()
	c := cli.Pos("c", "desc").AsVariadic()
	cli.ParseArgs(make_args("1", "2"))

	assertEqual(t, cli.NArgs(), 2)
	assertEqual(t, cli.NExtraArgs(), 0)
	assertEqual(t, a.Value(), "1")
	assertEqual(t, b.Value(), "2")
	assertEqual(t, len(c.Values()), 0)
}

func TestPositionalMultipleVariadic(t *testing.T) {
	defer cli.Clear()

	cli.Pos("a", "desc")
	cli.Pos("b", "desc").AsVariadic()

	expectPanicWith(t, func() {
		cli.Pos("c", "desc").AsVariadic()
	}, "cannot add a positional argument after a variadic positional argument")
}

func TestPositionalWithEndOfOption(t *testing.T) {
	defer cli.Clear()

	a := cli.Pos("a", "desc")
	b := cli.Pos("b", "desc").AsVariadic()

	cli.ParseArgs(make_args("1", "a", "--", "--not-a-flag", "3"))
	assertEqual(t, cli.NArgs(), 4)
	assertEqual(t, cli.NExtraArgs(), 0)
	assertEqual(t, a.Value(), "1")
	assertEqual(t, b.Values()[0], "a")
	assertEqual(t, b.Values()[1], "--not-a-flag")
	assertEqual(t, b.Values()[2], "3")
}

func TestPositionalWithSingleDash(t *testing.T) {
	defer cli.Clear()
	a := cli.Pos("a", "desc")
	cli.ParseArgs(make_args("-"))
	assertEqual(t, a.Value(), "-")
}

func TestPositionalWithNumber(t *testing.T) {
	defer cli.Clear()
	a := cli.Pos("a", "desc")
	cli.ParseArgs(make_args("-1"))
	assertEqual(t, a.Value(), "-1")
}
