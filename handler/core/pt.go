package core

import (
	"drylang/core"
	"fmt"
)

func BuiltinPt(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	for i, arg := range args {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arg.String())
	}
	fmt.Println()
	return core.BoolVal(true), nil
}
