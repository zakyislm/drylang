package oshandler

import (
	"drylang/core"
	"time"
)

func BuiltinNow(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	result = core.NumberVal(float64(time.Now().UnixMilli()))
	return result, nil
}
