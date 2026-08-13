package system

import (
	"drylang/core"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// BuiltinLog handles log("inf"|"obv"|"err", ...args). With DRY_LOG_JSON=1
// output is one JSON object per line; otherwise plain text with timestamp.
func BuiltinLog(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: want log level", line, col)
	}

	level := args[0].String()
	ts := time.Now().Format("2006/01/02 15:04:05")

	var output string
	for i := 1; i < len(args); i++ {
		output += args[i].String()
		if i < len(args)-1 {
			output += " "
		}
	}

	jsonMode := os.Getenv("DRY_LOG_JSON") == "1"
	if jsonMode {
		rec := map[string]string{"ts": ts, "level": level, "msg": output}
		b, _ := json.Marshal(rec)
		fmt.Println(string(b))
		return core.BoolVal(true), nil
	}

	switch level {
	case "inf":
		fmt.Printf("%s [INF] %s\n", ts, output)
	case "obv":
		fmt.Printf("%s [OBV] %s\n", ts, output)
	case "err":
		fmt.Printf("%s [ERR] %s\n", ts, output)
	default:
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: unknown log level %s", line, col, level)
	}

	return core.BoolVal(true), nil
}
