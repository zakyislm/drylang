package op

import (
	"drylang/core"
)

func OpNeg(vm core.VMCore, line, col int) error {
	a := vm.Pop()
	if a.Type != core.ValNumber {
		return vm.Errorf("E300 at %d:%d: want number", line, col)
	}
	vm.Push(core.NumberVal(-a.Data.(float64)))
	return nil
}
