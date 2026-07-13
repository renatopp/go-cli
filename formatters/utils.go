package formatters

import (
	"os"
	"strings"
)

var no_color = os.Getenv("NO_COLOR")

// var columns, _ = strconv.Atoi(os.Getenv("COLUMNS"))

func width() int {
	return 100
}

func titleStyle(s string) string {
	s = strings.ToUpper(s)
	if no_color != "" {
		return s
	}
	return "\033[1m\033[38;5;5m" + s + "\033[0m"
}

func descriptionStyle(s string) string {
	s = strings.TrimSpace(s)
	// s = wrap(s, width())
	if no_color != "" {
		return s
	}
	return "\033[2m" + s + "\033[0m"
}

func shortDescriptionStyle(s string) string {
	s = strings.ReplaceAll(s, "\t", "")
	// s = wrap(s, width())
	s = strings.ReplaceAll(s, "\n", "\n\t")
	if no_color != "" {
		return s
	}
	return "\033[2m" + s + "\033[0m"
}

func tagStyle(s string) string {
	if no_color != "" {
		return s
	}
	return "\033[1m\033[3m" + s + "\033[0m"
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
