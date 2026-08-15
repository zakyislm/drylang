package core

import (
	"drylang/core"
)

func BuiltinGet(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want 1 arg", line, col)
	}
	result = core.StringVal(args[0].Type)
	return result, nil
}
