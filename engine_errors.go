package engine

import (
	"fmt"
)

type ScriptError struct {
	Err   error
	Stack string
}

func (s ScriptError) Unwrap() error {
	if s.Stack == "" {
		return s.Err
	}
	return fmt.Errorf("script error: %s\nLocate: %s", s.Err.Error(), s.Stack)
}
func (s ScriptError) Error() string {
	if s.Err == nil {
		return "<ni>"
	}
	if s.Stack == "" {
		return s.Err.Error()
	}
	return fmt.Sprintf("script error: %s\nLocate: %s", s.Err.Error(), s.Stack)
}
