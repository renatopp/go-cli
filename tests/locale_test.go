package tests

import (
	"strings"
	"testing"

	"github.com/renatopp/go-cli"
	"github.com/renatopp/go-cli/pkg/locales"
)

func TestLocaleCustomErrorMessage(t *testing.T) {
	defer cli.Clear()
	defer cli.Locale(locales.Locale{})

	cli.Locale(locales.Locale{
		ErrUnknownFlag: "bandeira desconhecida %s",
	})
	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "bandeira desconhecida"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("--a", "1"))
	}, 1)
}

func TestLocaleFallsBackToDefault(t *testing.T) {
	defer cli.Clear()
	defer cli.Locale(locales.Locale{})

	// Only override one field; the rest should keep the English defaults.
	cli.Locale(locales.Locale{
		ErrMissingRequiredFlag: "falta a bandeira obrigatória %s",
	})
	cli.UsePanic(true)
	cli.Stderr(printfContains(t, "unknown flag"))
	expectPanicWith(t, func() {
		cli.ParseArgs(make_args("--a", "1"))
	}, 1)
}

func TestLocaleHelpLabels(t *testing.T) {
	defer cli.Clear()
	defer cli.Locale(locales.Locale{})

	cli.Locale(locales.Locale{
		UsageLabel: "Uso",
	})
	cli.FlagString("name", "n", "your name")
	help := cli.GetHelp()
	assertTrue(t, strings.Contains(help, "Uso:"), "expected help to contain translated usage label")
}
