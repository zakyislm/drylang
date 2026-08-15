package core

import (
	"drylang/core"
	"strings"
)

func BuiltinJ(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 2 || args[0].Type != core.ValArray || args[1].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want array, string", line, col)
	}
	arr := args[0].Data.([]core.Value)
	strs := make([]string, len(arr))
	for i, v := range arr {
		strs[i] = v.String()
	}
	result = core.StringVal(strings.Join(strs, args[1].Data.(string)))
	return result, nil
}
