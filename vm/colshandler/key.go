package colshandler

import (
	"drylang/core"
)

func BuiltinKey(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 || args[0].Type != core.ValMap {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want map", line, col)
	}
	m := args[0].Data.(map[string]core.Value)
	keys := make([]core.Value, 0, len(m))
	for k := range m {
		keys = append(keys, core.StringVal(k))
	}
	result = core.ArrayVal(keys)
	return result, nil
}
