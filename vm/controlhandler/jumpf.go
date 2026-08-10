package controlhandler

import (
	"drylang/core"
)

func OpJumpIfFalse(vm core.VMCore, target int) error {
	if !core.IsTruthy(vm.Peek()) {
		vm.SetIP(target)
	}
	return nil
}
