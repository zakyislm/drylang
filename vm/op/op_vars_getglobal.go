package op

import (
	"drylang/core"
)

func OpGetGlobal(vm core.VMCore, name string, line, col int) error {
	if closure := vm.GetCurrentClosure(); closure != nil {
		if val, ok := closure.Env[name]; ok {
			vm.Push(val)
			return nil
		}
	}
	if val, ok := vm.GetGlobal(name); ok {
		vm.Push(val)
		return nil
	}
	return vm.Errorf("E301 at %d:%d: %s not found", line, col, name)
}
