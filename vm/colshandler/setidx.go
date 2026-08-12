package colshandler

import (
	"drylang/core"
)

func OpSetIndex(vm core.VMCore, line, col int) error {
	val, idx, obj := vm.Pop(), vm.Pop(), vm.Pop()
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
		arr[i] = val
	} else if obj.Type == core.ValMap {
		m := obj.Data.(map[string]core.Value)
		k := idx.String()
		m[k] = val
	} else {
		return vm.Errorf("E302 at %d:%d: object not indexable for assignment", line, col)
	}
	return nil
}
