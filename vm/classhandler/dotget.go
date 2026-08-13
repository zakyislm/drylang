package classhandler

import (
	"drylang/core"
)

func OpDotGet(vm core.VMCore, line, col int, fieldName string, optional bool) error {
	obj := vm.Pop()
	if optional && obj.Type == core.ValUnknown {
		vm.Push(core.Value{Type: core.ValUnknown})
		return nil
	}
	
	if obj.Type == core.ValBuiltinModule {
		mod := obj.Data.(core.BuiltinModule)
		vm.Push(core.Value{
			Type: core.ValBuiltinFn,
			Data: core.BuiltinFn{ModuleID: mod.ModuleID, Method: fieldName},
		})
		return nil
	}

	if obj.Type == core.ValStructInstance {
		fields := obj.Data.(map[string]core.Value)
		if val, ok := fields[fieldName]; ok {
			vm.Push(val)
			return nil
		}
		if optional {
			vm.Push(core.Value{Type: core.ValUnknown})
			return nil
		}
		return vm.Errorf("E303 at %d:%d: undefined property: %s", line, col, fieldName)
	}

	if obj.Type != core.ValInstance {
		if optional {
			vm.Push(core.Value{Type: core.ValUnknown})
			return nil
		}
		return vm.Errorf("E303 at %d:%d: property access on non-instance", line, col)
	}

	instObj := obj.Data.(*core.Instance)
	if m, ok := instObj.Class.Methods[fieldName]; ok {
		vm.Push(core.Value{Type: core.ValBoundMethod, Data: &core.BoundMethod{Instance: instObj, Method: m}})
		return nil
	}
	
	// Search in parents (DFS)
	var findMethod func(c *core.ClassDef) *core.ClassMethod
	findMethod = func(c *core.ClassDef) *core.ClassMethod {
		for _, p := range c.Parents {
			if p == nil {
				continue
			}
			if m, ok := p.Methods[fieldName]; ok {
				return &m
			}
			if found := findMethod(p); found != nil {
				return found
			}
		}
		return nil
	}
	
	if foundM := findMethod(&instObj.Class); foundM != nil {
		vm.Push(core.Value{Type: core.ValBoundMethod, Data: &core.BoundMethod{Instance: instObj, Method: *foundM}})
		return nil
	}
	if val, ok := instObj.Fields[fieldName]; ok {
		vm.Push(val)
		return nil
	}
	
	if optional {
		vm.Push(core.Value{Type: core.ValUnknown})
		return nil
	}
	return vm.Errorf("E303 at %d:%d: undefined property: %s", line, col, fieldName)
}

