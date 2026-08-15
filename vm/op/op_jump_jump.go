package op

import (
	"drylang/core"
)

func OpJump(vm core.VMCore, target int) error {
	vm.SetIP(target)
	return nil
}
