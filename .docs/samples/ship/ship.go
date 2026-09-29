package main

import (
	"fmt"

	"github.com/renatopp/go-cli"
)

func main() {
	cli.Name("ship")
	cli.Description("Shell completion sample. Run `source <(ship completion bash)` to enable it.")
	cli.AutoCompletion(true)
	cli.AutoHelp(true)
	cli.FlagBool("verbose", "v", "Verbose output.").AsGlobal()

	cli.Command("deploy", "Deploy a service.", func() {
		env := cli.Flag("env", "e", "Target environment.").
			WithCompletion(func(prefix string) []string {
				return []string{"dev\tDevelopment", "staging\tStaging", "prod\tProduction"}
			})
		out := cli.Flag("output", "o", "Output file.")
		service := cli.Pos("service", "Service name.").AsRequired().
			WithCompletion(func(prefix string) []string {
				return []string{"api", "web", "worker"}
			})
		cli.Parse()
		fmt.Println("deploy", service.Value(), env.Value(), out.Value())
	})

	cli.Command("status", "Show the status.", func() {
		fmt.Println("ok")
	})

	cli.Parse()
	cli.Help()
}
