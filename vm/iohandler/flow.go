package iohandler

import (
	"drylang/core"
	"sync"
)

type flowState struct {
	State string
}

var (
	flowMutex sync.RWMutex
	flowStore = make(map[string]*flowState)
)

// BuiltinFlow handles flow("create", id, initialState) and flow("transition", id, nextState)
func BuiltinFlow(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want flow(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "create":
		if len(args) != 3 {
			return core.UnknownValue, v.Errorf("flow.create wants (id, initialState)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValString {
			return core.UnknownValue, v.Errorf("flow.create args must be (string, string)")
		}
		id := args[1].Data.(string)
		state := args[2].Data.(string)

		flowMutex.Lock()
		flowStore[id] = &flowState{State: state}
		flowMutex.Unlock()
		return core.Value{Type: core.ValBool, Data: true}, nil

	case "transition":
		if len(args) != 3 {
			return core.UnknownValue, v.Errorf("flow.transition wants (id, nextState)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValString {
			return core.UnknownValue, v.Errorf("flow.transition args must be (string, string)")
		}
		id := args[1].Data.(string)
		nextState := args[2].Data.(string)

		flowMutex.Lock()
		f, exists := flowStore[id]
		if !exists {
			flowMutex.Unlock()
			return core.Value{Type: core.ValBool, Data: false}, nil
		}
		f.State = nextState
		flowMutex.Unlock()
		return core.Value{Type: core.ValBool, Data: true}, nil

	case "get":
		if len(args) != 2 || args[1].Type != core.ValString {
			return core.UnknownValue, v.Errorf("flow.get wants (id)")
		}
		id := args[1].Data.(string)

		flowMutex.RLock()
		f, exists := flowStore[id]
		flowMutex.RUnlock()

		if !exists {
			return core.UnknownValue, nil
		}
		return core.Value{Type: core.ValString, Data: f.State}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown flow method: %s", method)
	}
}
