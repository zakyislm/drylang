package mathhandler

import (
	"drylang/core"
	"math"
)

func BuiltinMax(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 2 || args[0].Type != core.ValNumber || args[1].Type != core.ValNumber {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want 2 numbers", line, col)
	}
	result = core.NumberVal(math.Max(args[0].Data.(float64), args[1].Data.(float64)))
	return result, nil
}
