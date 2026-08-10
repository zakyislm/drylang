package colshandler

import (
	"drylang/core"
)

func BuiltinLen(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want 1 arg", line, col)
	}
	switch args[0].Type {
	case core.ValString:
		result = core.NumberVal(float64(len(args[0].Data.(string))))
	case core.ValArray:
		result = core.NumberVal(float64(len(args[0].Data.([]core.Value))))
	case core.ValMap:
		result = core.NumberVal(float64(len(args[0].Data.(map[string]core.Value))))
	default:
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want string|array|map", line, col)
	}
	return result, nil
}
