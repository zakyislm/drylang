package errfmt

import (
	"fmt"
	"strings"
)

var sourceLines []string

// Init initializes the source lines for error formatting.
func Init(source string) {
	sourceLines = strings.Split(source, "\n")
}

// Format formats an error with the code, line, column, and a visual pointer.
func Format(code string, line, col int, message string) error {
	if line < 1 || line > len(sourceLines) {
		return fmt.Errorf("%d:%d [%s] %s", line, col, code, message)
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

	return fmt.Errorf("%d:%d [%s] %s\n  %s\n  %s", line, col, code, message, srcLine, pointer)
}
