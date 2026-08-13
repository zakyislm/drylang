package state

import (
	"drylang/core"
	"sync"
	"time"
)

type memData struct {
	Val        core.Value
	ExpiryTime time.Time
}

var (
	memMutex sync.RWMutex
	memStore = make(map[string]*memData)
)

func init() {
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			memMutex.Lock()
			now := time.Now()
			for k, m := range memStore {
				if now.After(m.ExpiryTime) {
					delete(memStore, k)
				}
			}
			memMutex.Unlock()
		}
	}()
}

// BuiltinMem handles mem("set", key, val, ttl) and mem("get", key)
func BuiltinMem(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want mem(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "set":
		if len(args) != 4 {
			return core.UnknownValue, v.Errorf("mem.set wants (key, val, ttl_sec)")
		}
		if args[1].Type != core.ValString || args[3].Type != core.ValNumber {
			return core.UnknownValue, v.Errorf("mem.set args must be (string, any, number)")
		}

		key := args[1].Data.(string)
		val := args[2]
		ttl := int(args[3].Data.(float64))

		memMutex.Lock()
		memStore[key] = &memData{
			Val:        val,
			ExpiryTime: time.Now().Add(time.Duration(ttl) * time.Second),
		}
		memMutex.Unlock()
		return core.UnknownValue, nil

	case "get":
		if len(args) != 2 || args[1].Type != core.ValString {
			return core.UnknownValue, v.Errorf("mem.get wants (key)")
		}
		key := args[1].Data.(string)

		memMutex.RLock()
		m, exists := memStore[key]
		memMutex.RUnlock()

		if !exists || time.Now().After(m.ExpiryTime) {
			return core.UnknownValue, nil
		}
		return m.Val, nil

	case "del":
		if len(args) != 2 || args[1].Type != core.ValString {
			return core.UnknownValue, v.Errorf("mem.del wants (key)")
		}
		key := args[1].Data.(string)

		memMutex.Lock()
		delete(memStore, key)
		memMutex.Unlock()
		return core.UnknownValue, nil

	default:
		return core.UnknownValue, v.Errorf("unknown mem method: %s", method)
	}
}
