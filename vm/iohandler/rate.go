package iohandler

import (
	"drylang/core"
	"sync"
	"time"
)

type rateLimitEntry struct {
	Count      int
	WindowEnds time.Time
}

var (
	rateMutex sync.Mutex
	rateStore = make(map[string]*rateLimitEntry)
)

// cleanup loop for rate limiter
func init() {
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			rateMutex.Lock()
			now := time.Now()
			for ip, entry := range rateStore {
				if now.After(entry.WindowEnds) {
					delete(rateStore, ip)
				}
			}
			rateMutex.Unlock()
		}
	}()
}

// BuiltinRate handles rate.check(ip, maxReq, windowSec)
func BuiltinRate(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want rate(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "check":
		// rate.check(ip, maxReq, windowSec)
		if len(args) != 4 {
			return core.UnknownValue, v.Errorf("rate.check wants (ip, maxReq, windowSec)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValNumber || args[3].Type != core.ValNumber {
			return core.UnknownValue, v.Errorf("rate.check args must be (string, number, number)")
		}

		ip := args[1].Data.(string)
		maxReq := int(args[2].Data.(float64))
		windowSec := int(args[3].Data.(float64))

		rateMutex.Lock()
		defer rateMutex.Unlock()

		now := time.Now()
		entry, exists := rateStore[ip]

		if !exists || now.After(entry.WindowEnds) {
			rateStore[ip] = &rateLimitEntry{
				Count:      1,
				WindowEnds: now.Add(time.Duration(windowSec) * time.Second),
			}
			return core.Value{Type: core.ValBool, Data: true}, nil // Allowed
		}

		if entry.Count >= maxReq {
			return core.Value{Type: core.ValBool, Data: false}, nil // Rate Limited (Denied)
		}

		entry.Count++
		return core.Value{Type: core.ValBool, Data: true}, nil // Allowed

	default:
		return core.UnknownValue, v.Errorf("unknown rate method: %s", method)
	}
}
