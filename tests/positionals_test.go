package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/renatopp/go-cli"
)

func TestPositionalBasic(t *testing.T) {
	defer cli.Clear()

	a := cli.Pos("a", "desc")
	b := cli.PosInt("b", "desc")
	c := cli.PosUint("c", "desc")
	arguments := cli.ParseArgs(make_args("1", "2"))

	assertEqual(t, "1", a.Value())
	assertEqual(t, 2, b.Value())
	assertEqual(t, uint(0), c.Value())
	assertEqual(t, arguments.PosCount(), 2)
	assertEqual(t, arguments.PosAt(0), "1")
	assertEqual(t, arguments.PosAt(1), "2")
	assertEqual(t, arguments.PosAt(2), "")
	assertEqual(t, len(arguments.Pos()), 2)
	assertEqual(t, arguments.ExtraPosCount(), 0)
}

func TestPositionalWithExtra(t *testing.T) {
	defer cli.Clear()

	cli.AllowExtraPos(true)
	cli.Pos("a", "desc")
	arguments := cli.ParseArgs(make_args("1", "2", "3", "4"))

	assertEqual(t, arguments.ExtraPosCount(), 3)
	assertEqual(t, arguments.ExtraPosAt(0), "2")
	assertEqual(t, arguments.ExtraPosAt(1), "3")
	assertEqual(t, arguments.ExtraPosAt(2), "4")
	assertEqual(t, len(arguments.ExtraPos()), 3)
}

func TestPositionalExtraNotAllowed(t *testing.T) {
	defer cli.Clear()

	cli.UsePanic(true)
	cli.AllowExtraPos(false) // default is false, but just to be explicit
	cli.Pos("a", "desc")
	cli.Stderr(printfContains(t, "extra"))
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

	cli.UsePanic(true)
	a := cli.Pos("a", "desc").WithValidation(validFn)
	cli.Pos("b", "desc").WithValidation(validFn)
	cli.Stderr(printfContains(t, "invalid value"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("1", "2"))
	}, 1)
	assertEqual(t, a.Value(), "1")
}

func TestPositionalRequired(t *testing.T) {
	defer cli.Clear()

	cli.UsePanic(true)
	cli.Pos("a", "desc").AsRequired()
	cli.Pos("b", "desc").AsRequired()
	cli.Stderr(printfContains(t, ": b"))
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
	arguments := cli.ParseArgs(make_args("1", "2", "3", "4"))

	assertEqual(t, arguments.PosCount(), 4)
	assertEqual(t, arguments.ExtraPosCount(), 0)
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
	arguments := cli.ParseArgs(make_args("1", "2"))

	assertEqual(t, arguments.PosCount(), 2)
	assertEqual(t, arguments.ExtraPosCount(), 0)
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

	arguments := cli.ParseArgs(make_args("1", "a", "--", "--not-a-flag", "3"))
	assertEqual(t, arguments.PosCount(), 4)
	assertEqual(t, arguments.ExtraPosCount(), 0)
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

func TestPositionalHiddenNotInHelpButParses(t *testing.T) {
	defer cli.Clear()

	hidden := cli.Pos("secret", "hidden positional").AsHidden().AsRequired()
	visible := cli.Pos("name", "visible positional").AsRequired()

	help := cli.GetHelp()
	assertFalse(t, strings.Contains(help, "<secret>"))
	assertFalse(t, strings.Contains(help, "hidden positional"))
	assertTrue(t, strings.Contains(help, "<name>"))

	cli.ParseArgs(make_args("token", "john"))
	assertEqual(t, hidden.Value(), "token")
	assertEqual(t, visible.Value(), "john")
}

func TestPositionalHiddenRequiredStillValidated(t *testing.T) {
	defer cli.Clear()

	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "missing required positional argument"))
	cli.Pos("secret", "hidden required positional").AsHidden().AsRequired()

	expectPanicWith(t, func() {
		cli.ParseArgs(make_args())
	}, 1)
}

func TestPositionalEnvGetters(t *testing.T) {
	defer cli.Clear()
	a := cli.Pos("a", "desc")
	assertFalse(t, a.HasEnv())
	assertEqual(t, a.Env(), "")

	a.WithEnv("MY_POS_ENV")
	assertTrue(t, a.HasEnv())
	assertEqual(t, a.Env(), "MY_POS_ENV")
}

func TestPositionalEnvUsedWhenNotProvided(t *testing.T) {
	defer cli.Clear()
	t.Setenv("GO_CLI_TEST_POS_A", "from-env")

	a := cli.Pos("a", "desc").WithEnv("GO_CLI_TEST_POS_A")
	cli.ParseArgs(make_args())

	assertEqual(t, a.Value(), "from-env")
	assertTrue(t, a.IsProvided())
}

func TestPositionalEnvNotUsedWhenProvidedByUser(t *testing.T) {
	defer cli.Clear()
	t.Setenv("GO_CLI_TEST_POS_A", "from-env")

	a := cli.Pos("a", "desc").WithEnv("GO_CLI_TEST_POS_A")
	cli.ParseArgs(make_args("from-cli"))

	assertEqual(t, a.Value(), "from-cli")
}

func TestPositionalEnvFallsBackToDefaultWhenUnset(t *testing.T) {
	defer cli.Clear()

	a := cli.Pos("a", "desc").WithEnv("GO_CLI_TEST_POS_UNSET").WithDefault("defaulted")
	cli.ParseArgs(make_args())

	assertEqual(t, a.Value(), "defaulted")
	assertFalse(t, a.IsProvided())
}

func TestPositionalEnvSatisfiesRequired(t *testing.T) {
	defer cli.Clear()
	t.Setenv("GO_CLI_TEST_POS_REQUIRED", "from-env")

	a := cli.Pos("a", "desc").WithEnv("GO_CLI_TEST_POS_REQUIRED").AsRequired()
	cli.ParseArgs(make_args())

	assertEqual(t, a.Value(), "from-env")
}

// Regression test: resolving one positional's value from the environment must
// not block validation of the other required positionals.
func TestPositionalEnvDoesNotBlockOtherRequiredPositionals(t *testing.T) {
	defer cli.Clear()

	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "missing required positional argument"))
	cli.Pos("a", "desc").WithEnv("GO_CLI_TEST_POS_UNSET_2")
	cli.Pos("b", "desc").AsRequired()

	expectPanicWith(t, func() {
		cli.ParseArgs(make_args())
	}, 1)
}
