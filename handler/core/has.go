package core

import (
	"drylang/core"
	"strings"
)

func BuiltinHas(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 2 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want 2 args", line, col)
	}
	switch args[0].Type {
	case core.ValString:
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want string arg", line, col)
		}
		result = core.BoolVal(strings.Contains(args[0].Data.(string), args[1].Data.(string)))
	case core.ValMap:
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want string key", line, col)
		}
		m := args[0].Data.(map[string]core.Value)
		_, ok := m[args[1].Data.(string)]
		result = core.BoolVal(ok)
	case core.ValArray:
		arr := args[0].Data.([]core.Value)
		found := false
		for _, v := range arr {
			if v.String() == args[1].String() {
				found = true
				break
			}
		}
		result = core.BoolVal(found)
	default:
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want str|map|arr", line, col)
	}
	return result, nil
}
