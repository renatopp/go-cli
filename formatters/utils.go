package formatters

import (
	"os"
	"strings"
)

var colorMode = "none"

const bold = "\033[1m"
const dim = "\033[2m"
const italic = "\033[3m"
const underline = "\033[4m"
const reset = "\033[0m"
const accent = "\033[38;5;171m"
const red = "\033[38;5;9m"

func init() {
	if supportsColor() {
		colorMode = "color"
	}
	println("mode:", colorMode)
}

func width() int {
	return 100
}

func titleStyle(s string) string {
	s = strings.ToUpper(s)

	switch colorMode {
	case "color":
		return bold + accent + s + reset
	default:
		return s
	}
}

func descriptionStyle(s string) string {
	s = strings.TrimSpace(s)
	// s = wrap(s, width())
	switch colorMode {
	case "color":
		return dim + s + reset
	default:
		return s
	}
}

func shortDescriptionStyle(s string) string {
	s = strings.ReplaceAll(s, "\t", "")
	// s = wrap(s, width())
	s = strings.ReplaceAll(s, "\n", "\n\t")
	switch colorMode {
	case "color":
		return dim + s + reset
	default:
		return s
	}
}

func tagStyle(s string) string {
	switch colorMode {
	case "color":
		return dim + italic + s + reset
	default:
		return s
	}
}

func errorStyle(s string) string {
	switch colorMode {
	case "color":
		return bold + red + s + reset
	default:
		return s
	}
}

func argStyle(s string) string {
	switch colorMode {
	case "color":
		return s
	default:
		return s
	}
}

func indent(s string, spaces int, char string) string {
	indent := strings.Repeat(char, spaces)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = indent + line
	}
	return strings.Join(lines, "\n")
}

func wrap(s string, width int) string {
	if width <= 0 {
		return s
	}
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}
	var lines []string
	var currentLine string
	for _, word := range words {
		if len(currentLine)+len(word)+1 > width {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return strings.Join(lines, "\n")
}

func isTTY() bool {
	info, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) != 0
}

// supportsColor returns whether ANSI colors are likely supported.
func supportsColor() bool {
	if !isTTY() {
		return false
	}

	// Explicitly disabled.
	if os.Getenv("NO_COLOR") != "" {
		return false
	}

	// Explicitly forced.
	if os.Getenv("CLICOLOR_FORCE") != "" {
		return true
	}

	term := os.Getenv("TERM")
	if term == "" || term == "dumb" {
		return false
	}

	return true
}
