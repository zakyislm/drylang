package oshandler

import (
	"drylang/core"
	"os"
)

func BuiltinArg(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	osArgs := os.Args
	// Skip binary name and script name
	startIdx := 2
	if len(osArgs) > startIdx {
		arr := make([]core.Value, len(osArgs)-startIdx)
		for i := startIdx; i < len(osArgs); i++ {
			arr[i-startIdx] = core.StringVal(osArgs[i])
		}
		result = core.ArrayVal(arr)
	} else {
		result = core.ArrayVal([]core.Value{})
	}
	return result, nil
}
