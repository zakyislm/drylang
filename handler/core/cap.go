package core

import (
	"drylang/core"
	"strings"
)

func BuiltinCap(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want string", line, col)
	}
	result = core.StringVal(strings.ToUpper(args[0].Data.(string)))
	return result, nil
}
