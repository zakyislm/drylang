package core

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Value types
const (
	ValNumber         = "number"
	ValString         = "string"
	ValBool           = "bool"
	ValArray          = "array"
	ValMap            = "map"
	ValFn             = "fn"
	ValBoundMethod    = "bound_method"
	ValStructDef      = "struct_def"
	ValStructInstance = "struct_instance"
	ValClass          = "class"
	ValInstance       = "instance"
	ValUnknown        = "unknown"
)

// Value wraps any dryLang runtime value.
type Value struct {
	Type string
	Data interface{}
}

var UnknownValue = Value{Type: ValUnknown, Data: nil}

func NumberVal(v float64) Value       { return Value{ValNumber, v} }
func StringVal(v string) Value        { return Value{ValString, v} }
func BoolVal(v bool) Value            { return Value{ValBool, v} }
func ArrayVal(v []Value) Value        { return Value{ValArray, v} }
func MapVal(v map[string]Value) Value { return Value{ValMap, v} }
type Closure struct {
	Fn  *CompiledFn
	Env map[string]Value
}

func FnVal(v *Closure) Value       { return Value{ValFn, v} }

func (v Value) String() string {
	switch v.Type {
	case ValNumber:
		f := v.Data.(float64)
		if f == math.Trunc(f) {
			return strconv.FormatInt(int64(f), 10)
		}
		return strconv.FormatFloat(f, 'f', -1, 64)
	case ValString:
		return v.Data.(string)
	case ValBool:
		if v.Data.(bool) {
			return "t"
		}
		return "f"
	case ValArray:
		arr := v.Data.([]Value)
		parts := make([]string, len(arr))
		for i, item := range arr {
			parts[i] = item.String()
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case ValMap:
		m := v.Data.(map[string]Value)
		parts := make([]string, 0, len(m))
		for k, val := range m {
			parts = append(parts, fmt.Sprintf("%s: %s", k, val.String()))
		}
		return "{" + strings.Join(parts, ", ") + "}"
	case ValFn:
		c := v.Data.(*Closure)
		return fmt.Sprintf("<fn %s>", c.Fn.Name)
	case ValStructDef:
		sd := v.Data.(StructDef)
		return fmt.Sprintf("<struct %s>", sd.Name)
	case ValStructInstance:
		m := v.Data.(map[string]Value)
		parts := make([]string, 0, len(m))
		for k, val := range m {
			parts = append(parts, fmt.Sprintf("%s: %s", k, val.String()))
		}
		return "{" + strings.Join(parts, ", ") + "}"
	case ValClass:
		cd := v.Data.(ClassDef)
		return fmt.Sprintf("<class %s>", cd.Name)
	case ValInstance:
		inst := v.Data.(*Instance)
		return fmt.Sprintf("<instance %s>", inst.Class.Name)
	case ValUnknown:
		return "unknown"
	}
	return "?"
}

func IsTruthy(v Value) bool {
	switch v.Type {
	case ValBool:
		return v.Data.(bool)
	case ValNumber:
		return v.Data.(float64) != 0
	case ValString:
		return v.Data.(string) != ""
	case ValArray:
		return len(v.Data.([]Value)) > 0
	case ValMap:
		return len(v.Data.(map[string]Value)) > 0
	case ValUnknown:
		return false
	}
	return true
}

type Instance struct {
	Class  ClassDef
	Fields map[string]Value
}

type BoundMethod struct {
	Instance *Instance
	Method   ClassMethod
}
