package oshandler

import (
	"drylang/core"
	"fmt"
	"os"
)

func BuiltinDie(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) > 0 {
		msg := args[0].String()
		if msg != "" {
			fmt.Fprintln(os.Stderr, msg)
		}
	}
	os.Exit(1)
	result = core.UnknownValue // unreachable
	return result, nil
}
