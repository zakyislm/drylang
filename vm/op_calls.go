package vm

import (
	"drylang/core"
)

// execAsyncCall spawns an async task on a per-function worker pool.
func (vm *VM) execAsyncCall(inst core.Instruction, line, col int) error {
	argCount := inst.Operand
	workers := inst.Operand2

	args := make([]core.Value, argCount)
	for i := argCount - 1; i >= 0; i-- {
		args[i] = vm.pop()
	}
	callee := vm.pop()

	if callee.Type != core.ValFn {
		return vm.runtimeErr("E300", line, col, "want fn for async call")
	}
	closureVal := callee.Data.(*core.Closure)
	fnVal := closureVal.Fn

	if !fnVal.IsAsync {
		return vm.runtimeErr("E300", line, col, "cannot mul a non-async fn")
	}

	vm.poolMutex.Lock()
	pool, exists := vm.asyncPools[fnVal]
	if !exists {
		pool = make(chan asyncTask, 1000) // generous buffer
		vm.asyncPools[fnVal] = pool

		// Spawn workers
		for i := 0; i < workers; i++ {
			go func(ch chan asyncTask) {
				for task := range ch {
					execVM := &VM{
						chunk:      vm.chunk,
						fns:        vm.fns,
						globals:    vm.CloneGlobals(),
						stack:      make([]core.Value, 4096),
						sp:         0,
						ip:         0,
						asyncPools: vm.asyncPools, // share pools for nested calls
					}

					execVM.push(core.Value{Type: core.ValFn, Data: &core.Closure{Fn: task.fn, Env: make(map[string]core.Value)}})
					for _, a := range task.args {
						execVM.push(a)
					}

					// CallFunction internally sets up the frame and calls vm.execute()
					execVM.CallFunction(task.fn, len(task.args))

					vm.asyncWg.Done()
				}
			}(pool)
		}
	}
	vm.poolMutex.Unlock()

	vm.asyncWg.Add(1)
	pool <- asyncTask{fn: fnVal, args: args}
	return nil
}

func (vm *VM) callOp(argCount, line, col int) error {
	chunk := vm.chunk
	callee := vm.stack[vm.sp-argCount-1]

	if callee.Type == core.ValStructDef {
		def := callee.Data.(core.StructDef)
		if argCount != len(def.Fields) {
			return vm.runtimeErr("E300", line, col, "want %d args for struct", len(def.Fields))
		}
		fields := make(map[string]core.Value)
		fields["__struct__"] = core.StringVal(def.Name)
		for i := 0; i < argCount; i++ {
			fields[def.Fields[i]] = vm.stack[vm.sp-argCount+i]
		}
		vm.sp -= argCount + 1
		vm.push(core.Value{Type: core.ValStructInstance, Data: fields})
		return nil
	}

	if callee.Type == core.ValClass {
		return vm.callClass(callee, argCount, line, col, chunk)
	}

	if callee.Type == core.ValBoundMethod {
		return vm.callBoundMethod(callee, argCount, line, col, chunk)
	}

	if callee.Type == core.ValBuiltinFn {
		return vm.callBuiltinFn(callee, argCount, line, col)
	}

	if callee.Type != core.ValFn {
		return vm.runtimeErr("E300", line, col, "want fn")
	}

	closure := callee.Data.(*core.Closure)
	fn := closure.Fn
	if argCount != fn.ParamCount {
		return vm.runtimeErr("E300", line, col, "want %d args", fn.ParamCount)
	}

	frame := callFrame{
		closure: closure,
		ip:    vm.ip,
		bp:    vm.sp - argCount,
		chunk: chunk,
	}
	vm.frames = append(vm.frames, frame)

	for vm.sp < frame.bp+fn.Chunk.LocalsCount+1 {
		vm.push(core.UnknownValue)
	}

	savedIP := vm.ip
	savedChunk := chunk
	err := vm.execute(fn.Chunk)
	if err != nil {
		return err
	}

	vm.frames = vm.frames[:len(vm.frames)-1]
	result := vm.pop()

	vm.sp = frame.bp - 1
	if vm.sp < 0 {
		vm.sp = 0
	}
	vm.push(result)

	vm.ip = savedIP
	chunk = savedChunk
	vm.chunk = chunk
	return nil
}

func (vm *VM) callClass(callee core.Value, argCount, line, col int, chunk *core.Chunk) error {
	def := callee.Data.(core.ClassDef)
	inst := &core.Instance{
		Class:  def,
		Fields: make(map[string]core.Value),
	}

	// Initialize fields from all parents (DFS)
	var initFields func(c *core.ClassDef)
	initFields = func(c *core.ClassDef) {
		for _, p := range c.Parents {
			if p != nil {
				initFields(p)
			}
		}
		for _, fieldName := range c.Fields {
			inst.Fields[fieldName] = core.UnknownValue
		}
	}
	initFields(&def)

	// Find init method on the class or any parent (DFS).
	var findInit func(c *core.ClassDef) *core.ClassMethod
	findInit = func(c *core.ClassDef) *core.ClassMethod {
		if m, ok := c.Methods["init"]; ok {
			return &m
		}
		for _, p := range c.Parents {
			if p == nil {
				continue
			}
			if found := findInit(p); found != nil {
				return found
			}
		}
		return nil
	}

	if initM := findInit(&def); initM != nil {
		vm.stack[vm.sp-argCount-1] = core.Value{Type: core.ValInstance, Data: inst}

		frame := callFrame{
			closure: &core.Closure{
				Fn: &core.CompiledFn{Chunk: initM.Chunk, Name: "init", ParamCount: argCount},
				Env: make(map[string]core.Value),
			},
			ip:    vm.ip,
			bp:    vm.sp - argCount - 1,
			chunk: chunk,
		}
		vm.frames = append(vm.frames, frame)

		savedIP := vm.ip
		savedChunk := chunk
		err := vm.execute(initM.Chunk)
		if err != nil {
			return err
		}

		vm.frames = vm.frames[:len(vm.frames)-1]
		vm.pop() // init returns nothing

		vm.sp = frame.bp
		if vm.sp < 0 {
			vm.sp = 0
		}
		vm.push(core.Value{Type: core.ValInstance, Data: inst})

		vm.ip = savedIP
		chunk = savedChunk
		vm.chunk = chunk
		return nil
	}

	if argCount > 0 {
		return vm.runtimeErr("E300", line, col, "class has no init but got args")
	}
	vm.sp -= argCount + 1
	vm.push(core.Value{Type: core.ValInstance, Data: inst})
	return nil
}

func (vm *VM) callBoundMethod(callee core.Value, argCount, line, col int, chunk *core.Chunk) error {
	bound := callee.Data.(*core.BoundMethod)
	vm.stack[vm.sp-argCount-1] = core.Value{Type: core.ValInstance, Data: bound.Instance}

	frame := callFrame{
		closure: &core.Closure{
			Fn: &core.CompiledFn{Chunk: bound.Method.Chunk, Name: "<method>", ParamCount: argCount},
			Env: make(map[string]core.Value),
		},
		ip:    vm.ip,
		bp:    vm.sp - argCount - 1,
		chunk: chunk,
	}
	vm.frames = append(vm.frames, frame)

	for vm.sp < frame.bp+bound.Method.Chunk.LocalsCount+1 {
		vm.push(core.UnknownValue)
	}

	savedIP := vm.ip
	savedChunk := chunk
	err := vm.execute(bound.Method.Chunk)
	if err != nil {
		return err
	}

	vm.frames = vm.frames[:len(vm.frames)-1]
	result := vm.pop()

	vm.sp = frame.bp
	if vm.sp < 0 {
		vm.sp = 0
	}
	vm.push(result)

	vm.ip = savedIP
	chunk = savedChunk
	vm.chunk = chunk
	return nil
}

func (vm *VM) callBuiltinFn(callee core.Value, argCount, line, col int) error {
	bFn := callee.Data.(core.BuiltinFn)
	// Builtins expect the method name as the first string argument.
	// Shift arguments up by 1 on the stack and insert the method string.
	for i := 0; i < argCount; i++ {
		vm.stack[vm.sp-i] = vm.stack[vm.sp-i-1]
	}
	vm.stack[vm.sp-argCount] = core.StringVal(bFn.Method)
	vm.sp++ // We added one argument

	// Now call executeBuiltin with argCount + 1
	err := vm.executeBuiltin(core.BuiltinID(bFn.ModuleID), argCount+1, line, col)
	if err != nil {
		return err
	}

	// executeBuiltin pushed the result. Remove the callee.
	result := vm.pop()
	vm.pop() // remove ValBuiltinFn
	vm.push(result)
	return nil
}
