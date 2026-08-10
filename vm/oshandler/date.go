package oshandler

import (
	"drylang/core"
	"time"
)

func BuiltinDate(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	now := time.Now()
	m := make(map[string]core.Value)
	m["year"] = core.NumberVal(float64(now.Year()))
	m["month"] = core.NumberVal(float64(now.Month()))
	m["day"] = core.NumberVal(float64(now.Day()))
	m["hour"] = core.NumberVal(float64(now.Hour()))
	m["min"] = core.NumberVal(float64(now.Minute()))
	m["sec"] = core.NumberVal(float64(now.Second()))
	m["format"] = core.StringVal(now.Format("2006-01-02 15:04:05"))
	result = core.Value{core.ValMap, m}
	return result, nil
}
