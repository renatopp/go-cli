package internal

import "slices"

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
	if flag.Long() != "" && c.HasFlag(flag.Long()) {
		panic("flag with the same long name already exists: " + flag.Long())
	}
	if flag.Short() != "" && c.HasFlag(flag.Short()) {
		panic("flag with the same short name already exists: " + flag.Short())
	}

	c.flags = append(c.flags, flag)
	return c
}

func (c *Command) WithSubcommand(cmd *Command) *Command {
	if slices.IndexFunc(c.subcommands, func(sc *Command) bool { return sc.name == cmd.name }) != -1 {
		panic("subcommand with the same name already exists: " + cmd.name)
	}
	c.subcommands = append(c.subcommands, cmd)
	return c
}

func (c *Command) HasFlag(n string) bool {
	for _, f := range c.flags {
		if f.Long() == n || f.Short() == n {
			return true
		}
	}
	return false
}
