package v2

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

type Typewriter struct {
	buf    *strings.Builder
	writer *tabwriter.Writer
	// buffer := &strings.Builder{}
	// writer := tabwriter.NewWriter(buffer, 0, 0, 1, ' ', 0)
}

func NewDefaultTypewriter() *Typewriter {
	return NewTypewriter(20, 2, 2, ' ', 0)
}

func NewTypewriter(minwidth, tabwidth, padding int, padchar byte, flags uint) *Typewriter {
	buf := &strings.Builder{}
	writer := tabwriter.NewWriter(buf, minwidth, tabwidth, padding, padchar, flags)
	return &Typewriter{
		buf:    buf,
		writer: writer,
	}
}

func (t *Typewriter) Write(s string, args ...any) {
	t.writer.Write([]byte(fmt.Sprintf(s, args...)))
}

func (t *Typewriter) WriteLine(s string, args ...any) {
	t.Write(s+"\n", args...)
}

func (t *Typewriter) Flush() string {
	t.writer.Flush()
	s := t.buf.String()
	t.buf.Reset()
	return s
}
