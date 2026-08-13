package data

import (
	"drylang/core"
	"encoding/json"
)

func BuiltinJson(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) != 1 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want 1 arg", line, col)
	}

	if args[0].Type == core.ValString {
		// Parse JSON string
		jsonStr := args[0].Data.(string)
		parsed, err := parseJSON(jsonStr)
		if err != nil {
			return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"json parse fail: %v", err, line, col)
		}
		return parsed, nil
	}

	// Stringify to JSON
	raw := valueToInterface(args[0])
	bytes, err := json.Marshal(raw)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"json stringify fail: %v", err, line, col)
	}
	return core.StringVal(string(bytes)), nil
}

func valueToInterface(v core.Value) interface{} {
	switch v.Type {
	case core.ValNumber:
		return v.Data.(float64)
	case core.ValString:
		return v.Data.(string)
	case core.ValBool:
		return v.Data.(bool)
	case core.ValArray:
		arr := v.Data.([]core.Value)
		res := make([]interface{}, len(arr))
		for i, item := range arr {
			res[i] = valueToInterface(item)
		}
		return res
	case core.ValMap:
		m := v.Data.(map[string]core.Value)
		res := make(map[string]interface{}, len(m))
		for k, item := range m {
			res[k] = valueToInterface(item)
		}
		return res
	default:
		return nil
	}
}

func parseJSON(input string) (core.Value, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return core.UnknownValue, err
	}
	return jsonToValue(raw), nil
}

func jsonToValue(v interface{}) core.Value {
	if v == nil {
		return core.UnknownValue
	}
	switch val := v.(type) {
	case float64:
		return core.NumberVal(val)
	case string:
		return core.StringVal(val)
	case bool:
		return core.BoolVal(val)
	case []interface{}:
		arr := make([]core.Value, len(val))
		for i, item := range val {
			arr[i] = jsonToValue(item)
		}
		return core.ArrayVal(arr)
	case map[string]interface{}:
		m := make(map[string]core.Value, len(val))
		for k, item := range val {
			m[k] = jsonToValue(item)
		}
		return core.MapVal(m)
	}
	return core.UnknownValue
}
