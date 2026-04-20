package v2

type State struct {
	originalRaw    []string
	currentRaw     []string
	rootCommand    *Command
	currentCommand *Command
	strict         bool
	autoHelp       bool
	stdout         func(format string, a ...any)
}

type Command struct {
	name             string
	description      string
	shortDescription string
	execute          func()
	positionals      []Positional
	flags            []Flag
	subcommands      []*Command
}

type argument interface {
	Description() string
	RawValue() string
	Parse(value string) error
	IsParsed() bool
	IsRequired() bool
	IsPassed() bool
	HasDefault() bool
	RawDefault() string
	SetRawDefault(rawDefault string)
}

type Flag interface {
	argument
	Long() string
	Short() string
}

type Positional interface {
	argument
	Index() int
	Name() string
}
