package op

import (
	"drylang/core"
	"strings"
)

func OpGreaterEq(vm core.VMCore, line, col int) error {
	b, a := vm.Pop(), vm.Pop()
	switch {
	case a.Type == core.ValNumber && b.Type == core.ValNumber:
		vm.Push(core.BoolVal(a.Data.(float64) >= b.Data.(float64)))
	case a.Type == core.ValString && b.Type == core.ValString:
		vm.Push(core.BoolVal(strings.Compare(a.Data.(string), b.Data.(string)) >= 0))
	default:
		return vm.Errorf("E300 at %d:%d: want number or string, got %s vs %s", line, col, a.Type, b.Type)
	}
	return nil
}
