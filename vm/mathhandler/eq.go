package mathhandler

import (
	"drylang/core"
)

func OpEqual(vm core.VMCore, line, col int) error {
	b, a := vm.Pop(), vm.Pop()
	vm.Push(core.BoolVal(a.Data == b.Data))
	return nil
}
