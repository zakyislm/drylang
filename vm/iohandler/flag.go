package iohandler

import (
	"drylang/core"
	"os"
	"strings"
	"sync"
)

var (
	flagMutex sync.RWMutex
	flagStore = make(map[string]bool)
)

// BuiltinFlag handles flag.set, flag.check
func BuiltinFlag(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want flag(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "set":
		// flag("set", name, bool)
		if len(args) != 3 {
			return core.UnknownValue, v.Errorf("flag.set wants (name, bool)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValBool {
			return core.UnknownValue, v.Errorf("flag.set args must be (string, bool)")
		}
		name := args[1].Data.(string)
		val := args[2].Data.(bool)

		flagMutex.Lock()
		flagStore[name] = val
		flagMutex.Unlock()
		return core.UnknownValue, nil

	case "check":
		// flag("check", name)
		if len(args) != 2 || args[1].Type != core.ValString {
			return core.UnknownValue, v.Errorf("flag.check wants (name string)")
		}
		name := args[1].Data.(string)

		// First check in-memory flags
		flagMutex.RLock()
		val, exists := flagStore[name]
		flagMutex.RUnlock()

		if exists {
			return core.Value{Type: core.ValBool, Data: val}, nil
		}

		// Fallback to environment variables
		envVal := strings.ToLower(os.Getenv("FLAG_" + strings.ToUpper(name)))
		if envVal == "true" || envVal == "1" || envVal == "t" {
			return core.Value{Type: core.ValBool, Data: true}, nil
		}

		return core.Value{Type: core.ValBool, Data: false}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown flag method: %s", method)
	}
}
