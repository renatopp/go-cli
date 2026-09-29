package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/renatopp/go-cli"
)

func complete(t *testing.T, setup func(), args ...string) []string {
	t.Helper()
	defer cli.Clear()

	var buf bytes.Buffer
	cli.Name("app")
	cli.Stdout(&buf)
	cli.UsePanic(true)
	cli.AutoCompletion(true)
	setup()

	expectPanicWith(t, func() {
		cli.ParseArgs(append([]string{"__complete"}, args...))
	}, 0)

	out := strings.TrimSpace(buf.String())
	if out == "" {
		return nil
	}
	return strings.Split(out, "\n")
}

func completionTree() {
	cli.FlagBool("verbose", "v", "Verbose output.").AsGlobal()
	cli.Command("deploy", "Deploy the app.", func() {
		env := cli.Flag("env", "e", "Target environment.")
		env.WithCompletion(func(string) []string { return []string{"dev", "prod\tProduction"} })
		cli.Flag("output", "o", "Output file.")
		cli.FlagBool("dry", "d", "Dry run.")
		cli.Pos("service", "Service name.").WithCompletion(func(prefix string) []string {
			return []string{"api", "web", "worker"}
		})
		cli.Pos("files", "Files.").AsVariadic()
		cli.Parse()
	})
	cli.Command("delete", "Delete the app.", func() {})
	cli.Command("secret", "Hidden.", func() {}).AsHidden()
}

func TestCompletionSubcommands(t *testing.T) {
	got := complete(t, completionTree, "de")
	assertEqual(t, "deploy\tDeploy the app.|delete\tDelete the app.", strings.Join(got, "|"))
}

func TestCompletionHidesHiddenCommands(t *testing.T) {
	got := complete(t, completionTree, "")
	assertEqual(t, 2, len(got))
}

func TestCompletionFlags(t *testing.T) {
	got := complete(t, completionTree, "deploy", "--")
	assertEqual(t, "--verbose\tVerbose output.|--env\tTarget environment.|--output\tOutput file.|--dry\tDry run.", strings.Join(got, "|"))
}

func TestCompletionSkipsUsedFlags(t *testing.T) {
	got := complete(t, completionTree, "deploy", "-vd", "--env=dev", "--")
	assertEqual(t, "--output\tOutput file.", strings.Join(got, "|"))
}

func TestCompletionFlagValue(t *testing.T) {
	assertEqual(t, "prod\tProduction", strings.Join(complete(t, completionTree, "deploy", "--env", "p"), "|"))
	assertEqual(t, "dev|prod\tProduction", strings.Join(complete(t, completionTree, "deploy", "-e", ""), "|"))
	assertEqual(t, "--env=dev", strings.Join(complete(t, completionTree, "deploy", "--env=d"), "|"))
}

func TestCompletionFlagValueFallsBackToFiles(t *testing.T) {
	assertEqual(t, ":files", strings.Join(complete(t, completionTree, "deploy", "-o", ""), "|"))
}

func TestCompletionBoolFlagValue(t *testing.T) {
	assertEqual(t, "--dry=true|--dry=false", strings.Join(complete(t, completionTree, "deploy", "--dry="), "|"))
}

func TestCompletionPositionals(t *testing.T) {
	assertEqual(t, "web|worker", strings.Join(complete(t, completionTree, "deploy", "-e", "dev", "w"), "|"))
	assertEqual(t, ":files", strings.Join(complete(t, completionTree, "deploy", "api", ""), "|"))
	assertEqual(t, ":files", strings.Join(complete(t, completionTree, "deploy", "api", "a.txt", ""), "|"))
}

func TestCompletionCommandWithoutParse(t *testing.T) {
	assertEqual(t, 0, len(complete(t, completionTree, "delete", "")))
}

func TestCompletionShells(t *testing.T) {
	assertEqual(t, "bash|zsh", strings.Join(complete(t, completionTree, "completion", ""), "|"))
}

func TestCompletionScript(t *testing.T) {
	defer cli.Clear()

	for _, shell := range []string{"bash", "zsh"} {
		var buf bytes.Buffer
		cli.Clear()
		cli.Name("app")
		cli.Stdout(&buf)
		cli.UsePanic(true)
		cli.AutoCompletion(true)
		expectPanicWith(t, func() { cli.ParseArgs(make_args("completion", shell)) }, 0)
		assertTrue(t, strings.Contains(buf.String(), "__complete"), shell)
		assertTrue(t, strings.Contains(buf.String(), "_app"), shell)
	}
}

func TestCompletionDisabledByDefault(t *testing.T) {
	defer cli.Clear()
	cli.Name("app")
	cli.AllowExtraPos(true)
	res := cli.ParseArgs(make_args("__complete", "x"))
	assertEqual(t, "__complete", res.PosAt(0))
}
