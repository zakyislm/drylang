package colshandler

import (
	"drylang/core"
)

func BuiltinRm(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
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
