package internal

import (
	"errors"
	"slices"
	"strings"
)

type Arguments struct {
	Raw         []string
	Args        []string
	Positionals map[int]Positional
	Flags       map[string]Flag
	HasHelp     bool

	args   []string
	strict bool
}

func parseArguments(raw []string, args []Positional, flags []Flag, strict bool) (*Arguments, error) {
	arguments := &Arguments{
		Raw:         raw,
		Args:        []string{},
		Positionals: map[int]Positional{},
		Flags:       map[string]Flag{},
		args:        slices.Clone(raw),
		strict:      strict,
	}
	for _, arg := range args {
		arguments.Positionals[arg.Index()] = arg
	}
	for _, flag := range flags {
		if flag.Long() != "" {
			arguments.Flags[flag.Long()] = flag
		}
		if flag.Short() != "" {
			arguments.Flags[flag.Short()] = flag
		}
	}

	if err := arguments.parse(); err != nil {
		return nil, err
	}
	for _, positional := range arguments.Positionals {
		positional.setParsed()
		if positional.IsRequired() && !positional.IsSet() {
			return nil, errors.New("missing required positional argument: " + positional.Name())
		}
	}
	for _, flag := range arguments.Flags {
		flag.setParsed()
		if flag.IsRequired() && !flag.IsSet() {
			if flag.Long() != "" {
				return arguments, errors.New("missing required flag: --" + flag.Long())
			}
			if flag.Short() != "" {
				return arguments, errors.New("missing required flag: -" + flag.Short())
			}
			return arguments, errors.New("missing required flag")
		}
	}
	return arguments, nil
}

func (a *Arguments) next() (value string, ok bool) {
	if len(a.args) == 0 {
		return "", false
	}
	next := a.args[0]
	a.args = a.args[1:]
	return next, true
}

func (a *Arguments) nextValue() (value string, ok bool) {
	if len(a.args) == 0 {
		return "", false
	}
	if strings.HasPrefix(a.args[0], "-") {
		return "", false
	}
	next := a.args[0]
	a.args = a.args[1:]
	return next, true
}

func (a *Arguments) parse() error {

	ignoreFlags := false
	for {
		arg, ok := a.next()
		if !ok {
			break
		}

		if arg == "--" {
			ignoreFlags = true
			continue
		}

		if !ignoreFlags && strings.HasPrefix(arg, "--") {
			if err := a.parseLong(arg); err != nil {
				return err
			}
		} else if !ignoreFlags && strings.HasPrefix(arg, "-") {
			if err := a.parseShort(arg); err != nil {
				return err
			}
		} else {
			if err := a.parsePositional(arg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *Arguments) getFlag(name string) (Flag, error) {
	if name == "help" || name == "h" {
		a.HasHelp = true
	}

	flag, ok := a.Flags[name]
	if !ok && a.strict {
		return nil, errors.New("unknown flag: " + name)
	}
	if ok && flag.IsSet() {
		return nil, errors.New("flag already set: " + name)
	}
	return flag, nil
}

// --name=value | --name value | --name
func (a *Arguments) parseLong(arg string) error {
	// Check for --name=value format
	if equalIndex := strings.Index(arg, "="); equalIndex != -1 {
		return a.parseCompleteLong(arg, equalIndex)
	}

	// No "=" found, just a name
	return a.parseIncompleteLong(arg)
}

// --name=value
func (a *Arguments) parseCompleteLong(arg string, equalIndex int) error {
	name := arg[2:equalIndex] // Remove "--" prefix
	value := arg[equalIndex+1:]
	flag, err := a.getFlag(name)
	if flag == nil {
		return err
	}
	err = flag.fromString(value)
	if err != nil {
		return err
	}
	return nil
}

// --name | --name value
func (a *Arguments) parseIncompleteLong(arg string) error {
	name := arg[2:] // Remove "--" prefix
	flag, err := a.getFlag(name)
	if flag == nil {
		return err
	}

	// Check if the flag expects a value
	if flag.acceptsValue() {
		next, ok := a.nextValue()
		if !ok {
			return errors.New("flag requires a value: --" + name)
		}
		err = flag.fromString(next)
		if err != nil {
			return err
		}
		return nil
	}

	return flag.fromString("true")
}

// -a | -o output.txt | -ooutput.txt | -abc
func (a *Arguments) parseShort(arg string) error {
	name := arg[1:] // Remove "-" prefix
	if len(name) > 1 {
		flag, err := a.getFlag(name[:1])
		if flag == nil {
			return err
		}

		if !flag.acceptsValue() {
			return a.parseCombinedShort(name)
		}
		return a.parseValuedShort(name)
	}
	return a.parseFlaglessShort(name)
}

// -abc (for boolean flags)
func (a *Arguments) parseCombinedShort(arg string) error {
	for _, char := range arg {
		flag, err := a.getFlag(string(char))
		if err != nil {
			return err
		}
		if flag == nil {
			continue
		}
		if flag.acceptsValue() {
			return errors.New("flag requires a value: -" + string(char))
		}
		if err := flag.fromString("true"); err != nil {
			return err
		}
	}
	return nil
}

// -ooutput.txt
func (a *Arguments) parseValuedShort(arg string) error {
	flag, err := a.getFlag(arg[:1])
	if flag == nil {
		return err
	}
	if !flag.acceptsValue() {
		return errors.New("flag does not accept a value: -" + arg[:1])
	}
	if err := flag.fromString(arg[1:]); err != nil {
		return err
	}
	return nil
}

// -v | -o output.txt
func (a *Arguments) parseFlaglessShort(arg string) error {
	flag, err := a.getFlag(arg)
	if flag == nil {
		return err
	}

	// Check if the flag expects a value
	if flag.acceptsValue() {
		next, ok := a.next()
		if !ok {
			return errors.New("flag requires a value: -" + arg)
		}
		err = flag.fromString(next)
		if err != nil {
			return err
		}
		return nil
	}

	return flag.fromString("true")
}

func (a *Arguments) parsePositional(arg string) error {
	a.Args = append(a.Args, arg)

	index := len(a.Args)
	positional, ok := a.Positionals[index-1]
	if !ok {
		return nil
	}
	if positional.IsSet() {
		return errors.New("positional argument already set: " + positional.Name())
	}

	return positional.fromString(arg)
}
