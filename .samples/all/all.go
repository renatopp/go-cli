package all

import "github.com/renatopp/go-cli"

func main() {
	cli.Name("all")
	cli.Description("This is a sample description")
	cli.Flag("longs", "s", "string flag").AsGlobal().AsRepeatable().AsRequired()
	cli.FlagInt("longi", "i", "int flag").WithDefault(42).WithValidation(func(s int) error { return nil })
	cli.FlagBool("longb", "b", "bool flag").WithEnv("BOOLEAN").AsRepeatable()
	cli.FlagDuration("longd", "d", "duration flag").AsHidden()
	cli.Pos("sample", "sample positional").
		WithDefault("default").
		WithValidation(func(s string) error { return nil })
	cli.Pos("files", "variadic positional").
		AsRequired().
		AsVariadic()
	cli.Example("all --longs value xxx", "sample example 1")
	cli.Example("all --longs value yyy", "sample example 2")
	cli.Parse()
	cli.Help()
}
