package trace

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	// tracer = &tracer{out: &buf}
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		// buf.Write("....")
		tracer.Trace("Hello trace package")
		fmt.Println(tracer.GetWriter())
		if buf.String() != "Hello trace package\n" {
			t.Errorf("Trace should not write %s", buf.String())
		}
	}
}
