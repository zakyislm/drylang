package vm

import (
	"drylang/core"
	"fmt"
)

func (vm *VM) Push(val core.Value) {
	vm.push(val)
}

func (vm *VM) Pop() core.Value {
	return vm.pop()
}

func (vm *VM) Peek() core.Value {
	return vm.stack[vm.sp-1]
}

func (vm *VM) SetGlobal(name string, val core.Value) {
	vm.globals[name] = val
}

func (vm *VM) GetGlobal(name string) (core.Value, bool) {
	val, ok := vm.globals[name]
	return val, ok
}

func (vm *VM) SetLocal(slot int, val core.Value) {
	vm.stack[vm.frames[len(vm.frames)-1].bp+slot] = val
}

func (vm *VM) GetLocal(slot int) core.Value {
	return vm.stack[vm.frames[len(vm.frames)-1].bp+slot]
}

func (vm *VM) CallFunction(fn *core.CompiledFn, argCount int) error {
	return vm.callOp(argCount, 0, 0) // We can pass dummy line/col or update the signature
}

func (vm *VM) PushFrame(fn *core.CompiledFn, argCount int) error {
	frame := callFrame{
		fn:    fn,
		ip:    0,
		bp:    vm.sp - argCount,
		chunk: fn.Chunk,
	}
	vm.frames = append(vm.frames, frame)
	return nil
}

func (vm *VM) PopFrame() {
	vm.frames = vm.frames[:len(vm.frames)-1]
}

func (vm *VM) GetIP() int {
	return vm.ip
}

func (vm *VM) SetIP(ip int) {
	vm.ip = ip
}

func (vm *VM) GetChunk() *core.Chunk {
	return vm.chunk
}

func (vm *VM) Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func (vm *VM) Try(catchIP int) error {
	vm.tryStack = append(vm.tryStack, tryFrame{
		catchIP:    catchIP,
		sp:         vm.sp,
		frameDepth: len(vm.frames),
	})
	return nil
}

func (vm *VM) EndTry() error {
	if len(vm.tryStack) > 0 {
		vm.tryStack = vm.tryStack[:len(vm.tryStack)-1]
	}
	return nil
}

func (vm *VM) Throw() error {
	v := vm.pop()
	if len(vm.tryStack) > 0 {
		tf := vm.tryStack[len(vm.tryStack)-1]
		if tf.frameDepth == len(vm.frames) {
			vm.tryStack = vm.tryStack[:len(vm.tryStack)-1]
			vm.sp = tf.sp
			vm.push(v)
			vm.ip = tf.catchIP
			return nil
		}
	}
	return fmt.Errorf("%s", v.String())
}

