package oshandler

import (
	"drylang/core"
	"os"
)

func BuiltinDel(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want path string", line, col)
	}
	err := os.Remove(args[0].Data.(string))
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"del fail: %v", err, line, col)
	}
	result = core.BoolVal(true)
	return result, nil
}
