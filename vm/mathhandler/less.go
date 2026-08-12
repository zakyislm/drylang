package mathhandler

import (
	"drylang/core"
)

func OpLess(vm core.VMCore, line, col int) error {
	b, a := vm.Pop(), vm.Pop()
	if a.Type != core.ValNumber || b.Type != core.ValNumber {
		return vm.Errorf("E300 at %d:%d: want number, got a=%v (type %s), b=%v (type %s)", line, col, a.Data, a.Type, b.Data, b.Type)
	}
	vm.Push(core.BoolVal(a.Data.(float64) < b.Data.(float64)))
	return nil
}
