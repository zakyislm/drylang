package core

import (
	"bufio"
	"drylang/core"
	"os"
	"strings"
)

func BuiltinIn(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	result = core.StringVal(strings.TrimSpace(text))
	return result, nil
}
