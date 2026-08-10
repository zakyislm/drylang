package mathhandler

import (
	"drylang/core"
)

func OpAdd(vm core.VMCore, line, col int) error {
	b, a := vm.Pop(), vm.Pop()
	if a.Type == core.ValString || b.Type == core.ValString {
		vm.Push(core.StringVal(a.String() + b.String()))
	} else if a.Type == core.ValNumber && b.Type == core.ValNumber {
		vm.Push(core.NumberVal(a.Data.(float64) + b.Data.(float64)))
	} else {
		return vm.Errorf("E300 at %d:%d: want number", line, col)
	}
	return nil
}
