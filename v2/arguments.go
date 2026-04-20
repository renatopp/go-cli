package v2

import (
	"strings"
	"unicode/utf8"

	"golang.org/x/text/runes"
)

func parseArguments(app *App) {
	parser := &parser{
		queue: app.queue,
	}
	parser.parse()
}

type parser struct {
	queue []string
}

func (p *parser) parse() (err error) {
	eoo := false // end of options, i.e., after -- is encountered
	for {
		tok, ok := p.nextToken()
		if !ok {
			break
		}

		if tok == "--" {
			eoo = true
			continue
		}

		switch {
		case !eoo && strings.HasPrefix(tok, "--"):
			err = p.parseLong(tok)
		case !eoo && strings.HasPrefix(tok, "-"):
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

// nextToken returns the next token from the queue if any. It removes the token
// from the queue and returns it. Token if the next argument as passed from
// os.Args.
func (p *parser) nextToken() (string, bool) {
	if len(p.queue) == 0 {
		return "", false
	}
	next := p.queue[0]
	p.queue = p.queue[1:]
	return next, true
}

// nextValue returns the next token from the queue if any and if it is not
// considered a flag.
func (p *parser) nextValue() (string, bool) {
	if len(p.queue) == 0 {
		return "", false
	}
	next := p.queue[0]
	if p.isFlagToken(next) {
		return "", false
	}
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
		r, _ := utf8.DecodeRuneInString(token[1:])
		return runes.IsLetter(r) || runes.IsMark(r) || token[1] == '_'
	}
	return false
}

// isBooleanFlag checks if the flag object is a boolean flag
func (p *parser) isBooleanFlag(flag Flag) bool {
	_, ok := flag.(*GenericFlag[bool])
	return ok
}

// tryGetFlag tries to get the flag from the command
func (p *parser) tryGetFlag(name string) (Flag, error) {
	return nil, nil
}

// --name=value | --name value | --name
func (p *parser) parseLong(token string) error {
	switch {
	case strings.Contains(token, "="):
		return p.parseLongAssined(token)
	default:
		return p.parseLongUnassigned(token)
	}
}

// --name=value
func (p *parser) parseLongAssined(token string) error {
	index := strings.Index(token, "=")
	name := token[2:index]
	value := token[index+1:]
	flag, err := p.tryGetFlag(name)
	if err != nil {
		return err
	}
	return flag.Parse(value)
}

// --name value | --name
func (p *parser) parseLongUnassigned(token string) error {
	name := token[2:]
	flag, err := p.tryGetFlag(name)
	if err != nil {
		return err
	}

	if p.isBooleanFlag(flag) {
		return flag.Parse("true")
	}

	value, ok := p.nextValue()
	if !ok {
		// TODO: this should be an error
		return flag.Parse("")
	}
	return flag.Parse(value)
}

// -f value | -fvalue | -f
func (p *parser) parseShort(token string) error {
	// name := token[1:]
	// size := utf8.RuneCountInString(name)
	// firstFlag, err := p.tryGetFlag(name[:1])
	// switch {
	// 	case size <= 1:
	// 		return p.parseShortUncombined(token)
	// 	case strings.Contains(name, "="):
	// 		return p.parseShortCombinedWithValue(token)
	// 	case err == nil && p.isBooleanFlag(firstFlag):
	// 		return p.parseShortUncombined(token)
	// 	default:
	// 		return p.parseShortCombined(token)
	// }
}

// -abc (for boolean flags)
func (p *parser) parseShortCombined(token string) {}

// -foutput
func (p *parser) parseShortCombinedWithValue(token string) {}

// -v | -o output.txt
func (p *parser) parseShortUncombined(token string) {}

// value
func (p *parser) parsePositional(token string) error { return nil }
