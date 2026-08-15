package http

import (
	"crypto/sha1"
	"drylang/core"
	"encoding/base64"
)

var wsGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")

func computeAcceptKey(challengeKey string) string {
	h := sha1.New()
	h.Write([]byte(challengeKey))
	h.Write(wsGUID)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// BuiltinWs handles ws("serve", port)
func BuiltinWs(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want ws(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "serve":
		// ws("serve", port)
		if len(args) != 2 || args[1].Type != core.ValNumber {
			return core.UnknownValue, v.Errorf("ws.serve wants (port)")
		}
		
		return core.Value{Type: core.ValBool, Data: true}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown ws method: %s", method)
	}
}
