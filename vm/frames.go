package vm

import "drylang/core"

type asyncTask struct {
	fn   *core.CompiledFn
	args []core.Value
}

type callFrame struct {
	closure *core.Closure
	ip      int
	bp      int // base pointer (stack offset)
	chunk   *core.Chunk
}

type tryFrame struct {
	catchIP    int
	sp         int
	errVarName string
	frameDepth int
}

func (vm *VM) GetCurrentClosure() *core.Closure {
	if len(vm.frames) == 0 {
		return nil
	}
	return vm.frames[len(vm.frames)-1].closure
}
