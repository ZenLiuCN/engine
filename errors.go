package engine

import (
	"bytes"
	"fmt"
)

type ScriptError struct {
	Err   error
	Stack *bytes.Buffer
}

func (s ScriptError) Error() string {
	if s.Err == nil {
		return "<ni>"
	}
	if s.Stack == nil {
		return s.Err.Error()
	}
	return fmt.Sprintf("script error: %s\nscript stack: %s", s.Err.Error(), s.Stack.String())
}
