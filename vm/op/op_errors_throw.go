package op
import "drylang/core"
func OpThrow(vm core.VMCore) error {
	return vm.Throw()
}
