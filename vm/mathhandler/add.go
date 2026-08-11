package mathhandler

import (
	"drylang/core"
)

func OpAdd(vm core.VMCore, line, col int) error {
	b, a := vm.Pop(), vm.Pop()
	if a.Type == core.ValString || b.Type == core.ValString {
		vm.Push(core.StringVal(a.String() + b.String()))
	} else if a.Type == core.ValNumber && b.Type == core.ValNumber {
		vm.Push(core.NumberVal(a.Data.(float64) + b.Data.(float64)))
	} else if a.Type == core.ValArray && b.Type == core.ValArray {
		arrA := a.Data.([]core.Value)
		arrB := b.Data.([]core.Value)
		newArr := make([]core.Value, 0, len(arrA)+len(arrB))
		newArr = append(newArr, arrA...)
		newArr = append(newArr, arrB...)
		vm.Push(core.ArrayVal(newArr))
	} else {
		return vm.Errorf("E300 at %d:%d: cannot add %s + %s", line, col, a.Type, b.Type)
	}
	return nil
}
