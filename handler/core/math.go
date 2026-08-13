package core

import (
	"drylang/core"
	"math"
)

func BuiltinMath(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) < 2 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("want math(op, ...numbers)")
	}
	op := args[0].Data.(string)
	if args[1].Type != core.ValNumber {
		return core.UnknownValue, vm.Errorf("want number")
	}
	a := args[1].Data.(float64)
	switch op {
	case "sqrt":
		result = core.NumberVal(math.Sqrt(a))
	case "pow":
		if len(args) < 3 || args[2].Type != core.ValNumber {
			return core.UnknownValue, vm.Errorf("pow wants 2 numbers")
		}
		result = core.NumberVal(math.Pow(a, args[2].Data.(float64)))
	case "ceil":
		result = core.NumberVal(math.Ceil(a))
	case "floor":
		result = core.NumberVal(math.Floor(a))
	case "sin":
		result = core.NumberVal(math.Sin(a))
	case "cos":
		result = core.NumberVal(math.Cos(a))
	case "tan":
		result = core.NumberVal(math.Tan(a))
	case "log":
		result = core.NumberVal(math.Log(a))
	case "log10":
		result = core.NumberVal(math.Log10(a))
	default:
		return core.UnknownValue, vm.Errorf("unknown math op: %s", op)
	}
	return result, nil
}
