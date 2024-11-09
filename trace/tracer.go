package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
	GetWriter() io.Writer
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func (t *tracer) GetWriter() io.Writer {
	return t.out
}
