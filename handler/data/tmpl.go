package data

import (
	"bytes"
	"drylang/core"
	"html/template"
)

func tmplConvertMap(v core.Value) interface{} {
	if v.Type == core.ValMap {
		m := v.Data.(map[string]core.Value)
		out := make(map[string]interface{})
		for k, val := range m {
			out[k] = tmplConvertValue(val)
		}
		return out
	}
	return tmplConvertValue(v)
}

func tmplConvertValue(v core.Value) interface{} {
	switch v.Type {
	case core.ValString:
		return v.Data.(string)
	case core.ValNumber:
		return v.Data.(float64)
	case core.ValBool:
		return v.Data.(bool)
	case core.ValArray:
		arr := v.Data.([]core.Value)
		var out []interface{}
		for _, item := range arr {
			out = append(out, tmplConvertValue(item))
		}
		return out
	case core.ValMap:
		return tmplConvertMap(v)
	}
	return nil
}

// BuiltinTmpl handles tmpl("render", tmplStr, dataMap)
func BuiltinTmpl(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want tmpl(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "render":
		if len(args) != 3 {
			return core.UnknownValue, v.Errorf("tmpl.render wants (tmplString, dataObject)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValMap {
			return core.UnknownValue, v.Errorf("tmpl.render args must be (string, map)")
		}

		tmplStr := args[1].Data.(string)
		dataObj := tmplConvertMap(args[2])

		t, err := template.New("tpl").Parse(tmplStr)
		if err != nil {
			return core.UnknownValue, v.Errorf("tmpl parse error: %s", err)
		}

		var buf bytes.Buffer
		err = t.Execute(&buf, dataObj)
		if err != nil {
			return core.UnknownValue, v.Errorf("tmpl execute error: %s", err)
		}

		return core.Value{Type: core.ValString, Data: buf.String()}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown tmpl method: %s", method)
	}
}
