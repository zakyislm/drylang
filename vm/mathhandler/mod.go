package mathhandler

import (
	"drylang/core"
	"math"
)

func OpMod(vm core.VMCore, line, col int) error {
	b, a := vm.Pop(), vm.Pop()
	if a.Type != core.ValNumber || b.Type != core.ValNumber {
		return vm.Errorf("E300 at %d:%d: want number", line, col)
	}
	vm.Push(core.NumberVal(math.Mod(a.Data.(float64), b.Data.(float64))))
	return nil
}
