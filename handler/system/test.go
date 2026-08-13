package system

import (
	"drylang/core"
)

// BuiltinTest handles testing framework:
// test("eq", a, b, msg)
// test("neq", a, b, msg)
func BuiltinTest(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want test(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "eq":
		if len(args) < 3 {
			return core.UnknownValue, v.Errorf("test.eq wants (a, b, [msg])")
		}
		a := args[1]
		b := args[2]
		msg := "Assertion failed: not equal"
		if len(args) >= 4 && args[3].Type == core.ValString {
			msg = args[3].Data.(string)
		}

		// Simple equality check
		if a.Type != b.Type || a.Data != b.Data {
			return core.UnknownValue, v.Errorf("TEST FAIL: %s", msg)
		}
		return core.Value{Type: core.ValBool, Data: true}, nil

	case "neq":
		if len(args) < 3 {
			return core.UnknownValue, v.Errorf("test.neq wants (a, b, [msg])")
		}
		a := args[1]
		b := args[2]
		msg := "Assertion failed: equal"
		if len(args) >= 4 && args[3].Type == core.ValString {
			msg = args[3].Data.(string)
		}

		if a.Type == b.Type && a.Data == b.Data {
			return core.UnknownValue, v.Errorf("TEST FAIL: %s", msg)
		}
		return core.Value{Type: core.ValBool, Data: true}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown test method: %s", method)
	}
}
