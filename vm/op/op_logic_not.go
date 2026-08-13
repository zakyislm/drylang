package op

import (
	"drylang/core"
)

func OpNot(vm core.VMCore, line, col int) error {
	a := vm.Pop()
	vm.Push(core.BoolVal(!core.IsTruthy(a)))
	return nil
}
