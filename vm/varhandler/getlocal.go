package varhandler

import (
	"drylang/core"
)

func OpGetLocal(vm core.VMCore, slot, line, col int) error {
	vm.Push(vm.GetLocal(slot))
	return nil
}
