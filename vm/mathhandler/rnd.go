package mathhandler

import (
	"drylang/core"
	"math"
)

func BuiltinRnd(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 || args[0].Type != core.ValNumber {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want number", line, col)
	}
	result = core.NumberVal(math.Round(args[0].Data.(float64)))
	return result, nil
}
