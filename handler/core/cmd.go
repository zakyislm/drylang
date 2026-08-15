package core

import (
	"drylang/core"
	"os/exec"
	"strings"
)

func BuiltinCmd(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want command string", line, col)
	}
	cmdStr := args[0].Data.(string)

	if !core.GetSandbox().AllowCmd(cmdStr) {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: cmd denied: %s", line, col, cmdStr)
	}

	cmdArgs := make([]string, 0)
	for i := 1; i < len(args); i++ {
		cmdArgs = append(cmdArgs, args[i].String())
	}
	var cmd *exec.Cmd
	if len(cmdArgs) > 0 {
		cmd = exec.Command(cmdStr, cmdArgs...)
	} else {
		cmd = exec.Command(cmdStr)
	}
	out, err := cmd.Output()
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"cmd fail: %v\n%s", line, col, err, string(out))
	}
	result = core.StringVal(strings.TrimRight(string(out), "\n\r"))
	return result, nil
}
