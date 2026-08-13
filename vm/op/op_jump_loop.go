package op

import (
	"drylang/core"
)

func OpLoop(vm core.VMCore, offset int) error {
	vm.SetIP(vm.GetIP() - offset)
	return nil
}
