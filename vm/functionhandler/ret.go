package functionhandler

import (
	"drylang/core"
)

func OpReturn(vm core.VMCore) error {
	vm.PopFrame()
	return nil
}
