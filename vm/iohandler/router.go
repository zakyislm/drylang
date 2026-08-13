package iohandler

import (
	"drylang/core"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	routers    = make(map[int]*http.ServeMux)
	routerMws  = make(map[int][]*core.Closure)
	routerMut  sync.Mutex
	nextRtID   = 1
)

// BuiltinRt handles rt.new, rt.on, rt.serve, rt.use
func BuiltinRt(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("want rt(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "new":
		routerMut.Lock()
		id := nextRtID
		nextRtID++
		routers[id] = http.NewServeMux()
		routerMws[id] = make([]*core.Closure, 0)
		routerMut.Unlock()
		return core.NumberVal(float64(id)), nil

	case "on":
		if len(args) < 4 {
			return core.UnknownValue, vm.Errorf("rt.on wants (id, path, fn)")
		}
		if args[1].Type != core.ValNumber {
			return core.UnknownValue, vm.Errorf("rt.on wants router id")
		}
		id := int(args[1].Data.(float64))
		
		if args[2].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("rt.on wants path string")
		}
		path := args[2].Data.(string)

		if args[3].Type != core.ValFn {
			return core.UnknownValue, vm.Errorf("rt.on wants function")
		}
		fnClosure := args[3].Data.(*core.Closure)

		routerMut.Lock()
		mux, exists := routers[id]
		mws := routerMws[id]
		routerMut.Unlock()
		if !exists {
			return core.UnknownValue, vm.Errorf("router not found")
		}

		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			// CORS Headers (Default allow all for API development)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			reqMap := make(map[string]core.Value)
			reqMap["method"] = core.StringVal(r.Method)
			reqMap["url"] = core.StringVal(r.URL.Path)

			queryMap := make(map[string]core.Value)
			for k, v := range r.URL.Query() {
				if len(v) > 0 {
					queryMap[k] = core.StringVal(v[0])
				}
			}
			reqMap["query"] = core.MapVal(queryMap)

			if r.Body != nil {
				contentType := r.Header.Get("Content-Type")
				if len(contentType) >= 19 && contentType[:19] == "multipart/form-data" {
					r.ParseMultipartForm(10 << 20) // 10 MB limit
					
					formMap := make(map[string]core.Value)
					for k, v := range r.MultipartForm.Value {
						if len(v) > 0 {
							formMap[k] = core.StringVal(v[0])
						}
					}
					reqMap["form"] = core.MapVal(formMap)
					
					fileMap := make(map[string]core.Value)
					for k, v := range r.MultipartForm.File {
						if len(v) > 0 {
							// For files, we could expose filename and size
							fileInfo := make(map[string]core.Value)
							fileInfo["filename"] = core.StringVal(v[0].Filename)
							fileInfo["size"] = core.NumberVal(float64(v[0].Size))
							fileMap[k] = core.MapVal(fileInfo)
						}
					}
					reqMap["files"] = core.MapVal(fileMap)
				} else {
					bodyBytes, _ := io.ReadAll(r.Body)
					reqMap["body"] = core.StringVal(string(bodyBytes))
				}
			}

			// Apply Middlewares
			for _, mw := range mws {
				mwVM := vm.Clone()
				mwVM.Push(core.Value{Type: core.ValFn, Data: mw})
				mwVM.Push(core.MapVal(reqMap))
				if err := mwVM.CallFunction(mw.Fn, 1); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(fmt.Sprintf("Middleware Error: %v", err)))
					return
				}
				resVal := mwVM.Pop()
				// If middleware returns a map, we abort and send response
				if resVal.Type == core.ValMap {
					sendResponse(w, resVal.Data.(map[string]core.Value))
					return
				}
			}

			execVM := vm.Clone()
			execVM.Push(core.Value{Type: core.ValFn, Data: fnClosure})
			execVM.Push(core.MapVal(reqMap))
			err := execVM.CallFunction(fnClosure.Fn, 1)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Internal Server Error: %v", err)))
				return
			}
			
			resVal := execVM.Pop()
			if resVal.Type != core.ValMap {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Handler must return a map"))
				return
			}
			
			sendResponse(w, resVal.Data.(map[string]core.Value))
		})

		return core.UnknownValue, nil

	case "serve":
		if len(args) < 3 {
			return core.UnknownValue, vm.Errorf("rt.serve wants (id, addr)")
		}
		id := int(args[1].Data.(float64))
		addr := args[2].Data.(string)

		routerMut.Lock()
		mux, exists := routers[id]
		routerMut.Unlock()
		if !exists {
			return core.UnknownValue, vm.Errorf("router not found")
		}

		fmt.Printf("Server listening on %s\n", addr)
		err := http.ListenAndServe(addr, mux)
		if err != nil {
			return core.UnknownValue, vm.Errorf("server error: %v", err)
		}
		return core.UnknownValue, nil

	case "use":
		if len(args) < 3 {
			return core.UnknownValue, vm.Errorf("rt.use wants (id, fn)")
		}
		if args[1].Type != core.ValNumber {
			return core.UnknownValue, vm.Errorf("rt.use wants router id")
		}
		id := int(args[1].Data.(float64))
		
		if args[2].Type != core.ValFn {
			return core.UnknownValue, vm.Errorf("rt.use wants function")
		}
		mwClosure := args[2].Data.(*core.Closure)

		routerMut.Lock()
		defer routerMut.Unlock()
		if _, exists := routers[id]; !exists {
			return core.UnknownValue, vm.Errorf("router not found")
		}
		routerMws[id] = append(routerMws[id], mwClosure)
		return core.UnknownValue, nil

	default:
		return core.UnknownValue, vm.Errorf("unknown rt method: %s", method)
	}
}

func sendResponse(w http.ResponseWriter, resMap map[string]core.Value) {
	status := http.StatusOK
	if s, ok := resMap["status"]; ok && s.Type == core.ValNumber {
		status = int(s.Data.(float64))
	}
	
	w.Header().Set("Content-Type", "application/json") // Default
	if h, ok := resMap["headers"]; ok && h.Type == core.ValMap {
		headers := h.Data.(map[string]core.Value)
		for hk, hv := range headers {
			if hv.Type == core.ValString {
				w.Header().Set(hk, hv.Data.(string))
			}
		}
	}
	
	w.WriteHeader(status)
	
	if b, ok := resMap["body"]; ok && b.Type == core.ValString {
		w.Write([]byte(b.Data.(string)))
	} else if j, ok := resMap["json"]; ok {
		jsonBytes, _ := json.Marshal(convertValue(j))
		w.Write(jsonBytes)
	}
}



func convertValue(v core.Value) interface{} {
	switch v.Type {
	case core.ValString:
		return v.Data.(string)
	case core.ValNumber:
		return v.Data.(float64)
	case core.ValBool:
		return v.Data.(bool)
	case core.ValArray:
		arr := v.Data.([]core.Value)
		res := make([]interface{}, len(arr))
		for i, val := range arr {
			res[i] = convertValue(val)
		}
		return res
	case core.ValMap:
		m := v.Data.(map[string]core.Value)
		res := make(map[string]interface{})
		for k, val := range m {
			res[k] = convertValue(val)
		}
		return res
	default:
		return nil
	}
}
