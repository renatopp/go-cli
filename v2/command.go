package v2

type Command struct {
	name             string
	description      string
	shortDescription string
	execute          func()
	positionals      []Positional
	flags            []Flag
	subcommands      []*Command
}

func NewCommand() *Command {
	return &Command{
		name:        "",
		description: "",
		execute:     nil,
		positionals: []Positional{},
		flags:       []Flag{},
		subcommands: []*Command{},
	}
}

func (c *Command) WithName(name string) *Command {
	c.name = name
	return c
}

func (c *Command) WithDescription(description string) *Command {
	c.description = description
	return c
}

func (c *Command) WithShortDescription(shortDescription string) *Command {
	c.shortDescription = shortDescription
	return c
}

func (c *Command) WithExecute(execute func()) *Command {
	c.execute = execute
	return c
}

func (c *Command) WithPositional(arg Positional) *Command {
	if len(c.positionals) > 0 {
		last := c.positionals[len(c.positionals)-1]
		if last.IsVariadic() {
			panic("cannot add a positional argument after a variadic positional argument")
		}
	}

	c.positionals = append(c.positionals, arg)
	return c
}

func (c *Command) WithFlag(flag Flag) *Command {
	c.flags = append(c.flags, flag)
	return c
}

func (c *Command) WithSubcommand(cmd *Command) *Command {
	c.subcommands = append(c.subcommands, cmd)
	return c
}
