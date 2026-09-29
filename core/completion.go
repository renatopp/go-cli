package core

import (
	"fmt"
	"strings"

	"github.com/renatopp/go-cli/parsers"
)

const (
	completeCommandName      = "__complete" // hidden argument used by the shell scripts to request candidates
	completionCommandName    = "completion" // hidden command that prints the shell scripts
	completionFilesDirective = ":files"     // tells the shell script to also complete file names
)

// completionScript returns the completion script of the given shell ("bash"
// or "zsh") for the program with the given name. The script calls the program
// back on every TAB to get the candidates.
func completionScript(shell, name string) (string, error) {
	fn := "_" + strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			return r
		}
		return '_'
	}, name)

	var script string
	switch shell {
	case "bash":
		script = bashCompletionScript
	case "zsh":
		script = zshCompletionScript
	default:
		return "", fmt.Errorf("unsupported shell %q, use bash or zsh", shell)
	}

	return strings.NewReplacer("{{name}}", name, "{{fn}}", fn).Replace(script), nil
}

// newCompletionCommand creates the hidden `completion <shell>` command, which
// prints the completion script for the given shell.
func (a *App) newCompletionCommand() *Command {
	shell := NewPositional("shell", "The shell to generate the script for (bash or zsh).", parsers.String).
		AsRequired().
		WithCompletion(func(string) []string { return []string{"bash", "zsh"} }).
		WithValidation(func(s string) error {
			_, err := completionScript(s, "")
			return err
		})

	cmd := NewCommand(a.rootCommand).
		WithName(completionCommandName).
		WithShortDescription("Print the shell completion script.").
		WithDescription("Print the shell completion script for bash or zsh.").
		WithExample("source <("+a.rootCommand.name+" completion bash)", "Enable completion in the current bash session.").
		WithPositional(shell).
		AsHidden()

	return cmd.WithExecute(func() {
		a.Parse()
		script, _ := completionScript(shell.Value(), a.rootCommand.name)
		fmt.Fprint(a.stdout, script)
	})
}

// complete prints the completion candidates for the partial word (the last
// argument in the queue) of the current command and exits. Each candidate is
// printed in its own line, optionally followed by a tab and a description.
func (a *App) complete() {
	words, current := a.queue[:len(a.queue)-1], a.queue[len(a.queue)-1]
	cmd := a.currentCommand

	flags := map[string]AnyFlag{}
	for _, f := range cmd.flags {
		if f.Long() != "" {
			flags[f.Long()] = f
		}
		if f.Short() != "" {
			flags[f.Short()] = f
		}
	}

	// Leniently scan the previous words to find the used flags, the number of
	// positionals and if the last flag is waiting for a value.
	used := map[AnyFlag]bool{}
	var pending AnyFlag
	posCount := 0
	eoo := false
	for _, w := range words {
		switch {
		case pending != nil:
			pending = nil

		case !eoo && w == "--":
			eoo = true

		case !eoo && strings.HasPrefix(w, "--"):
			name, _, hasValue := strings.Cut(w[2:], "=")
			if f, ok := flags[name]; ok {
				used[f] = true
				if !hasValue && !isBoolFlag(f) {
					pending = f
				}
			}

		case !eoo && len(w) > 1 && w[0] == '-' && !(w[1] >= '0' && w[1] <= '9' || w[1] == '_'):
			for i := 1; i < len(w); i++ {
				f, ok := flags[w[i:i+1]]
				if !ok {
					break
				}
				used[f] = true
				if !isBoolFlag(f) {
					if i == len(w)-1 {
						pending = f
					}
					break
				}
			}

		default:
			posCount++
		}
	}

	var candidates []string
	files := false
	switch {
	case pending != nil:
		values, known := pending.complete(current)
		candidates, files = values, !known

	case !eoo && strings.HasPrefix(current, "--") && strings.Contains(current, "="):
		name, value, _ := strings.Cut(current[2:], "=")
		if f, ok := flags[name]; ok {
			values, known := f.complete(value)
			for _, v := range values {
				candidates = append(candidates, "--"+name+"="+v)
			}
			files = !known
		}

	case !eoo && strings.HasPrefix(current, "-"):
		for _, f := range cmd.flags {
			if f.IsHidden() || used[f] && !f.IsRepeatable() && !a.repeatedFlagsAllowed {
				continue
			}
			if f.Long() != "" {
				candidates = append(candidates, completionCandidate("--"+f.Long(), f.Description()))
			}
			if f.Short() != "" && !strings.HasPrefix(current, "--") {
				candidates = append(candidates, completionCandidate("-"+f.Short(), f.Description()))
			}
		}

	default:
		if posCount == 0 && !eoo {
			for _, sub := range cmd.subcommands {
				if !sub.hidden {
					candidates = append(candidates, completionCandidate(sub.name, sub.shortDescription))
				}
			}
		}

		var positional AnyPositional
		if posCount < len(cmd.positionals) {
			positional = cmd.positionals[posCount]
		} else if n := len(cmd.positionals); n > 0 && cmd.positionals[n-1].IsVariadic() {
			positional = cmd.positionals[n-1]
		}

		if positional != nil {
			values, known := positional.complete(current)
			candidates = append(candidates, values...)
			files = !known
		} else {
			files = a.extraPositionalsAllowed
		}
	}

	for _, c := range candidates {
		if strings.HasPrefix(c, current) {
			fmt.Fprintln(a.stdout, c)
		}
	}
	if files {
		fmt.Fprintln(a.stdout, completionFilesDirective)
	}
	a.Exit(0)
}

// completionCandidate joins a value and the first line of its description.
func completionCandidate(value, description string) string {
	description, _, _ = strings.Cut(description, "\n")
	description = strings.TrimSpace(strings.ReplaceAll(description, "\t", " "))
	if description == "" {
		return value
	}
	return value + "\t" + description
}

func isBoolFlag(f AnyFlag) bool {
	_, ok := f.(*Flag[bool])
	return ok
}

const bashCompletionScript = `# bash completion for {{name}}
{{fn}}() {
    local line="${COMP_LINE:0:COMP_POINT}" words cur out c strip=""
    COMPREPLY=()
    read -ra words <<< "$line"
    [[ $line == *[[:space:]] || ${#words[@]} -eq 0 ]] && words+=("")
    cur="${words[${#words[@]}-1]}"
    out=$("${words[0]/#\~/$HOME}" __complete "${words[@]:1}" 2>/dev/null) || return

    # bash splits the current word on "=", so candidates must not repeat the
    # "--flag=" part
    [[ $cur == *=* && $COMP_WORDBREAKS == *=* ]] && strip="${cur%=*}="

    while IFS= read -r c; do
        [[ -z $c ]] && continue
        if [[ $c == :files ]]; then
            compopt -o default 2>/dev/null
            continue
        fi
        c="${c%%$'\t'*}"
        COMPREPLY+=("${c#"$strip"}")
    done <<< "$out"
}
complete -F {{fn}} {{name}}
`

const zshCompletionScript = `#compdef {{name}}
# zsh completion for {{name}}
{{fn}}() {
    local -a out candidates
    local c value desc files=0
    out=("${(@f)$(${~words[1]} __complete "${(@Q)words[2,CURRENT]}" 2>/dev/null)}")

    for c in "${out[@]}"; do
        [[ -z $c ]] && continue
        if [[ $c == :files ]]; then
            files=1
            continue
        fi
        value=${c%%$'\t'*}
        desc=""
        [[ $c == *$'\t'* ]] && desc=${c#*$'\t'}
        candidates+=("${value//:/\\:}${desc:+:$desc}")
    done

    (( ${#candidates} )) && _describe -t values 'values' candidates
    if (( files )); then
        [[ $PREFIX == --*=* ]] && compset -P '*='
        _files
    fi
}

if [[ $funcstack[1] == {{fn}} ]]; then
    {{fn}} "$@"
else
    compdef {{fn}} {{name}}
fi
`
