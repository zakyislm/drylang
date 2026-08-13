package core

import (
	"drylang/core"
	"os"
)

func BuiltinRm(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) == 1 && args[0].Type == core.ValString {
		err := os.Remove(args[0].Data.(string))
		if err != nil {
			return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"rm fail: %v", err, line, col)
		}
		return core.BoolVal(true), nil
	}
	if len(args) != 2 || args[0].Type != core.ValArray || args[1].Type != core.ValNumber {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want array, number", line, col)
	}
	arr := args[0].Data.([]core.Value)
	idx := int(args[1].Data.(float64))
	if idx < 0 || idx >= len(arr) {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"bounds %d", idx, line, col)
	}
	newArr := make([]core.Value, 0, len(arr)-1)
	newArr = append(newArr, arr[:idx]...)
	newArr = append(newArr, arr[idx+1:]...)
	result = core.ArrayVal(newArr)
	return result, nil
}
