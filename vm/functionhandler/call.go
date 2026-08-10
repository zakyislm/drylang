package functionhandler

import (
	"drylang/core"
)

func OpCall(vm core.VMCore, line, col, argCount int) error {
	// The implementation of OpCall is very complex and relies on vm internal frames,
	// so for this MVP we just delegate to a method on VMCore that does it natively,
	// or we can implement it natively in VMCore.
	// Since CallFunction handles OpCall completely in our updated VMCore logic, we just call it.
	return vm.CallFunction(nil, argCount) // nil because callee is on stack
}
