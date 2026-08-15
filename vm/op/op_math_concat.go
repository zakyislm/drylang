package op

import (
	"drylang/core"
)

func OpConcat(vm core.VMCore, line, col int) error {
	b, a := vm.Pop(), vm.Pop()
	vm.Push(core.StringVal(a.String() + b.String()))
	return nil
}
