package iohandler

import (
	"drylang/core"
)

// BuiltinPipe handles pipe.make, pipe.send, pipe.recv
func BuiltinPipe(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("want pipe(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "make":
		// pipe.make(size)
		size := 0
		if len(args) > 1 && args[1].Type == core.ValNumber {
			size = int(args[1].Data.(float64))
		}
		ch := make(chan core.Value, size)
		return core.Value{Type: core.ValPipe, Data: ch}, nil

	case "send":
		// pipe.send(ch, val)
		if len(args) < 3 {
			return core.UnknownValue, vm.Errorf("pipe.send wants (pipe, value)")
		}
		if args[1].Type != core.ValPipe {
			return core.UnknownValue, vm.Errorf("pipe.send arg 1 must be pipe")
		}
		ch := args[1].Data.(chan core.Value)
		ch <- args[2]
		return core.UnknownValue, nil

	case "recv":
		// pipe.recv(ch)
		if len(args) < 2 {
			return core.UnknownValue, vm.Errorf("pipe.recv wants (pipe)")
		}
		if args[1].Type != core.ValPipe {
			return core.UnknownValue, vm.Errorf("pipe.recv arg 1 must be pipe")
		}
		ch := args[1].Data.(chan core.Value)
		val, ok := <-ch
		if !ok {
			return core.UnknownValue, vm.Errorf("pipe is closed")
		}
		return val, nil

	case "close":
		// pipe.close(ch)
		if len(args) < 2 {
			return core.UnknownValue, vm.Errorf("pipe.close wants (pipe)")
		}
		if args[1].Type != core.ValPipe {
			return core.UnknownValue, vm.Errorf("pipe.close arg 1 must be pipe")
		}
		ch := args[1].Data.(chan core.Value)
		close(ch)
		return core.UnknownValue, nil

	default:
		return core.UnknownValue, vm.Errorf("unknown pipe method: %s", method)
	}
}
