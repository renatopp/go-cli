package main

import "github.com/renatopp/go-cli"

func main() {
	cli.Name("stub")
	cli.Description("This is a stub cli to showcase the library features.")
	cli.AutoHelp()

	cli.Command("version", "Prints the version of the application.", cmdVersion)
	cli.Command("hello", "Prints Hello, World!", cmdHello)
	cli.Command("repeat", "Repeat the provided message.", cmdRepeat)
	cli.Command("booleans", "Demonstrates boolean flags.", cmdBooleans)
	cli.Command("restricted", "Demonstrates restricted commands.", cmdRestricted)
	cli.Parse()

	cli.ShowHelp()
}

func cmdVersion() {
	cli.Name("version")
	cli.Description("Prints the version of the application.")
	asJson := cli.FlagBool("json", "", "Prints the version in JSON format.")
	cli.Parse()

	if asJson.Value {
		println(`{"version": "1.0.0"}`)
		return
	}

	println("1.0.0")
}

func cmdHello() {
	cli.Name("hello")
	cli.Description("Prints Hello, World!")
	name := cli.PosString(0, "name", "The name to greet.")
	cli.Parse()

	if name.IsSet() {
		println("Hello, " + name.Value + "!")
		return
	}

	println("Hello, World!")
}

func cmdRepeat() {
	cli.Name("repeat")
	cli.Description("Repeat the provided message.")
	message := cli.PosString(0, "message", "The message to repeat.").AsRequired()
	times := cli.FlagInt("times", "t", "Number of times to repeat the message.").WithDefault(2)
	cli.Parse()

	for i := 0; i < times.Value; i++ {
		println(message.Value)
	}
}

func cmdBooleans() {
	cli.Name("booleans")
	cli.Description("Demonstrates boolean flags.")
	flagA := cli.FlagBool("", "a", "An example boolean flag.")
	flagB := cli.FlagBool("", "b", "An example boolean flag.")
	flagC := cli.FlagBool("", "c", "An example boolean flag.")
	flagLong := cli.FlagBool("long", "", "An example boolean flag.")
	cli.Parse()

	println("Flag A:   ", flagA.Value)
	println("Flag B:   ", flagB.Value)
	println("Flag C:   ", flagC.Value)
	println("Flag Long:", flagLong.Value)
}

func cmdRestricted() {
	cli.Name("restricted")
	cli.Description("Demonstrates restricted commands.")
	cli.Restricted()
	cli.Command("a", "Show a", func() { println("a") })
	cli.Command("b", "Show b", func() { println("b") })
	cli.Command("c", "Show c", func() { println("c") })
	cli.Parse()
}
