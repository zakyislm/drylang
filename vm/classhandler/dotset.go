package classhandler

import (
	"drylang/core"
)

func OpDotSet(vm core.VMCore, line, col int, fieldName string) error {
	val, obj := vm.Pop(), vm.Pop()
	if obj.Type == core.ValInstance {
		instObj := obj.Data.(*core.Instance)
		instObj.Fields[fieldName] = val
	} else {
		return vm.Errorf("E303 at %d:%d: property set on non-instance", line, col)
	}
	return nil
}
