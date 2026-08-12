package controlhandler

import (
	"drylang/core"
)

func OpJumpIfFalse(vm core.VMCore, target int) error {
	val := vm.Pop()
	if !core.IsTruthy(val) {
		vm.SetIP(target)
	}
	return nil
}
