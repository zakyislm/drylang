package vm

import (
	"drylang/core"
)

func registerBuiltinModules(vm *VM) {
	for name, id := range core.BuiltinNames {
		vm.globals[name] = core.Value{
			Type: core.ValBuiltinModule,
			Data: core.BuiltinModule{ModuleID: int(id)},
		}
	}
}
