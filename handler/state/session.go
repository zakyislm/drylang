package state

import (
	"drylang/core"
	"sync"
	"time"
)

type sessionData struct {
	Store      map[string]core.Value
	ExpiryTime time.Time
}

var (
	sessMutex sync.Mutex
	sessStore = make(map[string]*sessionData)
)

func init() {
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			sessMutex.Lock()
			now := time.Now()
			for id, sess := range sessStore {
				if now.After(sess.ExpiryTime) {
					delete(sessStore, id)
				}
			}
			sessMutex.Unlock()
		}
	}()
}

// BuiltinSess handles sess.set, sess.get, sess.del
func BuiltinSess(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want sess(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "set":
		// sess.set(id, key, val, expirySecs)
		if len(args) != 5 {
			return core.UnknownValue, v.Errorf("sess.set wants (id, key, val, expirySecs)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValString || args[4].Type != core.ValNumber {
			return core.UnknownValue, v.Errorf("sess.set args must be (string, string, any, number)")
		}

		id := args[1].Data.(string)
		key := args[2].Data.(string)
		val := args[3]
		expirySecs := int(args[4].Data.(float64))

		sessMutex.Lock()
		defer sessMutex.Unlock()

		sess, exists := sessStore[id]
		if !exists {
			sess = &sessionData{
				Store: make(map[string]core.Value),
			}
			sessStore[id] = sess
		}
		
		sess.Store[key] = val
		sess.ExpiryTime = time.Now().Add(time.Duration(expirySecs) * time.Second)
		return core.UnknownValue, nil

	case "get":
		// sess.get(id, key)
		if len(args) != 3 {
			return core.UnknownValue, v.Errorf("sess.get wants (id, key)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValString {
			return core.UnknownValue, v.Errorf("sess.get args must be (string, string)")
		}

		id := args[1].Data.(string)
		key := args[2].Data.(string)

		sessMutex.Lock()
		defer sessMutex.Unlock()

		sess, exists := sessStore[id]
		if !exists || time.Now().After(sess.ExpiryTime) {
			return core.UnknownValue, nil
		}

		val, ok := sess.Store[key]
		if !ok {
			return core.UnknownValue, nil
		}
		return val, nil

	case "del":
		// sess.del(id)
		if len(args) != 2 || args[1].Type != core.ValString {
			return core.UnknownValue, v.Errorf("sess.del wants (id string)")
		}
		id := args[1].Data.(string)

		sessMutex.Lock()
		delete(sessStore, id)
		sessMutex.Unlock()
		return core.UnknownValue, nil

	default:
		return core.UnknownValue, v.Errorf("unknown sess method: %s", method)
	}
}
