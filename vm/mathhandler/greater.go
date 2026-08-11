package mathhandler

import (
	"drylang/core"
)

func OpGreater(vm core.VMCore, line, col int) error {
	b, a := vm.Pop(), vm.Pop()
	if a.Type != core.ValNumber || b.Type != core.ValNumber {
		return vm.Errorf("E300 at %d:%d: want number (a=%v, b=%v)", line, col, a.Type, b.Type)
	}
	vm.Push(core.BoolVal(a.Data.(float64) > b.Data.(float64)))
	return nil
}
