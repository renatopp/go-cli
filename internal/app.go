package internal

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type App struct {
	originalRaw []string
	currentRaw  []string
	path        []string
	arguments   *Arguments
	cmd         *Command
	strict      bool
	autoHelp    bool
	printf      func(format string, a ...any)

	initialized bool
	parsed      bool
}

func (a *App) WithStrict(strict bool) *App {
	a.strict = strict
	return a
}

func (a *App) WithAutoHelp(autoHelp bool) *App {
	a.autoHelp = autoHelp
	return a
}

func (a *App) WithPrintFunc(printf func(format string, a ...any)) *App {
	a.printf = printf
	return a
}

func (a *App) NArgs() int {
	return len(a.arguments.Args)
}

func (a *App) Arg(i int) string {
	if i < 0 || i >= a.NArgs() {
		return ""
	}
	return a.arguments.Args[i]
}

func (a *App) Args() []string {
	return slices.Clone(a.arguments.Args)
}

func (a *App) Help() {
	name := strings.Join(a.path, " ")

	cmds := ""
	if len(a.cmd.subcommands) > 0 {
		cmds = " <command>"
	}

	opts := ""
	if len(a.cmd.flags) > 0 {
		opts = " [options]"
	}

	positionals := " "
	for _, p := range a.cmd.positionals {
		if p.IsRequired() {
			positionals += "<" + p.Name() + ">"
			continue
		}
		positionals += "[" + p.Name() + "]"
	}

	writer := NewDefaultTypewriter()
	writer.WriteLine("Usage: %s%s%s%s", name, cmds, opts, positionals)
	if a.cmd.description != "" {
		writer.WriteLine("\n%s", a.cmd.description)
	}

	if len(a.cmd.flags) > 0 {
		writer.WriteLine("")
		writer.WriteLine("Options:")
		for _, f := range a.cmd.flags {
			opts := ""
			if f.Short() != "" && f.Long() != "" {
				opts = fmt.Sprintf("-%s, --%s", f.Short(), f.Long())
			} else if f.Short() != "" {
				opts = fmt.Sprintf("-%s", f.Short())
			} else if f.Long() != "" {
				opts = fmt.Sprintf("--%s", f.Long())
			}
			desc := f.Description()
			if f.IsRequired() {
				desc += " Required."
			}
			writer.WriteLine("  %s\t%s", opts, desc)
		}
	}

	if len(a.cmd.subcommands) > 0 {
		writer.WriteLine("")
		writer.WriteLine("Commands:")
		for _, cmd := range a.cmd.subcommands {
			writer.WriteLine("  %s\t%s", cmd.name, cmd.shortDescription)
		}
	}
	s := writer.Flush()
	a.printf(s[:len(s)-1])
}

func (a *App) Parse() {
	if !a.initialized {
		a.originalRaw = os.Args
		a.currentRaw = os.Args[1:]
		a.path = []string{}
		if a.cmd.name != "" {
			a.path = append(a.path, a.cmd.name)
		} else {
			a.path = append(a.path, a.originalRaw[0])
		}
		a.initialized = true
	}

	if len(a.currentRaw) > 0 {
		name := a.currentRaw[0]
		for _, cmd := range a.cmd.subcommands {
			if cmd.name == name {
				a.currentRaw = a.currentRaw[1:]
				a.cmd = cmd
				a.path = append(a.path, name)
				cmd.execute()
				os.Exit(0)
			}
		}
	}
	if a.parsed {
		return
	}
	a.parseArguments()
	a.parsed = true
}

func (a *App) parseArguments() {
	args, err := parseArguments(a.currentRaw, a.cmd.positionals, a.cmd.flags, a.strict)
	if args != nil && args.HasHelp && a.autoHelp {
		a.Help()
		os.Exit(0)
	}
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	a.arguments = args

}
