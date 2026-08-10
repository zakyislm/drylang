package strhandler

import (
	"drylang/core"
	"strings"
)

func BuiltinMod(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 3 || args[0].Type != core.ValString || args[1].Type != core.ValString || args[2].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want 3 strings", line, col)
	}
	result = core.StringVal(strings.ReplaceAll(args[0].Data.(string), args[1].Data.(string), args[2].Data.(string)))
	return result, nil
}
