package core

import (
	"drylang/core"
	"sort"
)

func BuiltinSort(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 || args[0].Type != core.ValArray {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want array", line, col)
	}
	arr := make([]core.Value, len(args[0].Data.([]core.Value)))
	copy(arr, args[0].Data.([]core.Value))
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].Type == core.ValNumber && arr[j].Type == core.ValNumber {
			return arr[i].Data.(float64) < arr[j].Data.(float64)
		}
		return arr[i].String() < arr[j].String()
	})
	result = core.ArrayVal(arr)
	return result, nil
}
