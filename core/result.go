package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/renatopp/go-cli/errors"
	"github.com/renatopp/go-cli/parsers"
)

type Result struct {
	pos      []string // the list of registered and non-registered positional arguments
	extraPos []string // the list of extra positional arguments that are not defined in the command

	app                   *App     // reference to the main app for accessing configuration like AllowExtraPositionals
	queue                 []string // the queue of arguments to be parsed
	flags                 map[string]AnyFlag
	positionals           []AnyPositional
	hasVariadicPositional bool
	hasHelpFlag           bool
	hasVersionFlag        bool
}

func (a *Result) PosCount() int {
	return len(a.pos)
}

func (a *Result) ExtraPosCount() int {
	return len(a.extraPos)
}

func (a *Result) PosAt(index int) string {
	args := a.pos
	if index < 0 || index >= len(args) {
		return ""
	}
	return args[index]
}

func (a *Result) ExtraPosAt(index int) string {
	args := a.extraPos
	if index < 0 || index >= len(args) {
		return ""
	}
	return args[index]
}

func (a *Result) Pos() []string {
	return a.pos[:]
}

func (a *Result) ExtraPos() []string {
	return a.extraPos[:]
}

// next returns the next token from the queue if any. It removes the token
// from the queue and returns it. Token if the next argument as passed from
// os.Args.
func (a *Result) next() (string, bool) {
	if len(a.queue) == 0 {
		return "", false
	}
	next := a.queue[0]
	a.queue = a.queue[1:]
	return next, true
}

// isFlagToken checks if the token follows the flag pattern: --flag or -f
func (a *Result) isFlagToken(token string) bool {
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
func (a *Result) isBooleanFlag(name string) bool {
	if flag, ok := a.flags[name]; ok {
		_, ok := flag.(*Flag[bool])
		return ok
	}
	return false
}

// tryGetFlag tries to get the flag from the command
func (a *Result) tryGetFlag(name string) (AnyFlag, error) {
	if name == "help" || name == "h" {
		a.hasHelpFlag = true
	}
	if name == "version" || name == "v" {
		a.hasVersionFlag = true
	}

	if flag, ok := a.flags[name]; ok {
		return flag, nil
	}

	if a.app.extraFlagsAllowed {
		long := ""
		short := ""
		if len(name) == 1 {
			short = name
		} else {
			long = name
		}

		// If extra flags are allowed, we create a new generic string flag for the unknown flag and add it to the flags map. This allows users to access the value of the extra flag using the same API as regular flags.
		extraFlag := NewFlag(long, short, "", parsers.String)
		a.flags[name] = extraFlag
		return extraFlag, nil
	}
	return nil, errors.NewUnknownFlagError(name)
}

// parseFlag parses the flag with the given value. It checks for repeated flags
func (a *Result) parseFlag(name string, value string) error {
	flag, err := a.tryGetFlag(name)
	if flag == nil {
		return err
	}

	if flag.IsProvided() {
		if !a.app.repeatedFlagsAllowed && !flag.IsRepeatable() {
			return errors.NewRepeatedFlagError(name)
		}
	}

	return flag.parse(value)
}

// --name=value | --name value | --name
func (a *Result) parseLong(token string) error {
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
		if hasFlag {
			value, ok := a.next()
			if !ok {
				return errors.NewMissingFlagValueError(name)
			}
			return a.parseFlag(name, value)
		}
		return a.parseFlag(name, "")
	}
}

// -f value | -fvalue | -f | -abc
func (a *Result) parseShort(token string) error {
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
				return errors.NewMissingFlagValueError(name)
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
func (a *Result) parsePositional(token string) error {
	i := len(a.pos)
	var positional AnyPositional
	if i < len(a.positionals) {
		positional = a.positionals[i]
	}

	a.pos = append(a.pos, token)
	if positional != nil {
		return positional.parse(token)
	} else {
		if a.hasVariadicPositional {
			last := a.positionals[len(a.positionals)-1]
			return last.parse(token)
		}

		if !a.app.extraPositionalsAllowed {
			return errors.NewUnexpectedPosError(token)
		}

		a.extraPos = append(a.extraPos, token)
	}
	return nil
}

func parseArguments(app *App) (*Result, error) {
	args := &Result{
		pos:            []string{},
		extraPos:       []string{},
		app:            app,
		queue:          app.queue,
		flags:          map[string]AnyFlag{},
		positionals:    []AnyPositional{},
		hasHelpFlag:    false,
		hasVersionFlag: false,
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

	// check for auto help or auto version before performing other validations
	if app.autoHelp && args.hasHelpFlag {
		app.Help()
		app.Exit(0)
	}

	if app.version != "" && args.hasVersionFlag && app.CurrentCommand() == app.RootCommand() {
		fmt.Fprintf(app.stdout, "%s\n", app.version)
		app.Exit(0)
	}

	// check validations
	if err != nil {
		return args, err
	}

	// check for required flags and positionals
	for _, flag := range cmd.flags {
		if !flag.IsProvided() && flag.HasEnv() {
			if value, ok := os.LookupEnv(flag.Env()); ok {
				if err := flag.parse(value); err != nil {
					return args, err
				}
			}
		}

		if flag.IsRequired() && !flag.IsProvided() {
			return args, errors.NewMissingRequiredFlagError(flag.Signature())
		}
	}
	for _, positional := range cmd.positionals {
		if !positional.IsProvided() && positional.HasEnv() {
			if value, ok := os.LookupEnv(positional.Env()); ok {
				if err := positional.parse(value); err != nil {
					return args, err
				}
			}
		}

		if positional.IsRequired() && !positional.IsProvided() {
			return args, errors.NewMissingRequiredPosError(positional.Name())
		}
	}

	return args, nil
}
