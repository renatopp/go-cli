package formatters

import (
	"fmt"
	"slices"
	"strings"

	"github.com/renatopp/go-cli/core"
	"github.com/renatopp/go-cli/locales"
)

// DefaultHelpFormatter is the built-in help style. It renders the usage line,
// description, and the visible subcommands, options and arguments of the
// command, using the active locale for labels.
func DefaultHelpFormatter(cmd *core.Command, loc locales.Locale) string {

	var (
		path        = strings.Join(cmd.Path(), " ") // joined commands -- "git commit"
		subcommands = cmd.Subcommands()
		flags       = cmd.Flags()
		positionals = cmd.Positionals()
		examples    = cmd.Examples()
	)

	var (
		hasVisibleSubcommands = slices.ContainsFunc(subcommands, func(c *core.Command) bool { return !c.IsHidden() })
		hasVisibleFlags       = slices.ContainsFunc(flags, func(f core.AnyFlag) bool { return !f.IsHidden() })
		hasVisiblePositionals = slices.ContainsFunc(positionals, func(p core.AnyPositional) bool { return !p.IsHidden() })
		hasExamples           = len(examples) > 0
	)

	var (
		writer = newTabwriter()
	)

	// Usage
	{
		var (
			cmds = ""
			opts = ""
		)

		if hasVisibleSubcommands {
			cmds = fmt.Sprintf(" <%s>", loc.UsageCommandLabel)
		}
		if hasVisibleFlags {
			opts = fmt.Sprintf(" [%s]", loc.UsageOptionsLabel)
		}

		var pos strings.Builder
		for _, p := range positionals {
			if p.IsHidden() {
				continue
			}

			if p.IsRequired() {
				fmt.Fprintf(&pos, " <%s>", p.Name())
			} else {
				fmt.Fprintf(&pos, " [<%s>]", p.Name())
			}

			if p.IsVariadic() {
				fmt.Fprintf(&pos, "...")
			}
		}

		writer.WriteLine("%s %s%s%s%s", titleStyle(loc.UsageLabel), path, cmds, opts, pos.String())
	}

	// Description
	if cmd.Description() != "" {
		writer.WriteLine("\n%s", descriptionStyle(cmd.Description()))
	}

	// Subcommands
	if hasVisibleSubcommands {
		writer.WriteLine("")
		writer.WriteLine("%s", titleStyle(loc.CommandsLabel))
		for _, sub := range subcommands {
			if sub.IsHidden() {
				continue
			}
			writer.WriteLine("  %s\t%s",
				argStyle(sub.Name()),
				shortDescriptionStyle(sub.ShortDescription()),
			)
		}
	}

	// Flags
	if hasVisibleFlags {
		writer.WriteLine("")
		writer.WriteLine("%s", titleStyle(loc.OptionsLabel))
		for _, f := range flags {
			if f.IsHidden() {
				continue
			}

			opts := f.Signature()
			desc := f.Description()
			tags := make([]string, 0, 3)
			if f.IsGlobal() {
				tags = append(tags, loc.TagGlobalLabel)
			}
			if f.IsRequired() {
				tags = append(tags, loc.TagRequiredLabel)
			}
			if f.HasDefault() {
				tags = append(tags, fmt.Sprintf(loc.TagDefaultLabel, f.RawDefault()))
			}
			if f.HasEnv() {
				tags = append(tags, fmt.Sprintf(loc.TagEnvLabel, f.Env()))
			}
			tagsString := ""
			if len(tags) > 0 {
				tagsString = fmt.Sprintf(" (%s)", strings.Join(tags, ", "))
			}
			writer.WriteLine("  %s\t%s%s",
				argStyle(opts),
				shortDescriptionStyle(desc),
				tagStyle(tagsString),
			)
		}
	}

	// Positionals
	if hasVisiblePositionals {
		writer.WriteLine("")
		writer.WriteLine("%s", titleStyle(loc.ArgumentsLabel))
		for _, p := range positionals {
			if p.IsHidden() {
				continue
			}

			desc := p.Description()
			tags := make([]string, 0, 3)
			if p.IsRequired() {
				tags = append(tags, loc.TagRequiredLabel)
			}
			if p.HasDefault() {
				tags = append(tags, fmt.Sprintf(loc.TagDefaultLabel, p.RawDefault()))
			}
			if p.HasEnv() {
				tags = append(tags, fmt.Sprintf(loc.TagEnvLabel, p.Env()))
			}
			tagsString := ""
			if len(tags) > 0 {
				tagsString = fmt.Sprintf(" (%s)", strings.Join(tags, ", "))
			}
			writer.WriteLine("  %s\t%s%s",
				argStyle(p.Name()),
				shortDescriptionStyle(desc),
				tagStyle(tagsString),
			)
		}
	}

	// Examples
	if hasExamples {
		writer.WriteLine("")
		writer.WriteLine("%s", titleStyle(loc.ExamplesLabel))
		for i, ex := range cmd.Examples() {
			if i > 0 {
				writer.WriteLine("")
			}
			writer.WriteLine("  %s", argStyle(ex.Usage))
			writer.WriteLine("  %s", shortDescriptionStyle(ex.Description))
		}
	}

	return writer.Flush()
}
