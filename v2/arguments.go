package v2

import (
	"strings"
)

type Arguments struct {
	Args      []string // the list of registered and non-registered positional arguments
	ExtraArgs []string // the list of extra positional arguments that are not defined in the command
}

func parseArguments(app *App) (*Arguments, error) {
	parser := &parser{
		queue: app.queue,
	}
	err := parser.parse()
	if err != nil {
		return nil, err
	}
	return &Arguments{
		Args:      parser.args,
		ExtraArgs: parser.extraArgs,
	}, nil
}

type parser struct {
	queue       []string
	args        []string // the "raw" list of positional arguments, may contains non processed positionals
	extraArgs   []string // the list of extra positional arguments that are not defined in the command
	flags       map[string]Flag
	positionals []Positional
}

func (p *parser) parse() (err error) {
	eoo := false // end of options, i.e., after -- is encountered
	for {
		tok, ok := p.next()
		if !ok {
			break
		}

		switch {
		case !eoo && tok == "--":
			eoo = true
			err = nil

		case !eoo && p.isFlagToken(tok) && strings.HasPrefix(tok, "--"):
			err = p.parseLong(tok)

		case !eoo && p.isFlagToken(tok) && strings.HasPrefix(tok, "-"):
			err = p.parseShort(tok)

		default:
			err = p.parsePositional(tok)
		}

		if err != nil {
			return err
		}
	}
	return nil
}

// next returns the next token from the queue if any. It removes the token
// from the queue and returns it. Token if the next argument as passed from
// os.Args.
func (p *parser) next() (string, bool) {
	if len(p.queue) == 0 {
		return "", false
	}
	next := p.queue[0]
	p.queue = p.queue[1:]
	return next, true
}

// isFlagToken checks if the token follows the flag pattern: --flag or -f
func (p *parser) isFlagToken(token string) bool {
	if len(token) < 2 {
		return false
	}
	if strings.HasPrefix(token, "--") {
		return true
	}
	if token[0] == '-' {
		if token[1] >= '0' && token[1] <= '9' || token[1] == '_' {
			return false
		}
	}
	return false
}

// isBooleanFlag checks if the flag object is a boolean flag
func (p *parser) isBooleanFlag(name string) bool {
	if flag, ok := p.flags[name]; ok {
		_, ok := flag.(*GenericFlag[bool])
		return ok
	}
	return false
}

// tryGetFlag tries to get the flag from the command
func (p *parser) tryGetFlag(name string) (Flag, error) {
	if flag, ok := p.flags[name]; ok {
		return flag, nil
	}

	// TODO: this should be an error, but for now we just ignore unknown flags

	return nil, nil
}

// parseFlag parses the flag with the given value. It checks for repeated flags
func (p *parser) parseFlag(name string, value string) error {
	flag, err := p.tryGetFlag(name)
	if flag == nil {
		return err
	}
	// TODO: check for repeated flags if p.allowRepeatedFlags is false
	return flag.Parse(value)
}

// --name=value | --name value | --name
func (p *parser) parseLong(token string) error {
	switch {
	case strings.Contains(token, "="):
		// signed long flag, e.g., --name=value
		index := strings.Index(token, "=")
		name := token[2:index]
		value := token[index+1:]
		return p.parseFlag(name, value)

	default:
		// unsigned long flag, e.g., --name value or --name (for boolean flags)
		name := token[2:]
		if p.isBooleanFlag(name) {
			return p.parseFlag(name, "true")
		}
		value, ok := p.next()
		if !ok {
			// TODO: this should be an error
			return p.parseFlag(name, "")
		}
		return p.parseFlag(name, value)
	}
}

// -f value | -fvalue | -f | -abc
func (p *parser) parseShort(token string) error {
	name := token[1:]
	for {
		size := len(name)
		boolean := p.isBooleanFlag(name[:1])

		switch {
		case size <= 1:
			// -f or -f value
			if boolean {
				return p.parseFlag(name, "true")
			}

			value, ok := p.next()
			if !ok {
				// TODO: this should be an error
				return p.parseFlag(name, "")
			}
			return p.parseFlag(name, value)

		case boolean:
			// -abc (for boolean flags) or -abxvalue (for boolean flags followed by a non-boolean flag)
			if err := p.parseFlag(name[:1], "true"); err != nil {
				return err
			}
			name = name[1:]

		default:
			// -fvalue (for non-boolean flags)
			return p.parseFlag(name[:1], name[1:])
		}
	}
}

// value
func (p *parser) parsePositional(token string) error {
	i := len(p.args)
	var positional Positional
	if i < len(p.positionals) {
		positional = p.positionals[i]
	}

	p.args = append(p.args, token)
	if positional != nil {
		return positional.Parse(token)
	} else {
		p.extraArgs = append(p.extraArgs, token)
	}
	return nil
}
