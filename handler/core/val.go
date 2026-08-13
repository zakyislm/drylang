package core

import (
	"drylang/core"
)

func BuiltinVal(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 || args[0].Type != core.ValMap {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want map", line, col)
	}
	m := args[0].Data.(map[string]core.Value)
	vals := make([]core.Value, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	result = core.ArrayVal(vals)
	return result, nil
}
