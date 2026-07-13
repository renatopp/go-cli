package formatters

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

type typewriter struct {
	buf    *strings.Builder
	writer *tabwriter.Writer
	// buffer := &strings.Builder{}
	// writer := tabwriter.NewWriter(buffer, 0, 0, 1, ' ', 0)
}

func newDefaultTypewriter() *typewriter {
	return newTypewriter(20, 2, 2, ' ', 0)
}

func newTypewriter(minwidth, tabwidth, padding int, padchar byte, flags uint) *typewriter {
	buf := &strings.Builder{}
	writer := tabwriter.NewWriter(buf, minwidth, tabwidth, padding, padchar, flags)
	return &typewriter{
		buf:    buf,
		writer: writer,
	}
}

func (t *typewriter) Write(s string, args ...any) {
	fmt.Fprintf(t.writer, s, args...)
}

func (t *typewriter) WriteLine(s string, args ...any) {
	t.Write(s+"\n", args...)
}

func (t *typewriter) Flush() string {
	t.writer.Flush()
	s := t.buf.String()
	t.buf.Reset()
	return s
}
