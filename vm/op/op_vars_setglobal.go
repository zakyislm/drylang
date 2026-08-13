package op

import (
	"drylang/core"
)

func OpSetGlobal(vm core.VMCore, name string, line, col int) error {
	val := vm.Pop()
	if closure := vm.GetCurrentClosure(); closure != nil {
		if _, ok := closure.Env[name]; ok {
			closure.Env[name] = val
			return nil
		}
	}
	vm.SetGlobal(name, val)
	return nil
}
