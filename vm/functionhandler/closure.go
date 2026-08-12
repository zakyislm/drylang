package functionhandler

import (
	"drylang/core"
)

func OpClosure(vm core.VMCore, fn *core.CompiledFn) error {
	env := make(map[string]core.Value)

	// Inherit from parent closure (dynamic environment copy)
	if parent := vm.GetCurrentClosure(); parent != nil {
		for k, v := range parent.Env {
			env[k] = v
		}
		// Also capture current locals if available
		if parent.Fn != nil && parent.Fn.LocalNames != nil {
			for slot, name := range parent.Fn.LocalNames {
				env[name] = vm.GetLocal(slot)
			}
		}
	}

	vm.Push(core.FnVal(&core.Closure{Fn: fn, Env: env}))
	return nil
}
