package core

import (
	"drylang/core"
)

func BuiltinPop(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 || args[0].Type != core.ValArray {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want array", line, col)
	}
	arr := args[0].Data.([]core.Value)
	if len(arr) == 0 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"empty array", line, col)
	}
	result = arr[len(arr)-1]
	return result, nil
}
