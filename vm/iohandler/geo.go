package iohandler

import (
	"drylang/core"
	"encoding/json"
	"io"
	"net/http"
)

// BuiltinGeo handles geo("ip", "1.1.1.1")
func BuiltinGeo(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want geo(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "ip":
		if len(args) != 2 || args[1].Type != core.ValString {
			return core.UnknownValue, v.Errorf("geo.ip wants (ip_string)")
		}
		ip := args[1].Data.(string)

		resp, err := http.Get("http://ip-api.com/json/" + ip)
		if err != nil {
			return core.UnknownValue, v.Errorf("geo.ip fetch error: %s", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return core.UnknownValue, v.Errorf("geo.ip read error: %s", err)
		}

		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			return core.UnknownValue, v.Errorf("geo.ip parse error: %s", err)
		}

		out := make(map[string]core.Value)
		for k, val := range data {
			switch t := val.(type) {
			case string:
				out[k] = core.Value{Type: core.ValString, Data: t}
			case float64:
				out[k] = core.Value{Type: core.ValNumber, Data: t}
			case bool:
				out[k] = core.Value{Type: core.ValBool, Data: t}
			}
		}

		return core.Value{Type: core.ValMap, Data: out}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown geo method: %s", method)
	}
}
