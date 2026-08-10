package iohandler

import (
	"drylang/core"
	"encoding/json"
)

func BuiltinJson(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) != 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want json string", line, col)
	}
	// Simple JSON parser: converts JSON string to dryLang map/array
	jsonStr := args[0].Data.(string)
	parsed, err := parseJSON(jsonStr)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"json parse fail: %v", err, line, col)
	}
	result = parsed
	return result, nil
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
