package tests

import (
	"strings"
	"testing"

	"github.com/renatopp/go-cli"
)

func TestLocaleCustomErrorMessage(t *testing.T) {
	defer cli.Clear()
	defer cli.SetLocale(cli.Locale{})

	cli.SetLocale(cli.Locale{
		ErrUnknownFlag: "bandeira desconhecida %s",
	})
	cli.UsePanicInsteadOfExit(true)
	cli.StderrWith(printfContains(t, "bandeira desconhecida"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("--a", "1"))
	}, 1)
}

func TestLocaleFallsBackToDefault(t *testing.T) {
	defer cli.Clear()
	defer cli.SetLocale(cli.Locale{})

	// Only override one field; the rest should keep the English defaults.
	cli.SetLocale(cli.Locale{
		ErrMissingRequiredFlag: "falta a bandeira obrigatória %s",
	})
	cli.UsePanicInsteadOfExit(true)
	cli.StderrWith(printfContains(t, "unknown flag"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("--a", "1"))
	}, 1)
}

func TestLocaleHelpLabels(t *testing.T) {
	defer cli.Clear()
	defer cli.SetLocale(cli.Locale{})

	cli.SetLocale(cli.Locale{
		UsageLabel: "Uso",
	})
	cli.Flag("name", "n", "your name")
	help := cli.HelpString()
	assertTrue(t, strings.Contains(help, "Uso:"), "expected help to contain translated usage label")
}
