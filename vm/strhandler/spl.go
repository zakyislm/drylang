package strhandler

import (
	"drylang/core"
	"strings"
)

func BuiltinSpl(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 2 || args[0].Type != core.ValString || args[1].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want 2 strings", line, col)
	}
	parts := strings.Split(args[0].Data.(string), args[1].Data.(string))
	arr := make([]core.Value, len(parts))
	for i, p := range parts {
		arr[i] = core.StringVal(p)
	}
	result = core.ArrayVal(arr)
	return result, nil
}
