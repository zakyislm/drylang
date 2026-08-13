package vm

import (
	"drylang/core"
	"drylang/vm/op"
)

// executeInner runs the dispatch loop over a chunk's instructions.
func (vm *VM) executeInner(chunk *core.Chunk) error {
	for vm.ip < len(chunk.Code) {
		inst := chunk.Code[vm.ip]
		line := 0
		col := 0
		if vm.ip < len(chunk.Lines) {
			line = chunk.Lines[vm.ip]
			col = chunk.Cols[vm.ip]
		}
		vm.ip++

		switch inst.Op {
		case core.OpConst:
			if err := vm.execConst(chunk, inst, line, col); err != nil {
				return err
			}

		case core.OpPop:
			if vm.sp > 0 {
				vm.pop()
			}

		case core.OpTrue:
			vm.push(core.BoolVal(true))
		case core.OpFalse:
			vm.push(core.BoolVal(false))
		case core.OpUnknown:
			vm.push(core.UnknownValue)

		case core.OpAdd:
			if err := op.OpAdd(vm, line, col); err != nil {
				return err
			}

		case core.OpSub:
			if err := op.OpSub(vm, line, col); err != nil {
				return err
			}

		case core.OpMul:
			if err := op.OpMul(vm, line, col); err != nil {
				return err
			}

		case core.OpDiv:
			if err := op.OpDiv(vm, line, col); err != nil {
				return err
			}

		case core.OpMod:
			if err := op.OpMod(vm, line, col); err != nil {
				return err
			}

		case core.OpNeg:
			v := vm.pop()
			if v.Type != core.ValNumber {
				return vm.runtimeErr("E300", line, col, "want number")
			}
			vm.push(core.NumberVal(-v.Data.(float64)))

		case core.OpEqual:
			if err := op.OpEqual(vm, line, col); err != nil {
				return err
			}

		case core.OpNotEqual:
			if err := op.OpNotEqual(vm, line, col); err != nil {
				return err
			}

		case core.OpLess:
			if err := op.OpLess(vm, line, col); err != nil {
				return err
			}

		case core.OpGreater:
			if err := op.OpGreater(vm, line, col); err != nil {
				return err
			}

		case core.OpLessEq:
			if err := op.OpLessEq(vm, line, col); err != nil {
				return err
			}

		case core.OpGreaterEq:
			if err := op.OpGreaterEq(vm, line, col); err != nil {
				return err
			}

		case core.OpAnd:
			if err := op.OpAnd(vm, line, col); err != nil {
				return err
			}

		case core.OpOr:
			if err := op.OpOr(vm, line, col); err != nil {
				return err
			}

		case core.OpNot:
			v := vm.pop()
			vm.push(core.BoolVal(!core.IsTruthy(v)))

		case core.OpConcat:
			if err := op.OpConcat(vm, line, col); err != nil {
				return err
			}

		case core.OpGetGlobal:
			name := chunk.Constants[inst.Operand].(string)
			if err := op.OpGetGlobal(vm, name, line, col); err != nil {
				// varhandler returns E301 for not found, but tests expect E300 "unknown x"
				if err.Error()[:4] == "E301" {
					return vm.runtimeErr("E300", line, col, "unknown %s", name)
				}
				return err
			}

		case core.OpSetGlobal:
			name := chunk.Constants[inst.Operand].(string)
			if err := op.OpSetGlobal(vm, name, line, col); err != nil {
				return err
			}

		case core.OpGetLocal:
			if err := op.OpGetLocal(vm, inst.Operand, line, col); err != nil {
				return err
			}

		case core.OpSetLocal:
			if err := op.OpSetLocal(vm, inst.Operand, line, col); err != nil {
				return err
			}

		case core.OpJump:
			if err := op.OpJump(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpJumpIfFalse:
			if err := op.OpJumpIfFalse(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpJumpIfNotUnknown:
			if err := op.OpJumpIfNotUnknown(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpLoop:
			if err := op.OpLoop(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpArray:
			if err := op.OpArray(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpMap:
			if err := op.OpMap(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpIndex:
			if err := op.OpIndex(vm, line, col); err != nil {
				return err
			}

		case core.OpSetIndex:
			if err := op.OpSetIndex(vm, line, col); err != nil {
				return err
			}

		case core.OpDotGet:
			field := chunk.Constants[inst.Operand].(string)
			if err := op.OpDotGet(vm, line, col, field, false); err != nil {
				return err
			}

		case core.OpOptDotGet:
			field := chunk.Constants[inst.Operand].(string)
			if err := op.OpDotGet(vm, line, col, field, true); err != nil {
				return err
			}

		case core.OpDotSet:
			field := chunk.Constants[inst.Operand].(string)
			if err := op.OpDotSet(vm, line, col, field); err != nil {
				return err
			}

		case core.OpClosure:
			fn := chunk.Constants[inst.Operand].(*core.CompiledFn)
			if err := op.OpClosure(vm, fn); err != nil {
				return err
			}

		case core.OpCall:
			if err := op.OpCall(vm, line, col, inst.Operand); err != nil {
				return err
			}

		case core.OpReturn:
			if len(vm.frames) > 0 {
				return nil // return from function execution
			}
			// Top-level rev = exit
			return nil

		case core.OpTry:
			if err := op.OpTry(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpEndTry:
			if err := op.OpEndTry(vm); err != nil {
				return err
			}

		case core.OpThrow:
			if err := op.OpThrow(vm); err != nil {
				return err
			}

		case core.OpBuiltin:
			if err := vm.executeBuiltin(core.BuiltinID(inst.Operand), inst.Operand2, line, col); err != nil {
				return err
			}

		case core.OpGetLoopCounter:
			// handled via locals

		case core.OpAsyncCall:
			if err := vm.execAsyncCall(inst, line, col); err != nil {
				return err
			}

		case core.OpAwait:
			vm.asyncWg.Wait()
			// ponytail: awt returns void for now (doesn't push a promise result)
		}
	}

	return nil
}

// execConst pushes a constant value from the chunk's constant pool.
func (vm *VM) execConst(chunk *core.Chunk, inst core.Instruction, line, col int) error {
	val := chunk.Constants[inst.Operand]
	switch v := val.(type) {
	case float64:
		vm.push(core.NumberVal(v))
	case string:
		vm.push(core.StringVal(v))
	case *core.CompiledFn:
		vm.push(core.FnVal(&core.Closure{Fn: v, Env: make(map[string]core.Value)}))
	case core.StructDef:
		vm.push(core.Value{Type: core.ValStructDef, Data: v})
	case core.ClassDef:
		// Resolve parents at runtime
		var resolvedParents []*core.ClassDef
		for _, pName := range v.ParentNames {
				if parentVal, ok := vm.GetGlobal(pName); ok && parentVal.Type == core.ValClass {
				pDef := parentVal.Data.(core.ClassDef)
				resolvedParents = append(resolvedParents, &pDef)
			} else {
				return vm.Errorf("E300 at %d:%d: parent class '%s' not found", line, col, pName)
			}
		}
		v.Parents = resolvedParents
		vm.push(core.Value{Type: core.ValClass, Data: v})
	default:
		vm.push(core.Value{Type: core.ValUnknown, Data: nil})
	}
	return nil
}
