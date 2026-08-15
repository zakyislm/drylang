package core

import (
	"drylang/core"
	"fmt"
)

// BuiltinFmt handles fmt builtin, similar to sprintf
func BuiltinFmt(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) == 0 {
		return core.StringVal(""), nil
	}
	if args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("fmt first arg must be format string")
	}
	format := args[0].Data.(string)

	goArgs := make([]interface{}, len(args)-1)
	for i := 1; i < len(args); i++ {
		val := args[i].Data
		// Intelligently convert whole numbers to int so %d works
		if f, ok := val.(float64); ok && f == float64(int(f)) {
			val = int(f)
		}
		goArgs[i-1] = val
	}

	result := fmt.Sprintf(format, goArgs...)
	return core.StringVal(result), nil
}
