package errorhandler
import "drylang/core"
func OpEndTry(vm core.VMCore) error {
	return vm.EndTry()
}
