package mathhandler

import (
	"drylang/core"
	"strconv"
)

func BuiltinNum(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want 1 arg", line, col)
	}
	f, err := strconv.ParseFloat(args[0].String(), 64)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"bad number", line, col)
	}
	result = core.NumberVal(f)
	return result, nil
}
