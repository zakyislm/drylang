package errfmt

import (
	"fmt"
	"strings"
	"sync"
)

// DryError is the unified runtime/compile error carrying source position.
type DryError struct {
	Code    string
	Line    int
	Col     int
	Message string
	Detail  string // optional extra context, e.g. source line + pointer
}

func (e *DryError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("%d:%d [%s] %s\n%s", e.Line, e.Col, e.Code, e.Message, e.Detail)
	}
	return fmt.Sprintf("%d:%d [%s] %s", e.Line, e.Col, e.Code, e.Message)
}

var (
	mu          sync.RWMutex
	sourceLines []string
)

// Init sets the source lines for error formatting. Thread-safe.
func Init(source string) {
	mu.Lock()
	sourceLines = strings.Split(source, "\n")
	mu.Unlock()
}

// Format formats an error with the code, line, column, and a visual pointer.
func Format(code string, line, col int, message string) error {
	mu.RLock()
	defer mu.RUnlock()

	e := &DryError{Code: code, Line: line, Col: col, Message: message}
	if line < 1 || line > len(sourceLines) {
		return e
	}

	srcLine := sourceLines[line-1]

	// Create the pointer string
	pointer := ""
	for i := 0; i < col-1 && i < len(srcLine); i++ {
		if srcLine[i] == '\t' {
			pointer += "\t"
		} else {
			pointer += " "
		}
	}
	pointer += "^"

	e.Detail = "  " + srcLine + "\n  " + pointer
	return e
}
