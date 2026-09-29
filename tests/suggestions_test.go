package tests

import (
	"strings"
	"testing"

	"github.com/renatopp/go-cli"
	"github.com/renatopp/go-cli/locales"
)

func parseStderr(t *testing.T, args ...any) string {
	var stderr strings.Builder
	cli.UsePanic(true)
	cli.Stderr(&stderr)
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args(args...))
	}, 1)
	return stderr.String()
}

func TestSuggestionUnknownFlag(t *testing.T) {
	defer cli.Clear()
	cli.FlagBool("verbose", "V", "verbose output")
	out := parseStderr(t, "--verbos")
	assertTrue(t, strings.Contains(out, "unknown flag verbos, did you mean --verbose?"), out)
}

func TestSuggestionUnknownFlagTooFar(t *testing.T) {
	defer cli.Clear()
	cli.FlagBool("verbose", "V", "verbose output")
	out := parseStderr(t, "--xyz")
	assertTrue(t, strings.Contains(out, "unknown flag xyz"), out)
	assertFalse(t, strings.Contains(out, "did you mean"), out)
}

func TestSuggestionSkipsShortFlags(t *testing.T) {
	defer cli.Clear()
	cli.FlagBool("verbose", "V", "verbose output")
	out := parseStderr(t, "-x")
	assertFalse(t, strings.Contains(out, "did you mean"), out)
}

func TestSuggestionSkipsHiddenFlags(t *testing.T) {
	defer cli.Clear()
	cli.FlagBool("verbose", "V", "verbose output").AsHidden()
	out := parseStderr(t, "--verbos")
	assertFalse(t, strings.Contains(out, "did you mean"), out)
}

func TestSuggestionGlobalFlag(t *testing.T) {
	defer cli.Clear()
	cli.FlagBool("verbose", "V", "verbose output").AsGlobal()
	cli.Command("commit", "commit changes", func() {
		cli.Parse()
	})
	out := parseStderr(t, "commit", "--verbos")
	assertTrue(t, strings.Contains(out, "did you mean --verbose?"), out)
}

func TestSuggestionUnknownCommand(t *testing.T) {
	defer cli.Clear()
	cli.Command("commit", "commit changes", func() {})
	cli.Command("push", "push changes", func() {})
	out := parseStderr(t, "comit")
	assertTrue(t, strings.Contains(out, "unknown command comit, did you mean commit?"), out)
}

func TestSuggestionUnknownCommandTooFar(t *testing.T) {
	defer cli.Clear()
	cli.Command("commit", "commit changes", func() {})
	out := parseStderr(t, "xyz")
	assertTrue(t, strings.Contains(out, "unknown command xyz"), out)
	assertFalse(t, strings.Contains(out, "did you mean"), out)
}

func TestSuggestionSkipsHiddenCommands(t *testing.T) {
	defer cli.Clear()
	cli.Command("commit", "commit changes", func() {}).AsHidden()
	out := parseStderr(t, "comit")
	assertTrue(t, strings.Contains(out, "unexpected extra positional argument: comit"), out)
	assertFalse(t, strings.Contains(out, "did you mean"), out)
}

func TestSuggestionOnlyFirstPositionalIsCommand(t *testing.T) {
	defer cli.Clear()
	cli.Command("commit", "commit changes", func() {})
	cli.Pos("name", "a name")
	out := parseStderr(t, "a", "comit")
	assertTrue(t, strings.Contains(out, "unexpected extra positional argument: comit"), out)
}

func TestSuggestionDisabled(t *testing.T) {
	defer cli.Clear()
	cli.Suggestions(false)
	cli.FlagBool("verbose", "V", "verbose output")
	cli.Command("commit", "commit changes", func() {})
	out := parseStderr(t, "--verbos")
	assertFalse(t, strings.Contains(out, "did you mean"), out)
	out = parseStderr(t, "comit")
	assertTrue(t, strings.Contains(out, "unknown command comit"), out)
	assertFalse(t, strings.Contains(out, "did you mean"), out)
}

func TestSuggestionLocalized(t *testing.T) {
	defer cli.Clear()
	cli.Locale(locales.PT_BR())
	cli.Command("commit", "commit changes", func() {})
	out := parseStderr(t, "comit")
	assertTrue(t, strings.Contains(out, "comando desconhecido comit, você quis dizer commit?"), out)
}
