package iohandler

import (
	"drylang/core"
	"time"
)

func BuiltinQ(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 || args[0].Type != core.ValNumber {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want number", line, col)
	}
	ms := int(args[0].Data.(float64))
	time.Sleep(time.Duration(ms) * time.Millisecond)
	result = core.UnknownValue
	return result, nil
}
