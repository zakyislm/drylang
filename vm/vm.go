package vm

import (
	"bufio"
	"drylang/core"
	"drylang/errfmt"
	"drylang/vm/classhandler"
	"drylang/vm/colshandler"
	"drylang/vm/controlhandler"
	"drylang/vm/errorhandler"
	"drylang/vm/functionhandler"
	"drylang/vm/mathhandler"
	"drylang/vm/varhandler"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"database/sql"
)

// VM executes dryLang bytecode.
type VM struct {
	mu           sync.Mutex
	chunk        *core.Chunk
	fns          []*core.CompiledFn
	globals      map[string]core.Value
	stack        []core.Value
	sp           int // stack pointer
	ip           int // instruction pointer
	frames       []callFrame
	tryStack     []tryFrame
	stdinScanner *bufio.Scanner
}

type callFrame struct {
	fn    *core.CompiledFn
	ip    int
	bp    int // base pointer (stack offset)
	chunk *core.Chunk
}

type tryFrame struct {
	catchIP    int
	sp         int
	errVarName string
	frameDepth int
}

type Instance struct {
	Class  compiler.ClassDef
	Fields map[string]Value
}

type BoundMethod struct {
	Instance *Instance
	Method   compiler.ClassMethod
}

// New creates a new VM.
func New(chunk *core.Chunk, fns []*core.CompiledFn) *VM {
	return &VM{
		chunk:   chunk,
		fns:     fns,
		globals: make(map[string]core.Value),
		stack:   make([]core.Value, 4096),
		sp:      0,
		ip:      0,
	}
}

func (vm *VM) push(v core.Value) {
	vm.stack[vm.sp] = v
	vm.sp++
}

func (vm *VM) pop() core.Value {
	vm.sp--
	return vm.stack[vm.sp]
}

func (vm *VM) Update(chunk *core.Chunk, fns []*core.CompiledFn) {
	vm.chunk = chunk
	vm.fns = fns
}

func (vm *VM) peek() core.Value {
	return vm.stack[vm.sp-1]
}

func (vm *VM) runtimeErr(code string, line, col int, format string, args ...interface{}) error {
	return errfmt.Format(code, line, col, fmt.Sprintf(format, args...))
}

// Run executes the main chunk.
func (vm *VM) Run() error {
	return vm.execute(vm.chunk)
}

func (vm *VM) execute(chunk *core.Chunk) error {
	vm.chunk = chunk
	vm.ip = 0

	for vm.ip < len(chunk.Code) {
		inst := chunk.Code[vm.ip]
		line := 0
		col := 0
		if vm.ip < len(chunk.Lines) {
			line = chunk.Lines[vm.ip]
			col = chunk.Cols[vm.ip]
		}
		vm.ip++

		switch inst.Op {
		case core.OpConst:
			val := chunk.Constants[inst.Operand]
			switch v := val.(type) {
			case float64:
				vm.push(core.NumberVal(v))
			case string:
				vm.push(core.StringVal(v))
			case *core.CompiledFn:
				vm.push(core.FnVal(v))
			case core.StructDef:
				vm.push(core.Value{core.ValStructDef, v})
			case core.ClassDef:
				// Resolve parents at runtime
				var resolvedParents []*core.ClassDef
				for _, pName := range v.ParentNames {
					if parentVal, ok := vm.globals[pName]; ok && parentVal.Type == core.ValClass {
						pDef := parentVal.Data.(core.ClassDef)
						resolvedParents = append(resolvedParents, &pDef)
					} else {
						return vm.Errorf("E300 at %d:%d: parent class '%s' not found", line, col, pName)
					}
				}
				v.Parents = resolvedParents
				vm.push(core.Value{core.ValClass, v})
			default:
				vm.push(core.Value{core.ValUnknown, nil})
			}

		case core.OpPop:
			if vm.sp > 0 {
				vm.pop()
			}

		case core.OpTrue:
			vm.push(core.BoolVal(true))
		case core.OpFalse:
			vm.push(core.BoolVal(false))
		case core.OpUnknown:
			vm.push(core.UnknownValue)

		case core.OpAdd:
			if err := mathhandler.OpAdd(vm, line, col); err != nil {
				return err
			}

		case core.OpSub:
			if err := mathhandler.OpSub(vm, line, col); err != nil {
				return err
			}

		case core.OpMul:
			if err := mathhandler.OpMul(vm, line, col); err != nil {
				return err
			}

		case core.OpDiv:
			if err := mathhandler.OpDiv(vm, line, col); err != nil {
				return err
			}

		case core.OpMod:
			if err := mathhandler.OpMod(vm, line, col); err != nil {
				return err
			}

		case core.OpNeg:
			v := vm.pop()
			if v.Type != core.ValNumber {
				return vm.runtimeErr("E300", line, col, "want number")
			}
			vm.push(core.NumberVal(-v.Data.(float64)))

		case core.OpEqual:
			if err := mathhandler.OpEqual(vm, line, col); err != nil {
				return err
			}

		case core.OpNotEqual:
			if err := mathhandler.OpNotEqual(vm, line, col); err != nil {
				return err
			}

		case core.OpLess:
			if err := mathhandler.OpLess(vm, line, col); err != nil {
				return err
			}

		case core.OpGreater:
			if err := mathhandler.OpGreater(vm, line, col); err != nil {
				return err
			}

		case core.OpLessEq:
			if err := mathhandler.OpLessEq(vm, line, col); err != nil {
				return err
			}

		case core.OpGreaterEq:
			if err := mathhandler.OpGreaterEq(vm, line, col); err != nil {
				return err
			}

		case core.OpAnd:
			if err := mathhandler.OpAnd(vm, line, col); err != nil {
				return err
			}

		case core.OpOr:
			if err := mathhandler.OpOr(vm, line, col); err != nil {
				return err
			}

		case core.OpNot:
			v := vm.pop()
			vm.push(core.BoolVal(!core.IsTruthy(v)))

		case core.OpConcat:
			if err := mathhandler.OpConcat(vm, line, col); err != nil {
				return err
			}

		case core.OpGetGlobal:
			name := chunk.Constants[inst.Operand].(string)
			val, ok := vm.globals[name]
			if !ok {
				return vm.runtimeErr("E300", line, col, "unknown %s", name)
			}
			vm.push(val)

		case core.OpSetGlobal:
			name := chunk.Constants[inst.Operand].(string)
			if err := varhandler.OpSetGlobal(vm, name, line, col); err != nil {
				return err
			}

		case core.OpGetLocal:
			if err := varhandler.OpGetLocal(vm, inst.Operand, line, col); err != nil {
				return err
			}

		case core.OpSetLocal:
			if err := varhandler.OpSetLocal(vm, inst.Operand, line, col); err != nil {
				return err
			}

		case core.OpJump:
			if err := controlhandler.OpJump(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpJumpIfFalse:
			if err := controlhandler.OpJumpIfFalse(vm, inst.Operand); err != nil {
				return err
			}
		
		case core.OpJumpIfNotUnknown:
			if err := controlhandler.OpJumpIfNotUnknown(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpLoop:
			if err := controlhandler.OpLoop(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpArray:
			if err := colshandler.OpArray(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpMap:
			if err := colshandler.OpMap(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpIndex:
			if err := colshandler.OpIndex(vm, line, col); err != nil {
				return err
			}

		case core.OpSetIndex:
			if err := colshandler.OpSetIndex(vm, line, col); err != nil {
				return err
			}

		case core.OpDotGet:
			field := chunk.Constants[inst.Operand].(string)
			if err := classhandler.OpDotGet(vm, line, col, field, false); err != nil {
				return err
			}
			
		case core.OpOptDotGet:
			field := chunk.Constants[inst.Operand].(string)
			if err := classhandler.OpDotGet(vm, line, col, field, true); err != nil {
				return err
			}

		case core.OpDotSet:
			field := chunk.Constants[inst.Operand].(string)
			if err := classhandler.OpDotSet(vm, line, col, field); err != nil {
				return err
			}
		case core.OpClosure:
			fn := chunk.Constants[inst.Operand].(*core.CompiledFn)
			if err := functionhandler.OpClosure(vm, fn); err != nil {
				return err
			}

		case core.OpCall:
			if err := functionhandler.OpCall(vm, line, col, inst.Operand); err != nil {
				return err
			}

		case core.OpReturn:
			if len(vm.frames) > 0 {
				return nil // return from function execution
			}
			// Top-level rev = exit
			return nil

		case core.OpTry:
			if err := errorhandler.OpTry(vm, inst.Operand); err != nil {
				return err
			}

		case core.OpEndTry:
			if err := errorhandler.OpEndTry(vm); err != nil {
				return err
			}

		case core.OpThrow:
			if err := errorhandler.OpThrow(vm); err != nil {
				return err
			}

		case core.OpBuiltin:
			if err := vm.executeBuiltin(core.BuiltinID(inst.Operand), inst.Operand2, line, col); err != nil {
				return err
			}

		case core.OpGetLoopCounter:
			// handled via locals

		case core.OpAsync:
			// simplified: just execute normally for now

		case core.OpAwait:
			// simplified: value is already on stack
		}
	}

	return nil
}

func (vm *VM) executeBuiltin(id core.BuiltinID, argCount int, line, col int) error {
	args := make([]core.Value, argCount)
	for i := argCount - 1; i >= 0; i-- {
		args[i] = vm.pop()
	}

	var result core.Value

	switch id {
	case core.BuiltinLen:
		if len(args) != 1 {
			return vm.runtimeErr("E300", line, col, "want 1 arg")
		}
		switch args[0].Type {
		case core.ValString:
			result = core.NumberVal(float64(len(args[0].Data.(string))))
		case core.ValArray:
			result = core.NumberVal(float64(len(args[0].Data.([]core.Value))))
		case core.ValMap:
			result = core.NumberVal(float64(len(args[0].Data.(map[string]core.Value))))
		default:
			return vm.runtimeErr("E300", line, col, "want string|array|map")
		}

	case core.BuiltinGet:
		if len(args) != 1 {
			return vm.runtimeErr("E300", line, col, "want 1 arg")
		}
		result = core.StringVal(args[0].Type)

	case core.BuiltinAdd:
		if len(args) != 2 {
			return vm.runtimeErr("E300", line, col, "want 2 args")
		}
		if args[0].Type != core.ValArray {
			return vm.runtimeErr("E300", line, col, "want array")
		}
		arr := args[0].Data.([]core.Value)
		arr = append(arr, args[1])
		// Modify in place via globals would need ref semantics.
		// For now, push new array.
		result = core.ArrayVal(arr)

	case core.BuiltinNum:
		if len(args) != 1 {
			return vm.runtimeErr("E300", line, col, "want 1 arg")
		}
		f, err := strconv.ParseFloat(args[0].String(), 64)
		if err != nil {
			return vm.runtimeErr("E300", line, col, "bad number")
		}
		result = core.NumberVal(f)

	case core.BuiltinStr:
		if len(args) != 1 {
			return vm.runtimeErr("E300", line, col, "want 1 arg")
		}
		result = core.StringVal(args[0].String())

	case core.BuiltinAbs:
		if len(args) != 1 || args[0].Type != core.ValNumber {
			return vm.runtimeErr("E300", line, col, "want number")
		}
		result = core.NumberVal(math.Abs(args[0].Data.(float64)))

	case core.BuiltinMin:
		if len(args) != 2 || args[0].Type != core.ValNumber || args[1].Type != core.ValNumber {
			return vm.runtimeErr("E300", line, col, "want 2 numbers")
		}
		result = core.NumberVal(math.Min(args[0].Data.(float64), args[1].Data.(float64)))

	case core.BuiltinMax:
		if len(args) != 2 || args[0].Type != core.ValNumber || args[1].Type != core.ValNumber {
			return vm.runtimeErr("E300", line, col, "want 2 numbers")
		}
		result = core.NumberVal(math.Max(args[0].Data.(float64), args[1].Data.(float64)))

	case core.BuiltinRnd:
		if len(args) != 1 || args[0].Type != core.ValNumber {
			return vm.runtimeErr("E300", line, col, "want number")
		}
		result = core.NumberVal(math.Round(args[0].Data.(float64)))

	case core.BuiltinCap:
		if len(args) != 1 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want string")
		}
		result = core.StringVal(strings.ToUpper(args[0].Data.(string)))

	case core.BuiltinLow:
		if len(args) != 1 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want string")
		}
		result = core.StringVal(strings.ToLower(args[0].Data.(string)))

	case core.BuiltinTrm:
		if len(args) != 1 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want string")
		}
		result = core.StringVal(strings.TrimSpace(args[0].Data.(string)))

	case core.BuiltinSpl:
		if len(args) != 2 || args[0].Type != core.ValString || args[1].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want 2 strings")
		}
		parts := strings.Split(args[0].Data.(string), args[1].Data.(string))
		arr := make([]core.Value, len(parts))
		for i, p := range parts {
			arr[i] = core.StringVal(p)
		}
		result = core.ArrayVal(arr)

	case core.BuiltinJ:
		if len(args) != 2 || args[0].Type != core.ValArray || args[1].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want array, string")
		}
		arr := args[0].Data.([]core.Value)
		strs := make([]string, len(arr))
		for i, v := range arr {
			strs[i] = v.String()
		}
		result = core.StringVal(strings.Join(strs, args[1].Data.(string)))

	case core.BuiltinMod:
		if len(args) != 3 || args[0].Type != core.ValString || args[1].Type != core.ValString || args[2].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want 3 strings")
		}
		result = core.StringVal(strings.ReplaceAll(args[0].Data.(string), args[1].Data.(string), args[2].Data.(string)))

	case core.BuiltinHas:
		if len(args) != 2 {
			return vm.runtimeErr("E300", line, col, "want 2 args")
		}
		switch args[0].Type {
		case core.ValString:
			if args[1].Type != core.ValString {
				return vm.runtimeErr("E300", line, col, "want string arg")
			}
			result = core.BoolVal(strings.Contains(args[0].Data.(string), args[1].Data.(string)))
		case core.ValMap:
			if args[1].Type != core.ValString {
				return vm.runtimeErr("E300", line, col, "want string key")
			}
			m := args[0].Data.(map[string]core.Value)
			_, ok := m[args[1].Data.(string)]
			result = core.BoolVal(ok)
		case core.ValArray:
			arr := args[0].Data.([]core.Value)
			found := false
			for _, v := range arr {
				if v.String() == args[1].String() {
					found = true
					break
				}
			}
			result = core.BoolVal(found)
		default:
			return vm.runtimeErr("E300", line, col, "want str|map|arr")
		}

	case core.BuiltinSort:
		if len(args) != 1 || args[0].Type != core.ValArray {
			return vm.runtimeErr("E300", line, col, "want array")
		}
		arr := make([]core.Value, len(args[0].Data.([]core.Value)))
		copy(arr, args[0].Data.([]core.Value))
		sort.Slice(arr, func(i, j int) bool {
			if arr[i].Type == core.ValNumber && arr[j].Type == core.ValNumber {
				return arr[i].Data.(float64) < arr[j].Data.(float64)
			}
			return arr[i].String() < arr[j].String()
		})
		result = core.ArrayVal(arr)

	case core.BuiltinPop:
		if len(args) != 1 || args[0].Type != core.ValArray {
			return vm.runtimeErr("E300", line, col, "want array")
		}
		arr := args[0].Data.([]core.Value)
		if len(arr) == 0 {
			return vm.runtimeErr("E300", line, col, "empty array")
		}
		result = arr[len(arr)-1]

	case core.BuiltinRm:
		if len(args) != 2 || args[0].Type != core.ValArray || args[1].Type != core.ValNumber {
			return vm.runtimeErr("E300", line, col, "want array, number")
		}
		arr := args[0].Data.([]core.Value)
		idx := int(args[1].Data.(float64))
		if idx < 0 || idx >= len(arr) {
			return vm.runtimeErr("E300", line, col, "bounds %d", idx)
		}
		newArr := make([]core.Value, 0, len(arr)-1)
		newArr = append(newArr, arr[:idx]...)
		newArr = append(newArr, arr[idx+1:]...)
		result = core.ArrayVal(newArr)

	case core.BuiltinKey:
		if len(args) != 1 || args[0].Type != core.ValMap {
			return vm.runtimeErr("E300", line, col, "want map")
		}
		m := args[0].Data.(map[string]core.Value)
		keys := make([]core.Value, 0, len(m))
		for k := range m {
			keys = append(keys, core.StringVal(k))
		}
		result = core.ArrayVal(keys)

	case core.BuiltinVal:
		if len(args) != 1 || args[0].Type != core.ValMap {
			return vm.runtimeErr("E300", line, col, "want map")
		}
		m := args[0].Data.(map[string]core.Value)
		vals := make([]core.Value, 0, len(m))
		for _, v := range m {
			vals = append(vals, v)
		}
		result = core.ArrayVal(vals)

	case core.BuiltinRan:
		result = core.NumberVal(rand.Float64())

	case core.BuiltinQ:
		if len(args) != 1 || args[0].Type != core.ValNumber {
			return vm.runtimeErr("E300", line, col, "want number")
		}
		ms := int(args[0].Data.(float64))
		time.Sleep(time.Duration(ms) * time.Millisecond)
		result = core.UnknownValue

	case core.BuiltinR:
		if len(args) < 1 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want string")
		}
		data, err := os.ReadFile(args[0].Data.(string))
		if err != nil {
			return vm.runtimeErr("E300", line, col, "read fail")
		}
		result = core.StringVal(string(data))

	case core.BuiltinW:
		if len(args) != 2 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want path, data")
		}
		err := os.WriteFile(args[0].Data.(string), []byte(args[1].String()), 0644)
		if err != nil {
			return vm.runtimeErr("E300", line, col, "write fail")
		}
		result = core.BoolVal(true)

	case core.BuiltinOp:
		if len(args) < 2 {
			return vm.runtimeErr("E300", line, col, "want port, handler")
		}
		port := args[0].String()
		if args[1].Type != core.ValFn {
			return vm.runtimeErr("E300", line, col, "want handler function")
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
			return vm.runtimeErr("E300", line, col, "server error: %v", err)
		}
		result = core.BoolVal(true)

	case core.BuiltinDb:
		if len(args) < 3 {
			return vm.runtimeErr("E300", line, col, "want driver, dsn, query")
		}
		driver := args[0].String()
		dsn := args[1].String()
		query := args[2].String()

		db, err := sql.Open(driver, dsn)
		if err != nil {
			return vm.runtimeErr("E300", line, col, "db open err: %v", err)
		}
		defer db.Close()

		var qargs []interface{}
		for i := 3; i < len(args); i++ {
			if args[i].Type == core.ValNumber {
				qargs = append(qargs, args[i].Data.(float64))
			} else if args[i].Type == core.ValBool {
				qargs = append(qargs, args[i].Data.(bool))
			} else {
				qargs = append(qargs, args[i].String())
			}
		}

		isSelect := strings.HasPrefix(strings.ToUpper(strings.TrimSpace(query)), "SELECT")
		if isSelect || strings.HasPrefix(strings.ToUpper(strings.TrimSpace(query)), "PRAGMA") || strings.HasPrefix(strings.ToUpper(strings.TrimSpace(query)), "SHOW") || strings.HasPrefix(strings.ToUpper(strings.TrimSpace(query)), "DESCRIBE") {
			rows, err := db.Query(query, qargs...)
			if err != nil {
				return vm.runtimeErr("E300", line, col, "db query err: %v", err)
			}
			defer rows.Close()

			cols, _ := rows.Columns()
			var resultArr []core.Value

			for rows.Next() {
				columns := make([]interface{}, len(cols))
				columnPointers := make([]interface{}, len(cols))
				for i := range columns {
					columnPointers[i] = &columns[i]
				}

				if err := rows.Scan(columnPointers...); err != nil {
					return vm.runtimeErr("E300", line, col, "db scan err: %v", err)
				}

				rowMap := make(map[string]core.Value)
				for i, colName := range cols {
					val := columns[i]
					if val == nil {
						rowMap[colName] = core.UnknownValue
						continue
					}
					switch v := val.(type) {
					case []byte:
						rowMap[colName] = core.StringVal(string(v))
					case string:
						rowMap[colName] = core.StringVal(v)
					case int64:
						rowMap[colName] = core.NumberVal(float64(v))
					case int:
						rowMap[colName] = core.NumberVal(float64(v))
					case int32:
						rowMap[colName] = core.NumberVal(float64(v))
					case float64:
						rowMap[colName] = core.NumberVal(v)
					case bool:
						rowMap[colName] = core.BoolVal(v)
					default:
						rowMap[colName] = core.StringVal(fmt.Sprintf("%v", v))
					}
				}
				resultArr = append(resultArr, core.Value{core.ValMap, rowMap})
			}
			result = core.Value{core.ValArray, resultArr}
		} else {
			res, err := db.Exec(query, qargs...)
			if err != nil {
				return vm.runtimeErr("E300", line, col, "db exec err: %v", err)
			}

			lastId, _ := res.LastInsertId()
			rowsAffected, _ := res.RowsAffected()

			m := make(map[string]core.Value)
			m["last_insert_id"] = core.NumberVal(float64(lastId))
			m["rows_affected"] = core.NumberVal(float64(rowsAffected))
			result = core.Value{core.ValMap, m}
		}

	case core.BuiltinNow:
		result = core.NumberVal(float64(time.Now().UnixMilli()))

	case core.BuiltinDate:
		now := time.Now()
		m := make(map[string]core.Value)
		m["year"] = core.NumberVal(float64(now.Year()))
		m["month"] = core.NumberVal(float64(now.Month()))
		m["day"] = core.NumberVal(float64(now.Day()))
		m["hour"] = core.NumberVal(float64(now.Hour()))
		m["min"] = core.NumberVal(float64(now.Minute()))
		m["sec"] = core.NumberVal(float64(now.Second()))
		m["format"] = core.StringVal(now.Format("2006-01-02 15:04:05"))
		result = core.Value{core.ValMap, m}

	case core.BuiltinReq:
		if len(args) < 1 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want url string")
		}
		reqURL := args[0].Data.(string)
		req, err := http.NewRequest("GET", reqURL, nil)
		if err != nil {
			return vm.runtimeErr("E300", line, col, "req fail: %v", err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return vm.runtimeErr("E300", line, col, "req fail: %v", err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return vm.runtimeErr("E300", line, col, "read fail: %v", err)
		}
		result = core.StringVal(string(body))

	case core.BuiltinJson:
		if len(args) != 1 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want json string")
		}
		// Simple JSON parser: converts JSON string to dryLang map/array
		jsonStr := args[0].Data.(string)
		parsed, err := parseJSON(jsonStr)
		if err != nil {
			return vm.runtimeErr("E300", line, col, "json parse fail: %v", err)
		}
		result = parsed

	case core.BuiltinArg:
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

	case core.BuiltinEnv:
		if len(args) != 1 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want string")
		}
		result = core.StringVal(os.Getenv(args[0].Data.(string)))

	case core.BuiltinCmd:
		if len(args) < 1 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want command string")
		}
		cmdStr := args[0].Data.(string)
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
		out, err := cmd.CombinedOutput()
		if err != nil {
			return vm.runtimeErr("E300", line, col, "cmd fail: %v\n%s", err, string(out))
		}
		result = core.StringVal(strings.TrimRight(string(out), "\n\r"))

	case core.BuiltinDir:
		if len(args) != 1 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want path string")
		}
		entries, err := os.ReadDir(args[0].Data.(string))
		if err != nil {
			return vm.runtimeErr("E300", line, col, "dir fail: %v", err)
		}
		arr := make([]core.Value, len(entries))
		for i, e := range entries {
			arr[i] = core.StringVal(e.Name())
		}
		result = core.ArrayVal(arr)

	case core.BuiltinDel:
		if len(args) != 1 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want path string")
		}
		err := os.Remove(args[0].Data.(string))
		if err != nil {
			return vm.runtimeErr("E300", line, col, "del fail: %v", err)
		}
		result = core.BoolVal(true)

	case core.BuiltinDie:
		if len(args) > 0 {
			msg := args[0].String()
			if msg != "" {
				fmt.Fprintln(os.Stderr, msg)
			}
		}
		os.Exit(1)
		result = core.UnknownValue // unreachable

	case core.BuiltinMath:
		if len(args) < 2 || args[0].Type != core.ValString {
			return vm.runtimeErr("E300", line, col, "want math(op, ...numbers)")
		}
		op := args[0].Data.(string)
		if args[1].Type != core.ValNumber {
			return vm.runtimeErr("E300", line, col, "want number")
		}
		a := args[1].Data.(float64)
		switch op {
		case "sqrt":
			result = core.NumberVal(math.Sqrt(a))
		case "pow":
			if len(args) < 3 || args[2].Type != core.ValNumber {
				return vm.runtimeErr("E300", line, col, "pow wants 2 numbers")
			}
			result = core.NumberVal(math.Pow(a, args[2].Data.(float64)))
		case "ceil":
			result = core.NumberVal(math.Ceil(a))
		case "floor":
			result = core.NumberVal(math.Floor(a))
		case "sin":
			result = core.NumberVal(math.Sin(a))
		case "cos":
			result = core.NumberVal(math.Cos(a))
		case "tan":
			result = core.NumberVal(math.Tan(a))
		case "log":
			result = core.NumberVal(math.Log(a))
		case "log10":
			result = core.NumberVal(math.Log10(a))
		default:
			return vm.runtimeErr("E300", line, col, "unknown math op: %s", op)
		}

	case core.BuiltinPt:
		var strs []string
		for _, arg := range args {
			strs = append(strs, arg.String())
		}
		fmt.Println(strings.Join(strs, " "))
		result = core.UnknownValue

	case core.BuiltinIn:
		if argCount > 0 {
			fmt.Print(args[0].String())
		}
		if vm.stdinScanner == nil {
			vm.stdinScanner = bufio.NewScanner(os.Stdin)
		}
		if vm.stdinScanner.Scan() {
			text := strings.TrimSpace(vm.stdinScanner.Text())
			result = core.StringVal(text)
		} else {
			result = core.StringVal("")
		}

	default:
		return vm.runtimeErr("E300", line, col, "unknown builtin")
	}

	vm.push(result)
	return nil
}

func valuesEqual(a, b core.Value) bool {
	if a.Type != b.Type {
		return false
	}
	switch a.Type {
	case core.ValNumber:
		return a.Data.(float64) == b.Data.(float64)
	case core.ValString:
		return a.Data.(string) == b.Data.(string)
	case core.ValBool:
		return a.Data.(bool) == b.Data.(bool)
	case core.ValUnknown:
		return b.Type == core.ValUnknown
	}
	return false
}

func parseJSON(input string) (core.Value, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return core.UnknownValue, err
	}
	return jsonToValue(raw), nil
}

func jsonToValue(v interface{}) core.Value {
	if v == nil {
		return core.UnknownValue
	}
	switch val := v.(type) {
	case float64:
		return core.NumberVal(val)
	case string:
		return core.StringVal(val)
	case bool:
		return core.BoolVal(val)
	case []interface{}:
		arr := make([]core.Value, len(val))
		for i, item := range val {
			arr[i] = jsonToValue(item)
		}
		return core.ArrayVal(arr)
	case map[string]interface{}:
		m := make(map[string]core.Value, len(val))
		for k, item := range val {
			m[k] = jsonToValue(item)
		}
		return core.MapVal(m)
	}
	return core.UnknownValue
}

func (vm *VM) callOp(argCount, line, col int) error {
	chunk := vm.chunk
	callee := vm.stack[vm.sp-argCount-1]

	if callee.Type == core.ValStructDef {
		def := callee.Data.(core.StructDef)
		if argCount != len(def.Fields) {
			return vm.runtimeErr("E300", line, col, "want %d args for struct", len(def.Fields))
		}
		fields := make(map[string]core.Value)
		fields["__struct__"] = core.StringVal(def.Name)
		for i := 0; i < argCount; i++ {
			fields[def.Fields[i]] = vm.stack[vm.sp-argCount+i]
		}
		vm.sp -= argCount + 1
		vm.push(core.Value{Type: core.ValStructInstance, Data: fields})
		return nil
	}

	if callee.Type == core.ValClass {
		def := callee.Data.(core.ClassDef)
		inst := &core.Instance{
			Class:  def,
			Fields: make(map[string]core.Value),
		}
		
		// Initialize fields from all parents (DFS)
		var initFields func(c *core.ClassDef)
		initFields = func(c *core.ClassDef) {
			for _, p := range c.Parents {
				if p != nil {
					initFields(p)
				}
			}
			for _, fieldName := range c.Fields {
				inst.Fields[fieldName] = core.UnknownValue
			}
		}
		initFields(&def)

		if initM, hasInit := def.Methods["init"]; hasInit {
			vm.stack[vm.sp-argCount-1] = core.Value{Type: core.ValInstance, Data: inst}

			frame := callFrame{
				fn:    &core.CompiledFn{Chunk: initM.Chunk, Name: "init", ParamCount: argCount},
				ip:    vm.ip,
				bp:    vm.sp - argCount - 1,
				chunk: chunk,
			}
			vm.frames = append(vm.frames, frame)

			savedIP := vm.ip
			savedChunk := chunk
			err := vm.execute(initM.Chunk)
			if err != nil {
				return err
			}

			vm.frames = vm.frames[:len(vm.frames)-1]
			vm.pop() // init returns nothing

			vm.sp = frame.bp
			if vm.sp < 0 {
				vm.sp = 0
			}
			vm.push(core.Value{Type: core.ValInstance, Data: inst})

			vm.ip = savedIP
			chunk = savedChunk
			vm.chunk = chunk
			return nil
		} else {
			if argCount > 0 {
				return vm.runtimeErr("E300", line, col, "class has no init but got args")
			}
			vm.sp -= argCount + 1
			vm.push(core.Value{Type: core.ValInstance, Data: inst})
			return nil
		}
	}

	if callee.Type == core.ValBoundMethod {
		bound := callee.Data.(*core.BoundMethod)
		vm.stack[vm.sp-argCount-1] = core.Value{Type: core.ValInstance, Data: bound.Instance}

		fn := &core.CompiledFn{Chunk: bound.Method.Chunk, Name: "<method>", ParamCount: argCount}

		frame := callFrame{
			fn:    fn,
			ip:    vm.ip,
			bp:    vm.sp - argCount - 1,
			chunk: chunk,
		}
		vm.frames = append(vm.frames, frame)

		savedIP := vm.ip
		savedChunk := chunk
		err := vm.execute(bound.Method.Chunk)
		if err != nil {
			return err
		}

		vm.frames = vm.frames[:len(vm.frames)-1]
		result := vm.pop()

		vm.sp = frame.bp
		if vm.sp < 0 {
			vm.sp = 0
		}
		vm.push(result)

		vm.ip = savedIP
		chunk = savedChunk
		vm.chunk = chunk
		return nil
	}

	if callee.Type != core.ValFn {
		return vm.runtimeErr("E300", line, col, "want fn")
	}

	fn := callee.Data.(*core.CompiledFn)
	if argCount != fn.ParamCount {
		return vm.runtimeErr("E300", line, col, "want %d args", fn.ParamCount)
	}

	frame := callFrame{
		fn:    fn,
		ip:    vm.ip,
		bp:    vm.sp - argCount,
		chunk: chunk,
	}
	vm.frames = append(vm.frames, frame)

	savedIP := vm.ip
	savedChunk := chunk
	err := vm.execute(fn.Chunk)
	if err != nil {
		if len(vm.tryStack) > 0 {
			tf := vm.tryStack[len(vm.tryStack)-1]
			vm.tryStack = vm.tryStack[:len(vm.tryStack)-1]
			vm.sp = tf.sp
			vm.push(core.StringVal(err.Error()))
			vm.ip = tf.catchIP
			chunk = savedChunk
			vm.chunk = chunk
			for len(vm.frames) > tf.frameDepth {
				vm.frames = vm.frames[:len(vm.frames)-1]
			}
			return nil
		}
		return err
	}

	vm.frames = vm.frames[:len(vm.frames)-1]
	result := vm.pop()

	vm.sp = frame.bp - 1
	if vm.sp < 0 {
		vm.sp = 0
	}
	vm.push(result)

	vm.ip = savedIP
	chunk = savedChunk
	vm.chunk = chunk
	return nil
}
