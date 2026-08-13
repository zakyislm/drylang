package vm

import (
	"bufio"
	"context"
	"drylang/core"
	"drylang/handler/system"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// builtinIn reads a line from standard input, printing an optional prompt.
func (vm *VM) builtinIn(args []core.Value, line, col int) (core.Value, error) {
	if len(args) > 0 {
		fmt.Print(args[0].String())
	}
	if vm.stdinScanner == nil {
		vm.stdinScanner = bufio.NewScanner(os.Stdin)
	}
	if vm.stdinScanner.Scan() {
		text := strings.TrimSpace(vm.stdinScanner.Text())
		return core.StringVal(text), nil
	}
	return core.StringVal(""), nil
}

// builtinOp starts an HTTP server. mode "uni" runs handlers on the shared VM;
// mode "mul" clones the VM per request with a worker semaphore.
func (vm *VM) builtinOp(args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 2 {
		return core.UnknownValue, vm.runtimeErr("E300", line, col, "want port, handler")
	}
	port := args[0].String()
	if args[1].Type != core.ValFn {
		return core.UnknownValue, vm.runtimeErr("E300", line, col, "want handler function")
	}
	closure := args[1].Data.(*core.Closure)
	handlerFn := closure.Fn

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
		system.Requests.Add(1)
		var execVM *VM

		if mode == "mul" {
			sem <- struct{}{}
			defer func() { <-sem }()

			execVM = &VM{
				chunk:   vm.chunk,
				fns:     vm.fns,
				globals: vm.CloneGlobals(),
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
		reqMap["query"] = core.Value{Type: core.ValMap, Data: queryMap}

		// Save old state if we are in uni mode
		savedIP := execVM.ip
		savedChunk := execVM.chunk

		execVM.push(core.FnVal(&core.Closure{Fn: handlerFn, Env: make(map[string]core.Value)}))
		execVM.push(core.Value{Type: core.ValMap, Data: reqMap})

		frame := callFrame{
			closure: &core.Closure{Fn: handlerFn, Env: make(map[string]core.Value)},
			ip:    0,
			bp:    execVM.sp - 1, // args
			chunk: handlerFn.Chunk,
		}
		execVM.frames = append(execVM.frames, frame)

		for execVM.sp < frame.bp+handlerFn.Chunk.LocalsCount+1 {
			execVM.push(core.UnknownValue)
		}

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
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           http.DefaultServeMux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1 MB
	}

	// Graceful shutdown on SIGINT/SIGTERM.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stop
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return core.UnknownValue, vm.runtimeErr("E300", line, col, "server error: %v", err)
	}
	return core.BoolVal(true), nil
}
