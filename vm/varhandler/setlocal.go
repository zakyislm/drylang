package varhandler

import (
	"drylang/core"
)

func OpSetLocal(vm core.VMCore, slot, line, col int) error {
	vm.SetLocal(slot, vm.Pop())
	return nil
}
