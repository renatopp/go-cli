package formatters

// import (
// 	"fmt"
// 	"strings"
// )

// // DefaultHelpFormatter is the built-in help style. It renders the usage line,
// // description, and the visible subcommands, options and arguments of the
// // command, using the active locale for labels.
// func DefaultHelpFormatter(cmd *Command, loc Locale) string {
// 	name := strings.Join(cmd.Path(), " ")

// 	hasVisibleSubcommands := false
// 	for _, sub := range cmd.subcommands {
// 		if !sub.IsHidden() {
// 			hasVisibleSubcommands = true
// 			break
// 		}
// 	}

// 	hasVisibleFlags := false
// 	for _, f := range cmd.flags {
// 		if !f.IsHidden() {
// 			hasVisibleFlags = true
// 			break
// 		}
// 	}

// 	hasVisiblePositionals := false
// 	for _, p := range cmd.positionals {
// 		if !p.IsHidden() {
// 			hasVisiblePositionals = true
// 			break
// 		}
// 	}

// 	cmds := ""
// 	if hasVisibleSubcommands {
// 		cmds = fmt.Sprintf(" <%s>", loc.UsageCommandLabel)
// 	}

// 	opts := ""
// 	if hasVisibleFlags {
// 		opts = fmt.Sprintf(" [%s]", loc.UsageOptionsLabel)
// 	}

// 	var positionals strings.Builder
// 	for _, p := range cmd.positionals {
// 		if p.IsHidden() {
// 			continue
// 		}

// 		if p.IsRequired() {
// 			positionals.WriteString(" <")
// 			positionals.WriteString(p.Name())
// 			positionals.WriteString(">")
// 			continue
// 		}
// 		positionals.WriteString(" [<")
// 		positionals.WriteString(p.Name())
// 		positionals.WriteString(">]")
// 	}

// 	writer := newDefaultTypewriter()
// 	writer.WriteLine("%s: %s%s%s%s", loc.UsageLabel, name, cmds, opts, positionals.String())
// 	if cmd.description != "" {
// 		writer.WriteLine("\n%s", strings.TrimSpace(cmd.description))
// 	}

// 	if hasVisibleSubcommands {
// 		writer.WriteLine("")
// 		writer.WriteLine("%s:", loc.CommandsLabel)
// 		for _, sub := range cmd.subcommands {
// 			if sub.IsHidden() {
// 				continue
// 			}
// 			writer.WriteLine("  %s\t%s", sub.name, sub.shortDescription)
// 		}
// 	}

// 	if hasVisibleFlags {
// 		writer.WriteLine("")
// 		writer.WriteLine("%s:", loc.OptionsLabel)
// 		for _, f := range cmd.flags {
// 			if f.IsHidden() {
// 				continue
// 			}

// 			opts := f.Signature()
// 			desc := f.Description()
// 			labels := make([]string, 0, 3)
// 			if f.IsGlobal() {
// 				labels = append(labels, loc.FlagGlobalLabel)
// 			}
// 			if f.IsRequired() {
// 				labels = append(labels, loc.FlagRequiredLabel)
// 			}
// 			if f.HasDefault() {
// 				labels = append(labels, fmt.Sprintf(loc.FlagDefaultLabel, f.RawDefault()))
// 			}
// 			label := ""
// 			if len(labels) > 0 {
// 				label = fmt.Sprintf("(%s) ", strings.Join(labels, ", "))
// 			}
// 			writer.WriteLine("  %s\t%s%s", opts, label, desc)
// 		}
// 	}

// 	if hasVisiblePositionals {
// 		writer.WriteLine("")
// 		writer.WriteLine("%s:", loc.ArgumentsLabel)
// 		for _, p := range cmd.positionals {
// 			if p.IsHidden() {
// 				continue
// 			}

// 			desc := p.Description()
// 			labels := make([]string, 0, 3)
// 			if p.IsRequired() {
// 				labels = append(labels, loc.FlagRequiredLabel)
// 			}
// 			if p.HasDefault() {
// 				labels = append(labels, fmt.Sprintf(loc.FlagDefaultLabel, p.RawDefault()))
// 			}
// 			label := ""
// 			if len(labels) > 0 {
// 				label = fmt.Sprintf("(%s) ", strings.Join(labels, ", "))
// 			}
// 			writer.WriteLine("  %s\t%s%s", p.Name(), label, desc)
// 		}
// 	}

// 	return writer.Flush()
// }

// // DefaultErrorFormatter is the built-in error style. It prefixes the error
// // message with the localized error label, e.g. "Error: unknown flag x".
// func DefaultErrorFormatter(err error, loc Locale) string {
// 	return fmt.Sprintf("%s: %s", loc.ErrorLabel, loc.LocalizeError(err))
// }
