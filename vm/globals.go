package vm

import "drylang/core"

func (vm *VM) SetGlobal(name string, val core.Value) {
	vm.gmu.Lock()
	vm.globals[name] = val
	vm.gmu.Unlock()
}

func (vm *VM) GetGlobal(name string) (core.Value, bool) {
	vm.gmu.RLock()
	val, ok := vm.globals[name]
	vm.gmu.RUnlock()
	return val, ok
}

// CloneGlobals returns a shallow copy of the globals map so a spawned VM
// (async task or HTTP request) owns its own variable space. Values are shared;
// keys are isolated. This is what makes concurrent execution race-free: no two
// goroutines ever write the same map.
func (vm *VM) CloneGlobals() map[string]core.Value {
	vm.gmu.RLock()
	defer vm.gmu.RUnlock()
	cp := make(map[string]core.Value, len(vm.globals))
	for k, v := range vm.globals {
		cp[k] = v
	}
	return cp
}
