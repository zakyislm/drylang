package core

import (
	"drylang/core"
	"math/rand"
)

func BuiltinRan(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	result = core.NumberVal(rand.Float64())
	return result, nil
}
