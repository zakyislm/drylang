package colshandler

import (
	"drylang/core"
)

func OpIndex(vm core.VMCore, line, col int) error {
	idx, obj := vm.Pop(), vm.Pop()
	if obj.Type == core.ValArray {

		arr := obj.Data.([]core.Value)
		var i int
		switch v := idx.Data.(type) {
		case float64:
			i = int(v)
		case int:
			i = v
		default:
			return vm.Errorf("E302 at %d:%d: array index must be number", line, col)
		}
		if i < 0 || i >= len(arr) {
			return vm.Errorf("E302 at %d:%d: array index out of bounds", line, col)
		}
		vm.Push(arr[i])
	} else if obj.Type == core.ValMap {
		m := obj.Data.(map[string]core.Value)
		k := idx.String()
		if val, ok := m[k]; ok {
			vm.Push(val)
		} else {
			vm.Push(core.UnknownValue)
		}
	} else if obj.Type == core.ValString {

		s := obj.Data.(string)
		var i int
		switch v := idx.Data.(type) {
		case float64:
			i = int(v)
		case int:
			i = v
		default:
			return vm.Errorf("E302 at %d:%d: string index must be number", line, col)
		}
		if i < 0 || i >= len(s) {
			return vm.Errorf("E302 at %d:%d: string index out of bounds", line, col)
		}
		vm.Push(core.StringVal(string(s[i])))
	} else {
		return vm.Errorf("E302 at %d:%d: object not indexable", line, col)
	}
	return nil
}
