package op
import "drylang/core"
func OpTry(vm core.VMCore, catchOffset int) error {
	return vm.Try(catchOffset)
}
