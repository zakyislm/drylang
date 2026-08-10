package oshandler

import (
	"drylang/core"
	"os"
)

func BuiltinDir(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want path string", line, col)
	}
	entries, err := os.ReadDir(args[0].Data.(string))
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"dir fail: %v", err, line, col)
	}
	arr := make([]core.Value, len(entries))
	for i, e := range entries {
		arr[i] = core.StringVal(e.Name())
	}
	result = core.ArrayVal(arr)
	return result, nil
}
