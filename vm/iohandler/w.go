package iohandler

import (
	"drylang/core"
	"os"
)

func BuiltinW(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 2 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want path, data", line, col)
	}
	err := os.WriteFile(args[0].Data.(string), []byte(args[1].String()), 0644)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"write fail", line, col)
	}
	result = core.BoolVal(true)
	return result, nil
}
