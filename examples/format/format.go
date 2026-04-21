package main

import "github.com/renatopp/go-cli"

func main() {
	cli.Name("format")
	cli.Description("Example of mutually exclusive options.")
	json := cli.FlagBool("json", "", "Output in JSON format")
	yaml := cli.FlagBool("yaml", "", "Output in YAML format")
	cli.Parse()

	// if you need at least one of the flags provided
	// cli.CheckAnyFlag(json, yaml)

	// if you need at most one of the flags provided
	cli.CheckExclusiveFlags(json, yaml)
	format := "plain text"
	if json.Value() {
		format = "JSON"
	} else if yaml.Value() {
		format = "YAML"
	}

	println("Output format:", format)
}
