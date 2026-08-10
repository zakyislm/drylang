package vm

import (
	"drylang/core"
	"fmt"
	"io"
	"net/http"
)

func BuiltinOp(vm *VM, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) < 2 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want port, handler", line, col)
	}
	port := args[0].String()
	if args[1].Type != core.ValFn {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want handler function", line, col)
	}
	handlerFn := args[1].Data.(*core.CompiledFn)

	mode := "uni"
	if len(args) > 2 {
		mode = args[2].String()
	}

	maxWorkers := 100
	if len(args) > 3 && args[3].Type == core.ValNumber {
		maxWorkers = int(args[3].Data.(float64))
	}

	sem := make(chan struct{}, maxWorkers)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var execVM *VM

		if mode == "mul" {
			sem <- struct{}{}
			defer func() { <-sem }()

			execVM = &VM{
				chunk:   vm.chunk,
				fns:     vm.fns,
				globals: vm.globals, // shared!
				stack:   make([]core.Value, 1024),
				sp:      0,
				ip:      0,
			}
		} else {
			vm.mu.Lock()
			defer vm.mu.Unlock()
			execVM = vm
		}

		reqMap := make(map[string]core.Value)
		method := r.Method
		switch method {
		case "GET":
			method = "G"
		case "POST":
			method = "PO"
		case "PUT":
			method = "PUT"
		case "PATCH":
			method = "PAT"
		case "DELETE":
			method = "DEL"
		case "OPTIONS":
			method = "OPT"
		case "HEAD":
			method = "H"
		}
		reqMap["method"] = core.StringVal(method)
		reqMap["path"] = core.StringVal(r.URL.Path)

		bodyBytes, _ := io.ReadAll(r.Body)
		reqMap["body"] = core.StringVal(string(bodyBytes))

		queryMap := make(map[string]core.Value)
		for k, v := range r.URL.Query() {
			if len(v) > 0 {
				queryMap[k] = core.StringVal(v[0])
			}
		}
		reqMap["query"] = core.Value{core.ValMap, queryMap}

		// Save old state if we are in uni mode
		savedIP := execVM.ip
		savedChunk := execVM.chunk

		execVM.push(core.FnVal(handlerFn))
		execVM.push(core.Value{core.ValMap, reqMap})

		frame := callFrame{
			fn:    handlerFn,
			ip:    execVM.ip,
			bp:    execVM.sp - 1, // 1 argument
			chunk: execVM.chunk,
		}
		execVM.frames = append(execVM.frames, frame)

		err := execVM.execute(handlerFn.Chunk)

		var res core.Value
		if err == nil {
			execVM.frames = execVM.frames[:len(execVM.frames)-1]
			res = execVM.pop()
			execVM.sp = frame.bp - 1
			if execVM.sp < 0 {
				execVM.sp = 0
			}
			execVM.ip = savedIP
			execVM.chunk = savedChunk
		} else {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		if res.Type == core.ValMap {
			m := res.Data.(map[string]core.Value)
			if status, ok := m["status"]; ok && status.Type == core.ValNumber {
				w.WriteHeader(int(status.Data.(float64)))
			}
			if body, ok := m["body"]; ok {
				w.Write([]byte(body.String()))
			}
		} else {
			w.Write([]byte(res.String()))
		}
	})

	fmt.Printf("Starting dryLang server on port %s (mode: %s)...\n", port, mode)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"server error: %v", err, line, col)
	}
	result = core.BoolVal(true)
	return result, nil
}
