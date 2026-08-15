package vm

import (
	"drylang/errfmt"
	"drylang/handler/system"
	"fmt"
	"strings"
)

func (vm *VM) runtimeErr(code string, line, col int, format string, args ...interface{}) error {
	system.Errors.Add(1)
	baseErr := errfmt.Format(code, line, col, fmt.Sprintf(format, args...))
	if len(vm.frames) <= 1 {
		return baseErr
	}

	var trace strings.Builder
	trace.WriteString(fmt.Sprintf("%v\nStack trace:", baseErr))

	// Start from the current frame (top of the stack) down to the bottom
	for i := len(vm.frames) - 1; i >= 0; i-- {
		frame := vm.frames[i]
		fnName := "main"
		if frame.closure != nil && frame.closure.Fn != nil && frame.closure.Fn.Name != "" {
			fnName = frame.closure.Fn.Name
		}

		// For the top frame, we already have exact line/col
		if i == len(vm.frames)-1 {
			trace.WriteString(fmt.Sprintf("\n  at %s() (line %d)", fnName, line))
		} else {
			// For previous frames, use frame.ip - 1 to get the call instruction's line
			ip := frame.ip - 1
			if ip < 0 {
				ip = 0
			}
			fLine := 0
			if ip < len(frame.chunk.Lines) {
				fLine = frame.chunk.Lines[ip]
			}
			trace.WriteString(fmt.Sprintf("\n  at %s() (line %d)", fnName, fLine))
		}
	}
	return fmt.Errorf("%s", trace.String())
}
