package v2

import (
	"fmt"
	"strings"
)

type Arguments struct {
	Args      []string // the list of registered and non-registered positional arguments
	ExtraArgs []string // the list of extra positional arguments that are not defined in the command

	app                   *App     // reference to the main app for accessing configuration like AllowExtraPositionals
	queue                 []string // the queue of arguments to be parsed
	flags                 map[string]Flag
	positionals           []Positional
	hasVariadicPositional bool
}

func parseArguments(app *App) (*Arguments, error) {
	args := &Arguments{
		Args:        []string{},
		ExtraArgs:   []string{},
		app:         app,
		queue:       app.queue,
		flags:       map[string]Flag{},
		positionals: []Positional{},
	}

	cmd := app.CurrentCommand()

	// prepare the flags and positionals maps for easy lookup during parsing
	for _, flag := range cmd.flags {
		args.flags[flag.Long()] = flag
		args.flags[flag.Short()] = flag
	}
	for _, positional := range cmd.positionals {
		args.positionals = append(args.positionals, positional)
		if positional.IsVariadic() {
			args.hasVariadicPositional = true
		}
	}

	// parse the arguments
	eoo := false // end of options, i.e., after -- is encountered
	var err error
	for {
		tok, ok := args.next()
		if !ok {
			break
		}

		switch {
		case !eoo && tok == "--":
			eoo = true
			err = nil

		case !eoo && args.isFlagToken(tok) && strings.HasPrefix(tok, "--"):
			err = args.parseLong(tok)

		case !eoo && args.isFlagToken(tok) && strings.HasPrefix(tok, "-"):
			err = args.parseShort(tok)

		default:
			err = args.parsePositional(tok)
		}

		if err != nil {
			break
		}
	}

	if err != nil {
		return args, err
	}

	// check for required flags and positionals
	for _, flag := range cmd.flags {
		if flag.IsRequired() && !flag.IsParsed() {
			return args, fmt.Errorf("missing required flag %s", flag.Signature())
		}
	}
	for i, positional := range cmd.positionals {
		if positional.IsRequired() && i >= len(args.Args) {
			return args, fmt.Errorf("missing required positional argument: %s", positional.Name())
		}
	}

	return args, nil
}

// next returns the next token from the queue if any. It removes the token
// from the queue and returns it. Token if the next argument as passed from
// os.Args.
func (a *Arguments) next() (string, bool) {
	if len(a.queue) == 0 {
		return "", false
	}
	next := a.queue[0]
	a.queue = a.queue[1:]
	return next, true
}

// isFlagToken checks if the token follows the flag pattern: --flag or -f
func (a *Arguments) isFlagToken(token string) bool {
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
	return true
}

// isBooleanFlag checks if the flag object is a boolean flag
func (a *Arguments) isBooleanFlag(name string) bool {
	if flag, ok := a.flags[name]; ok {
		_, ok := flag.(*GenericFlag[bool])
		return ok
	}
	return false
}

// tryGetFlag tries to get the flag from the command
func (a *Arguments) tryGetFlag(name string) (Flag, error) {
	if flag, ok := a.flags[name]; ok {
		return flag, nil
	}

	if a.app.ExtraFlagsAllowed {
		long := ""
		short := ""
		if len(name) == 1 {
			short = name
		} else {
			long = name
		}

		// If extra flags are allowed, we create a new generic string flag for the unknown flag and add it to the flags map. This allows users to access the value of the extra flag using the same API as regular flags.
		extraFlag := NewGenericFlag(long, short, "", ParseString)
		a.flags[name] = extraFlag
		return extraFlag, nil
	}
	return nil, fmt.Errorf("unknown flag %s", name)
}

// parseFlag parses the flag with the given value. It checks for repeated flags
func (a *Arguments) parseFlag(name string, value string) error {
	flag, err := a.tryGetFlag(name)
	if flag == nil {
		return err
	}

	if flag.IsParsed() {
		if !a.app.RepeatedFlagsAllowed && !flag.IsRepeatable() {
			return fmt.Errorf("flag %s was specified multiple times", name)
		}
	}

	return flag.Parse(value)
}

// --name=value | --name value | --name
func (a *Arguments) parseLong(token string) error {
	switch {
	case strings.Contains(token, "="):
		// signed long flag, e.g., --name=value
		index := strings.Index(token, "=")
		name := token[2:index]
		value := token[index+1:]
		return a.parseFlag(name, value)

	default:
		// unsigned long flag, e.g., --name value or --name (for boolean flags)
		name := token[2:]
		if a.isBooleanFlag(name) {
			return a.parseFlag(name, "true")
		}
		_, hasFlag := a.flags[name]
		value, ok := a.next()
		if hasFlag && !ok {
			return fmt.Errorf("missing value for flag %s", name)
		}
		return a.parseFlag(name, value)
	}
}

// -f value | -fvalue | -f | -abc
func (a *Arguments) parseShort(token string) error {
	name := token[1:]
	for {
		size := len(name)
		boolean := a.isBooleanFlag(name[:1])

		switch {
		case size <= 1:
			// -f or -f value
			if boolean {
				return a.parseFlag(name, "true")
			}

			_, hasFlag := a.flags[name]
			value, ok := a.next()
			if hasFlag && !ok {
				return fmt.Errorf("missing value for flag %s", name)
			}
			return a.parseFlag(name, value)

		case boolean:
			// -abc (for boolean flags) or -abxvalue (for boolean flags followed by a non-boolean flag)
			if err := a.parseFlag(name[:1], "true"); err != nil {
				return err
			}
			name = name[1:]

		default:
			// -fvalue (for non-boolean flags)
			return a.parseFlag(name[:1], name[1:])
		}
	}
}

// value
func (a *Arguments) parsePositional(token string) error {
	i := len(a.Args)
	var positional Positional
	if i < len(a.positionals) {
		positional = a.positionals[i]
	}

	a.Args = append(a.Args, token)
	if positional != nil {
		return positional.Parse(token)
	} else {
		if a.hasVariadicPositional {
			last := a.positionals[len(a.positionals)-1]
			return last.Parse(token)
		}

		if !a.app.ExtraPositionalsAllowed {
			return fmt.Errorf("unexpected extra positional argument: %s", token)
		}

		a.ExtraArgs = append(a.ExtraArgs, token)
	}
	return nil
}
