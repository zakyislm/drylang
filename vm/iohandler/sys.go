package iohandler

import (
	"bytes"
	"drylang/core"
	"os"
	"os/exec"
	"runtime"
)

// BuiltinSys handles sys.run, sys.env, sys.os
func BuiltinSys(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("want sys(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "run":
		if len(args) < 2 {
			return core.UnknownValue, vm.Errorf("sys.run wants (cmd, [args...])")
		}
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("sys.run wants command string")
		}
		cmdName := args[1].Data.(string)

		var cmdArgs []string
		if len(args) > 2 {
			if args[2].Type == core.ValArray {
				arr := args[2].Data.([]core.Value)
				for _, a := range arr {
					cmdArgs = append(cmdArgs, a.String())
				}
			} else {
				for i := 2; i < len(args); i++ {
					cmdArgs = append(cmdArgs, args[i].String())
				}
			}
		}

		cmd := exec.Command(cmdName, cmdArgs...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		err := cmd.Run()
		
		res := make(map[string]core.Value)
		res["stdout"] = core.StringVal(out.String())
		res["stderr"] = core.StringVal(stderr.String())
		
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				res["exitCode"] = core.NumberVal(float64(exitError.ExitCode()))
			} else {
				res["exitCode"] = core.NumberVal(-1)
				res["error"] = core.StringVal(err.Error())
			}
		} else {
			res["exitCode"] = core.NumberVal(0)
		}

		return core.MapVal(res), nil

	case "dotenv":
		filename := ".env"
		if len(args) > 1 && args[1].Type == core.ValString {
			filename = args[1].Data.(string)
		}
		
		data, err := os.ReadFile(filename)
		if err != nil {
			return core.UnknownValue, vm.Errorf("sys.dotenv error: %v", err)
		}
		
		lines := bytes.Split(data, []byte("\n"))
		for _, line := range lines {
			line = bytes.TrimSpace(line)
			if len(line) == 0 || line[0] == '#' {
				continue
			}
			parts := bytes.SplitN(line, []byte("="), 2)
			if len(parts) == 2 {
				key := string(bytes.TrimSpace(parts[0]))
				val := string(bytes.TrimSpace(parts[1]))
				// Remove quotes if present
				if len(val) >= 2 && ((val[0] == '"' && val[len(val)-1] == '"') || (val[0] == '\'' && val[len(val)-1] == '\'')) {
					val = val[1 : len(val)-1]
				}
				os.Setenv(key, val)
			}
		}
		return core.BoolVal(true), nil

	case "env":
		if len(args) == 1 {
			// Return all env vars as map
			envMap := make(map[string]core.Value)
			for _, e := range os.Environ() {
				for i, b := range e {
					if b == '=' {
						envMap[e[:i]] = core.StringVal(e[i+1:])
						break
					}
				}
			}
			return core.MapVal(envMap), nil
		}
		
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("sys.env wants var name")
		}
		key := args[1].Data.(string)
		
		if len(args) == 2 {
			// Get
			val, exists := os.LookupEnv(key)
			if !exists {
				return core.UnknownValue, nil
			}
			return core.StringVal(val), nil
		}
		
		// Set
		if args[2].Type == core.ValString {
			os.Setenv(key, args[2].Data.(string))
		}
		return core.UnknownValue, nil

	case "os":
		info := make(map[string]core.Value)
		info["name"] = core.StringVal(runtime.GOOS)
		info["arch"] = core.StringVal(runtime.GOARCH)
		info["cpus"] = core.NumberVal(float64(runtime.NumCPU()))
		return core.MapVal(info), nil

	default:
		return core.UnknownValue, vm.Errorf("unknown sys method: %s", method)
	}
}
