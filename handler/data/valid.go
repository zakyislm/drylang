package data

import (
	"drylang/core"
	"net/mail"
	"strconv"
)

// BuiltinValid handles valid.mail, valid.len, valid.num
func BuiltinValid(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 2 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("want valid(method, value)")
	}
	method := args[0].Data.(string)

	switch method {
	case "mail":
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("valid.mail wants string")
		}
		address := args[1].Data.(string)
		_, err := mail.ParseAddress(address)
		return core.BoolVal(err == nil), nil

	case "len":
		if args[1].Type == core.ValString {
			return core.NumberVal(float64(len(args[1].Data.(string)))), nil
		}
		if args[1].Type == core.ValArray {
			return core.NumberVal(float64(len(args[1].Data.([]core.Value)))), nil
		}
		if args[1].Type == core.ValMap {
			return core.NumberVal(float64(len(args[1].Data.(map[string]core.Value)))), nil
		}
		return core.UnknownValue, vm.Errorf("valid.len wants string, array, or map")

	case "num":
		if args[1].Type == core.ValNumber {
			return core.BoolVal(true), nil
		}
		if args[1].Type == core.ValString {
			_, err := strconv.ParseFloat(args[1].Data.(string), 64)
			return core.BoolVal(err == nil), nil
		}
		return core.BoolVal(false), nil

	default:
		return core.UnknownValue, vm.Errorf("unknown valid method: %s", method)
	}
}
