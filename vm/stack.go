package vm

import "drylang/core"

func (vm *VM) push(v core.Value) {
	if vm.sp >= len(vm.stack) {
		// Grow the stack instead of panicking on deep recursion.
		newCap := len(vm.stack) * 2
		if newCap < 16 {
			newCap = 16
		}
		bigger := make([]core.Value, newCap)
		copy(bigger, vm.stack)
		vm.stack = bigger
	}
	vm.stack[vm.sp] = v
	vm.sp++
}

func (vm *VM) pop() core.Value {
	vm.sp--
	return vm.stack[vm.sp]
}

func (vm *VM) peek() core.Value {
	return vm.stack[vm.sp-1]
}
