package iohandler

import (
	"drylang/core"
	"os"
)

func BuiltinR(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want string", line, col)
	}
	data, err := os.ReadFile(args[0].Data.(string))
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"read fail", line, col)
	}
	result = core.StringVal(string(data))
	return result, nil
}
