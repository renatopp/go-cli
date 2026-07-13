package core

import (
	"errors"
	"slices"
)

type Command struct {
	parent           *Command
	name             string
	description      string
	shortDescription string
	hidden           bool
	execute          func()
	positionals      []AnyPositional
	flags            []AnyFlag
	subcommands      []*Command
}

func NewCommand(parent *Command) *Command {
	return &Command{
		parent:      parent,
		name:        "",
		description: "",
		execute:     nil,
		positionals: []AnyPositional{},
		flags:       []AnyFlag{},
		subcommands: []*Command{},
	}
}

// WithName sets the name of the command. The name is used to identify the command in
// the command hierarchy and in the help output.
func (c *Command) WithName(name string) *Command {
	c.name = name
	return c
}

// WithDescription sets the description of the command. The description is used in the
// help output for this command.
func (c *Command) WithDescription(description string) *Command {
	c.description = description
	return c
}

// WithShortDescription sets the short description of the command. The short description
// is used in the help output for the parent command, to describe this command in a single line.
func (c *Command) WithShortDescription(shortDescription string) *Command {
	c.shortDescription = shortDescription
	return c
}

// WithExecute sets the function to be executed when the command is invoked.
func (c *Command) WithExecute(execute func()) *Command {
	c.execute = execute
	return c
}

// WithPositional adds a positional argument to the command.
func (c *Command) WithPositional(arg AnyPositional) *Command {
	if len(c.positionals) > 0 {
		last := c.positionals[len(c.positionals)-1]
		if last.IsVariadic() {
			panic("cannot add a positional argument after a variadic positional argument")
		}
	}

	c.positionals = append(c.positionals, arg)
	return c
}

// WithFlag adds a flag to the command.
func (c *Command) WithFlag(flag AnyFlag) *Command {
	if flag.Long() != "" && c.HasFlag(flag.Long()) {
		panic("flag with the same long name already exists: " + flag.Long())
	}
	if flag.Short() != "" && c.HasFlag(flag.Short()) {
		panic("flag with the same short name already exists: " + flag.Short())
	}

	c.flags = append(c.flags, flag)
	return c
}

// WithSubcommand adds a subcommand to the command.
func (c *Command) WithSubcommand(cmd *Command) *Command {
	if slices.IndexFunc(c.subcommands, func(sc *Command) bool { return sc.name == cmd.name }) != -1 {
		panic("subcommand with the same name already exists: " + cmd.name)
	}
	c.subcommands = append(c.subcommands, cmd)
	return c
}

// AsHidden marks the command as hidden. Hidden commands are not shown in the help
// output bu can be executed directly by the user.
func (c *Command) AsHidden() *Command {
	c.hidden = true
	return c
}

// Parent returns the parent command, or nil for the root command.
func (c *Command) Parent() *Command {
	return c.parent
}

// Name returns the name of the command.
func (c *Command) Name() string {
	return c.name
}

// IsHidden returns true if the command is hidden.
func (c *Command) IsHidden() bool {
	return c.hidden
}

// Path returns the names of the commands leading to this command, starting at
// the root, e.g. ["git", "commit"].
func (c *Command) Path() []string {
	if c.parent == nil {
		return []string{c.name}
	}
	return append(c.parent.Path(), c.name)
}

// Description returns the description of the command.
func (c *Command) Description() string {
	return c.description
}

// ShortDescription returns the short description of the command.
func (c *Command) ShortDescription() string {
	return c.shortDescription
}

// Subcommands returns the list of subcommands of the command.
func (c *Command) Subcommands() []*Command {
	return c.subcommands[:]
}

// Positionals returns the list of positional arguments of the command.
func (c *Command) Positionals() []AnyPositional {
	return c.positionals[:]
}

// Flags returns the list of flags of the command.
func (c *Command) Flags() []AnyFlag {
	return c.flags[:]
}

// HasFlag returns true if the command has a flag with the given long or short name.
func (c *Command) HasFlag(n string) bool {
	for _, f := range c.flags {
		if f.Long() == n || f.Short() == n {
			return true
		}
	}
	return false
}

// GetFlag returns the flag with the given long or short name, or an error if not found.
func (c *Command) GetFlag(n string) (AnyFlag, error) {
	for _, f := range c.flags {
		if f.Long() == n || f.Short() == n {
			return f, nil
		}
	}
	return nil, errors.New("flag not found")
}

// inheritFlags adds the global flags of the parent command to the command's flags.
// The global flags are not added on this command creation because the parent command may
// not have been fully constructed yet.
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
