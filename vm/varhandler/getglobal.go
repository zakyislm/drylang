package varhandler

import (
	"drylang/core"
)

func OpGetGlobal(vm core.VMCore, name string, line, col int) error {
	if val, ok := vm.GetGlobal(name); ok {
		vm.Push(val)
		return nil
	}
	return vm.Errorf("E301 at %d:%d: %s not found", line, col, name)
}
