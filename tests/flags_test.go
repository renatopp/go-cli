package tests

import (
	"strings"
	"testing"

	"github.com/renatopp/go-cli"
)

func TestFlagInvalidExtra(t *testing.T) {
	defer cli.Clear()
	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "unknown flag"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("--a", "1"))
	}, 1)
}

func TestFlagExtra(t *testing.T) {
	defer cli.Clear()
	cli.AllowExtraFlags(true)
	cli.ParseArgs(make_args("--a"))
}

func TestFlagAssignedLong(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagString("long", "", "")
	cli.ParseArgs(make_args("--long=1"))
	assertEqual(t, a.Value(), "1")
}

func TestFlagUnassignedLong(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagString("long", "", "")
	cli.ParseArgs(make_args("--long", "1"))
	assertEqual(t, a.Value(), "1")
}

func TestFlagBoolean(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagBool("long", "", "")
	cli.AllowExtraPos(true)
	arguments := cli.ParseArgs(make_args("--long", "1"))
	assertEqual(t, a.Value(), true)
	assertEqual(t, arguments.ExtraPosAt(0), "1")
}

func TestFlagLongWithDashedValue(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagString("long", "", "")
	cli.ParseArgs(make_args("--long", "--name"))
	assertEqual(t, a.Value(), "--name")
}

func TestFlagInvalidLongWithoutValue(t *testing.T) {
	defer cli.Clear()
	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "missing value for flag"))
	cli.FlagString("long", "", "")
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("--long"))
	}, 1)
}

func TestFlagShortUncombined(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagString("", "s", "")
	cli.ParseArgs(make_args("-s", "1"))
	assertEqual(t, a.Value(), "1")
}

func TestFlagShortCombinedValue(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagString("", "s", "")
	cli.ParseArgs(make_args("-s1"))
	assertEqual(t, a.Value(), "1")
}

func TestFlagShortBoolean(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagBool("", "s", "")
	cli.AllowExtraPos(true)
	arguments := cli.ParseArgs(make_args("-s", "1"))
	assertEqual(t, a.Value(), true)
	assertEqual(t, arguments.ExtraPosAt(0), "1")
}

func TestFlagShortCombinedDashedValue(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagString("", "s", "")
	cli.ParseArgs(make_args("-s--name"))
	assertEqual(t, a.Value(), "--name")
}

func TestFlagInvalidShortWithoutValue(t *testing.T) {
	defer cli.Clear()
	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "missing value for flag"))
	cli.FlagString("", "s", "")
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("-s"))
	}, 1)
}

func TestFlagCombined(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagBool("", "a", "")
	b := cli.FlagBool("", "b", "")
	c := cli.FlagBool("", "c", "")
	cli.ParseArgs(make_args("-ab"))
	assertEqual(t, a.Value(), true)
	assertEqual(t, b.Value(), true)
	assertEqual(t, c.Value(), false)
}

func TestFlagCombinedWithValue(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagBool("", "a", "")
	b := cli.FlagString("", "b", "")
	c := cli.FlagBool("", "c", "")
	cli.ParseArgs(make_args("-ab1"))
	assertEqual(t, a.Value(), true)
	assertEqual(t, b.Value(), "1")
	assertEqual(t, c.Value(), false)
}

func TestFlagCombinedWithDashedValue(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagBool("", "a", "")
	b := cli.FlagString("", "b", "")
	c := cli.FlagBool("", "c", "")
	cli.ParseArgs(make_args("-ab--name"))
	assertEqual(t, a.Value(), true)
	assertEqual(t, b.Value(), "--name")
	assertEqual(t, c.Value(), false)
}

func TestFlagCombinedThreeBooleans(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagBool("", "a", "")
	b := cli.FlagBool("", "b", "")
	c := cli.FlagBool("", "c", "")
	cli.ParseArgs(make_args("-abc"))
	assertEqual(t, a.Value(), true)
	assertEqual(t, b.Value(), true)
	assertEqual(t, c.Value(), true)
}

func TestFlagCombinedWithSpaceSeparatedValue(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagBool("", "a", "")
	b := cli.FlagBool("", "b", "")
	c := cli.FlagString("", "c", "")
	cli.ParseArgs(make_args("-abc", "value"))
	assertEqual(t, a.Value(), true)
	assertEqual(t, b.Value(), true)
	assertEqual(t, c.Value(), "value")
}

func TestFlagInvalidRepeatedLong(t *testing.T) {
	defer cli.Clear()
	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "specified multiple times"))
	cli.FlagString("a", "", "")
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("--a", "1", "--a", "2"))
	}, 1)
}

func TestFlagInvalidRepeatedShort(t *testing.T) {
	defer cli.Clear()
	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "specified multiple times"))
	cli.FlagString("", "a", "")
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("-a", "1", "-a", "2"))
	}, 1)
}

func TestFlagRepeatedAllowedGlobally(t *testing.T) {
	defer cli.Clear()
	cli.AllowRepeatedFlags(true)
	a := cli.FlagString("a", "", "")
	cli.ParseArgs(make_args("--a", "1", "--a", "2"))
	assertEqual(t, a.Value(), "2")
	assertEqual(t, a.Values()[0], "1")
	assertEqual(t, a.Values()[1], "2")
	assertTrue(t, a.IsRepeated())
}

func TestFlagRepeatedAllowedLocally(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagString("a", "", "").AsRepeatable()
	cli.ParseArgs(make_args("--a", "1", "--a", "2"))
	assertEqual(t, a.Value(), "2")
	assertEqual(t, a.Values()[0], "1")
	assertEqual(t, a.Values()[1], "2")
	assertTrue(t, a.IsRepeated())
}

func TestFlagRepeatedCombined(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagBool("", "a", "").AsRepeatable()
	cli.ParseArgs(make_args("-aa"))
	assertEqual(t, a.Value(), true)
	assertEqual(t, len(a.Values()), 2)
	assertTrue(t, a.IsRepeated())
}

func TestFlagRequired(t *testing.T) {
	defer cli.Clear()
	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "missing required flag"))
	cli.FlagString("a", "", "").AsRequired()
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args())
	}, 1)
}

func TestFlagRequiredWithValue(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagString("a", "", "").AsRequired()
	cli.ParseArgs(make_args("--a", "1"))
	assertEqual(t, a.Value(), "1")
}

func TestFlagRequiredCombined(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagBool("", "a", "").AsRequired()
	b := cli.FlagBool("", "b", "")
	cli.ParseArgs(make_args("-ab"))
	assertEqual(t, a.Value(), true)
	assertEqual(t, b.Value(), true)
}

func TestFlagOptional(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagString("a", "", "").WithDefault("defaulted")
	cli.ParseArgs(make_args())
	assertEqual(t, a.Value(), "defaulted")
}

func TestFlagOptionalWithValue(t *testing.T) {
	defer cli.Clear()
	a := cli.FlagString("a", "", "").WithDefault("defaulted")
	cli.ParseArgs(make_args("--a", "1"))
	assertEqual(t, a.Value(), "1")
}

func TestFlagHiddenNotInHelpButParses(t *testing.T) {
	defer cli.Clear()

	hidden := cli.FlagString("secret", "s", "hidden flag").AsHidden()
	visible := cli.FlagString("name", "n", "visible flag")

	help := cli.GetHelp()
	assertFalse(t, strings.Contains(help, "--secret"))
	assertFalse(t, strings.Contains(help, "hidden flag"))
	assertTrue(t, strings.Contains(help, "--name"))

	cli.ParseArgs(make_args("--secret", "token", "--name", "john"))
	assertEqual(t, hidden.Value(), "token")
	assertEqual(t, visible.Value(), "john")
}

func TestFlagHiddenRequiredStillValidated(t *testing.T) {
	defer cli.Clear()

	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "missing required flag"))
	cli.FlagString("secret", "s", "hidden required flag").AsHidden().AsRequired()

	expectPanicWith(t, func() {
		cli.ParseArgs(make_args())
	}, 1)
}
