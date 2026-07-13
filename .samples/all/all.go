package main

import "github.com/renatopp/go-cli"

func main() {
	cli.Name("all")
	cli.Description("This is a sample description lorem	 ipsum dolor sit amet, consectetur adipiscing elit. \tSed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
	cli.Flag("longs", "s", "string flag").AsGlobal().AsRepeatable().AsRequired()
	cli.FlagInt("longi", "i", "int flag").WithDefault(42).WithValidation(func(s int) error { return nil })
	cli.FlagBool("longb", "b", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi volutpat leo in neque efficitur finibus. Fusce imperdiet ut nulla a posuere. Nulla accumsan pellentesque gravida. Integer sed tortor et sapien vehicula rutrum tempus convallis nibh. Donec faucibus lacinia quam, id malesuada nunc gravida sed. Mauris eget pellentesque est. Donec sit amet nunc eu elit vehicula rhoncus eu ac dolor. ").WithEnv("BOOLEAN").AsRepeatable()
	cli.FlagDuration("longd", "d", "duration flag").AsHidden()
	cli.Pos("sample", "sample positional").
		WithDefault("default").
		WithValidation(func(s string) error { return nil })
	cli.Pos("files", "variadic positional").
		AsRequired().
		AsVariadic()
	cli.Command("commit", "commit command", func() {})
	cli.Command("push", "push command", func() {})
	cli.Command("pull", "pull command", func() {})
	cli.Command("commit2", "hidden command", func() {}).AsHidden()
	cli.Example("all --longs value xxx", "sample example 1")
	cli.Example("all --longs value yyy", "sample example 2")
	cli.AutoHelp(true)
	cli.Parse()
	cli.Help()
}
