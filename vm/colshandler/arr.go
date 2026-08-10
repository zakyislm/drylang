package colshandler

import (
	"drylang/core"
)

func OpArray(vm core.VMCore, count int) error {
	arr := make([]core.Value, count)
	for i := count - 1; i >= 0; i-- {
		arr[i] = vm.Pop()
	}
	vm.Push(core.ArrayVal(arr))
	return nil
}
