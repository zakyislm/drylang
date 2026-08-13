package iohandler

import (
	"drylang/core"
	"runtime"
)

// BuiltinMet handles met("sys")
func BuiltinMet(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want met(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "sys":
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		out := make(map[string]core.Value)
		out["goroutines"] = core.Value{Type: core.ValNumber, Data: float64(runtime.NumGoroutine())}
		out["alloc_bytes"] = core.Value{Type: core.ValNumber, Data: float64(m.Alloc)}
		out["sys_bytes"] = core.Value{Type: core.ValNumber, Data: float64(m.Sys)}
		out["num_gc"] = core.Value{Type: core.ValNumber, Data: float64(m.NumGC)}

		return core.Value{Type: core.ValMap, Data: out}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown met method: %s", method)
	}
}
