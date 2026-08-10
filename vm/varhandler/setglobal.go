package varhandler

import (
	"drylang/core"
)

func OpSetGlobal(vm core.VMCore, name string, line, col int) error {
	vm.SetGlobal(name, vm.Pop())
	return nil
}
