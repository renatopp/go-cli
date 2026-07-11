package internal

import "slices"

type Command struct {
	parent           *Command
	name             string
	description      string
	shortDescription string
	hidden           bool
	execute          func()
	positionals      []Positional
	flags            []Flag
	subcommands      []*Command
}

func NewCommand(parent *Command) *Command {
	return &Command{
		parent:      parent,
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

func (c *Command) AsHidden() *Command {
	c.hidden = true
	return c
}

func (c *Command) IsHidden() bool {
	return c.hidden
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

func (c *Command) Name() string {
	return c.name
}

func (c *Command) Description() string {
	return c.description
}

func (c *Command) ShortDescription() string {
	return c.shortDescription
}

func (c *Command) Commands() []*Command {
	return c.subcommands[:]
}

func (c *Command) Positionals() []Positional {
	return c.positionals[:]
}

func (c *Command) Flags() []Flag {
	return c.flags[:]
}

func (c *Command) HasFlag(n string) bool {
	for _, f := range c.flags {
		if f.Long() == n || f.Short() == n {
			return true
		}
	}
	return false
}

func (c *Command) GetFlag(n string) (Flag, error) {
	for _, f := range c.flags {
		if f.Long() == n || f.Short() == n {
			return f, nil
		}
	}
	return nil, ErrFlagNotFound
}

func (c *Command) inheritFlags() {
	if c.parent == nil {
		return
	}
	for _, f := range c.parent.flags {
		if f.IsGlobal() {
			c.flags = append(c.flags, f)
		}
	}
}
