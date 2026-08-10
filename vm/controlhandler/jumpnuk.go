package controlhandler

import (
	"drylang/core"
)

func OpJumpIfNotUnknown(vm core.VMCore, target int) error {
	val := vm.Peek()
	if val.Type != core.ValUnknown {
		vm.SetIP(target)
	}
	return nil
}
