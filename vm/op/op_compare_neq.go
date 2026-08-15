package op

import (
	"drylang/core"
)

func OpNotEqual(vm core.VMCore, line, col int) error {
	b, a := vm.Pop(), vm.Pop()
	vm.Push(core.BoolVal(a.Data != b.Data))
	return nil
}
