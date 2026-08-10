package colshandler

import (
	"drylang/core"
)

func OpMap(vm core.VMCore, count int) error {
	m := make(map[string]core.Value)
	for i := 0; i < count; i++ {
		val := vm.Pop()
		key := vm.Pop()
		m[key.String()] = val
	}
	vm.Push(core.MapVal(m))
	return nil
}
