package core

import (
	"drylang/core"
)

func BuiltinAdd(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 2 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want 2 args", line, col)
	}
	if args[0].Type != core.ValArray {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want array", line, col)
	}
	arr := args[0].Data.([]core.Value)
	arr = append(arr, args[1])
	// Modify in place via globals would need ref semantics.
	// For now, push new array.
	result = core.ArrayVal(arr)
	return result, nil
}
