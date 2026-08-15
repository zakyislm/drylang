package op

import (
	"drylang/core"
)

func OpOr(vm core.VMCore, line, col int) error {
	b, a := vm.Pop(), vm.Pop()
	vm.Push(core.BoolVal(core.IsTruthy(a) || core.IsTruthy(b)))
	return nil
}
