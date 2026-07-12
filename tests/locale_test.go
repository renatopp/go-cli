package tests

import (
	"strings"
	"testing"

	"github.com/renatopp/go-cli"
	"github.com/renatopp/go-cli/locales"
)

func TestLocaleCustomErrorMessage(t *testing.T) {
	defer cli.Clear()
	defer cli.SetLocale(locales.Locale{})

	cli.SetLocale(locales.Locale{
		ErrUnknownFlag: "bandeira desconhecida %s",
	})
	cli.UsePanicInsteadOfExit(true)
	cli.SetStderr(printfContains(t, "bandeira desconhecida"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("--a", "1"))
	}, 1)
}

func TestLocaleFallsBackToDefault(t *testing.T) {
	defer cli.Clear()
	defer cli.SetLocale(locales.Locale{})

	// Only override one field; the rest should keep the English defaults.
	cli.SetLocale(locales.Locale{
		ErrMissingRequiredFlag: "falta a bandeira obrigatória %s",
	})
	cli.UsePanicInsteadOfExit(true)
	cli.SetStderr(printfContains(t, "unknown flag"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("--a", "1"))
	}, 1)
}

func TestLocaleHelpLabels(t *testing.T) {
	defer cli.Clear()
	defer cli.SetLocale(locales.Locale{})

	cli.SetLocale(locales.Locale{
		UsageLabel: "Uso",
	})
	cli.Flag("name", "n", "your name")
	help := cli.HelpString()
	assertTrue(t, strings.Contains(help, "Uso:"), "expected help to contain translated usage label")
}
