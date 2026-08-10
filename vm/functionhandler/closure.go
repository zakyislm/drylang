package functionhandler

import (
	"drylang/core"
)

func OpClosure(vm core.VMCore, fn *core.CompiledFn) error {
	vm.Push(core.FnVal(fn))
	return nil
}
