package vm

import (
	"drylang/compiler"
	"drylang/errfmt"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Value types
const (
	ValNumber  = "number"
	ValString  = "string"
	ValBool    = "bool"
	ValArray   = "array"
	ValMap     = "map"
	ValFn      = "fn"
	ValUnknown = "unknown"
)

// Value wraps any dryLang runtime value.
type Value struct {
	Type string
	Data interface{}
}

var UnknownValue = Value{Type: ValUnknown, Data: nil}

func NumberVal(v float64) Value          { return Value{ValNumber, v} }
func StringVal(v string) Value           { return Value{ValString, v} }
func BoolVal(v bool) Value               { return Value{ValBool, v} }
func ArrayVal(v []Value) Value           { return Value{ValArray, v} }
func MapVal(v map[string]Value) Value    { return Value{ValMap, v} }
func FnVal(v *compiler.CompiledFn) Value { return Value{ValFn, v} }

func (v Value) String() string {
	switch v.Type {
	case ValNumber:
		f := v.Data.(float64)
		if f == math.Trunc(f) {
			return strconv.FormatInt(int64(f), 10)
		}
		return strconv.FormatFloat(f, 'f', -1, 64)
	case ValString:
		return v.Data.(string)
	case ValBool:
		if v.Data.(bool) {
			return "t"
		}
		return "f"
	case ValArray:
		arr := v.Data.([]Value)
		parts := make([]string, len(arr))
		for i, item := range arr {
			parts[i] = item.String()
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case ValMap:
		m := v.Data.(map[string]Value)
		parts := make([]string, 0, len(m))
		for k, val := range m {
			parts = append(parts, fmt.Sprintf("%s: %s", k, val.String()))
		}
		return "{" + strings.Join(parts, ", ") + "}"
	case ValFn:
		fn := v.Data.(*compiler.CompiledFn)
		return fmt.Sprintf("<fn %s>", fn.Name)
	case ValUnknown:
		return "unknown"
	}
	return "?"
}

func isTruthy(v Value) bool {
	switch v.Type {
	case ValBool:
		return v.Data.(bool)
	case ValNumber:
		return v.Data.(float64) != 0
	case ValString:
		return v.Data.(string) != ""
	case ValArray:
		return len(v.Data.([]Value)) > 0
	case ValMap:
		return len(v.Data.(map[string]Value)) > 0
	case ValUnknown:
		return false
	}
	return true
}

// VM executes dryLang bytecode.
type VM struct {
	chunk    *compiler.Chunk
	fns      []*compiler.CompiledFn
	globals  map[string]Value
	stack    []Value
	sp       int // stack pointer
	ip       int // instruction pointer
	frames   []callFrame
	tryStack []tryFrame
}

type callFrame struct {
	fn    *compiler.CompiledFn
	ip    int
	bp    int // base pointer (stack offset)
	chunk *compiler.Chunk
}

type tryFrame struct {
	catchIP    int
	sp         int
	frameDepth int
}

// New creates a new VM.
func New(chunk *compiler.Chunk, fns []*compiler.CompiledFn) *VM {
	return &VM{
		chunk:   chunk,
		fns:     fns,
		globals: make(map[string]Value),
		stack:   make([]Value, 4096),
		sp:      0,
		ip:      0,
	}
}

func (vm *VM) push(v Value) {
	vm.stack[vm.sp] = v
	vm.sp++
}

func (vm *VM) pop() Value {
	vm.sp--
	return vm.stack[vm.sp]
}

func (vm *VM) Update(chunk *compiler.Chunk, fns []*compiler.CompiledFn) {
	vm.chunk = chunk
	vm.fns = fns
}

func (vm *VM) peek() Value {
	return vm.stack[vm.sp-1]
}

func (vm *VM) runtimeErr(code string, line, col int, format string, args ...interface{}) error {
	return errfmt.Format(code, line, col, fmt.Sprintf(format, args...))
}

// Run executes the main chunk.
func (vm *VM) Run() error {
	return vm.execute(vm.chunk)
}

func (vm *VM) execute(chunk *compiler.Chunk) error {
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
		case compiler.OpConst:
			val := chunk.Constants[inst.Operand]
			switch v := val.(type) {
			case float64:
				vm.push(NumberVal(v))
			case string:
				vm.push(StringVal(v))
			case *compiler.CompiledFn:
				vm.push(FnVal(v))
			case compiler.StructDef:
				vm.push(Value{ValMap, map[string]Value{"__struct__": StringVal(v.Name)}})
			default:
				vm.push(Value{ValUnknown, nil})
			}

		case compiler.OpPop:
			if vm.sp > 0 {
				vm.pop()
			}

		case compiler.OpTrue:
			vm.push(BoolVal(true))
		case compiler.OpFalse:
			vm.push(BoolVal(false))
		case compiler.OpUnknown:
			vm.push(UnknownValue)

		case compiler.OpAdd:
			b, a := vm.pop(), vm.pop()
			if a.Type == ValString || b.Type == ValString {
				vm.push(StringVal(a.String() + b.String()))
			} else if a.Type == ValNumber && b.Type == ValNumber {
				vm.push(NumberVal(a.Data.(float64) + b.Data.(float64)))
			} else {
				return vm.runtimeErr("E300", line, col, "want number")
			}

		case compiler.OpSub:
			b, a := vm.pop(), vm.pop()
			if a.Type != ValNumber || b.Type != ValNumber {
				return vm.runtimeErr("E300", line, col, "want number")
			}
			vm.push(NumberVal(a.Data.(float64) - b.Data.(float64)))

		case compiler.OpMul:
			b, a := vm.pop(), vm.pop()
			if a.Type != ValNumber || b.Type != ValNumber {
				return vm.runtimeErr("E300", line, col, "want number")
			}
			vm.push(NumberVal(a.Data.(float64) * b.Data.(float64)))

		case compiler.OpDiv:
			b, a := vm.pop(), vm.pop()
			if a.Type != ValNumber || b.Type != ValNumber {
				return vm.runtimeErr("E300", line, col, "want number")
			}
			if b.Data.(float64) == 0 {
				return vm.runtimeErr("E300", line, col, "div by 0")
			}
			vm.push(NumberVal(a.Data.(float64) / b.Data.(float64)))

		case compiler.OpMod:
			b, a := vm.pop(), vm.pop()
			if a.Type != ValNumber || b.Type != ValNumber {
				return vm.runtimeErr("E300", line, col, "want number")
			}
			vm.push(NumberVal(math.Mod(a.Data.(float64), b.Data.(float64))))

		case compiler.OpNeg:
			v := vm.pop()
			if v.Type != ValNumber {
				return vm.runtimeErr("E300", line, col, "want number")
			}
			vm.push(NumberVal(-v.Data.(float64)))

		case compiler.OpEqual:
			b, a := vm.pop(), vm.pop()
			vm.push(BoolVal(valuesEqual(a, b)))

		case compiler.OpNotEqual:
			b, a := vm.pop(), vm.pop()
			vm.push(BoolVal(!valuesEqual(a, b)))

		case compiler.OpLess:
			b, a := vm.pop(), vm.pop()
			if a.Type == ValNumber && b.Type == ValNumber {
				vm.push(BoolVal(a.Data.(float64) < b.Data.(float64)))
			} else {
				vm.push(BoolVal(a.String() < b.String()))
			}

		case compiler.OpGreater:
			b, a := vm.pop(), vm.pop()
			if a.Type == ValNumber && b.Type == ValNumber {
				vm.push(BoolVal(a.Data.(float64) > b.Data.(float64)))
			} else {
				vm.push(BoolVal(a.String() > b.String()))
			}

		case compiler.OpLessEq:
			b, a := vm.pop(), vm.pop()
			if a.Type == ValNumber && b.Type == ValNumber {
				vm.push(BoolVal(a.Data.(float64) <= b.Data.(float64)))
			} else {
				vm.push(BoolVal(a.String() <= b.String()))
			}

		case compiler.OpGreaterEq:
			b, a := vm.pop(), vm.pop()
			if a.Type == ValNumber && b.Type == ValNumber {
				vm.push(BoolVal(a.Data.(float64) >= b.Data.(float64)))
			} else {
				vm.push(BoolVal(a.String() >= b.String()))
			}

		case compiler.OpAnd:
			b, a := vm.pop(), vm.pop()
			vm.push(BoolVal(isTruthy(a) && isTruthy(b)))

		case compiler.OpOr:
			b, a := vm.pop(), vm.pop()
			vm.push(BoolVal(isTruthy(a) || isTruthy(b)))

		case compiler.OpNot:
			v := vm.pop()
			vm.push(BoolVal(!isTruthy(v)))

		case compiler.OpConcat:
			b, a := vm.pop(), vm.pop()
			vm.push(StringVal(a.String() + b.String()))

		case compiler.OpGetGlobal:
			name := chunk.Constants[inst.Operand].(string)
			val, ok := vm.globals[name]
			if !ok {
				return vm.runtimeErr("E300", line, col, "unknown %s", name)
			}
			vm.push(val)

		case compiler.OpSetGlobal:
			name := chunk.Constants[inst.Operand].(string)
			vm.globals[name] = vm.pop()

		case compiler.OpGetLocal:
			bp := 0
			if len(vm.frames) > 0 {
				bp = vm.frames[len(vm.frames)-1].bp
			}
			vm.push(vm.stack[bp+inst.Operand])

		case compiler.OpSetLocal:
			bp := 0
			if len(vm.frames) > 0 {
				bp = vm.frames[len(vm.frames)-1].bp
			}
			val := vm.pop()
			idx := bp + inst.Operand
			for idx >= len(vm.stack) {
				vm.stack = append(vm.stack, UnknownValue)
			}
			vm.stack[idx] = val
			if idx >= vm.sp {
				vm.sp = idx + 1
			}

		case compiler.OpJump:
			vm.ip = inst.Operand

		case compiler.OpJumpIfFalse:
			cond := vm.pop()
			if !isTruthy(cond) {
				vm.ip = inst.Operand
			}

		case compiler.OpLoop:
			vm.ip = inst.Operand

		case compiler.OpArray:
			count := inst.Operand
			arr := make([]Value, count)
			for i := count - 1; i >= 0; i-- {
				arr[i] = vm.pop()
			}
			vm.push(ArrayVal(arr))

		case compiler.OpMap:
			count := inst.Operand
			m := make(map[string]Value, count)
			for i := 0; i < count; i++ {
				val := vm.pop()
				key := vm.pop()
				m[key.String()] = val
			}
			vm.push(MapVal(m))

		case compiler.OpIndex:
			idx := vm.pop()
			obj := vm.pop()
			switch obj.Type {
			case ValArray:
				arr := obj.Data.([]Value)
				i := int(idx.Data.(float64))
				if i < 0 || i >= len(arr) {
					return vm.runtimeErr("E300", line, col, "bounds %d", i)
				}
				vm.push(arr[i])
			case ValMap:
				m := obj.Data.(map[string]Value)
				key := idx.String()
				if val, ok := m[key]; ok {
					vm.push(val)
				} else {
					vm.push(UnknownValue)
				}
			case ValString:
				s := obj.Data.(string)
				i := int(idx.Data.(float64))
				if i < 0 || i >= len(s) {
					return vm.runtimeErr("E300", line, col, "bounds %d", i)
				}
				vm.push(StringVal(string(s[i])))
			default:
				return vm.runtimeErr("E300", line, col, "want array|map")
			}

		case compiler.OpSetIndex:
			val := vm.pop()
			idx := vm.pop()
			obj := vm.pop()
			switch obj.Type {
			case ValArray:
				arr := obj.Data.([]Value)
				i := int(idx.Data.(float64))
				if i < 0 || i >= len(arr) {
					return vm.runtimeErr("E300", line, col, "bounds %d", i)
				}
				arr[i] = val
			case ValMap:
				m := obj.Data.(map[string]Value)
				m[idx.String()] = val
			default:
				return vm.runtimeErr("E300", line, col, "want array|map")
			}

		case compiler.OpDotGet:
			field := chunk.Constants[inst.Operand].(string)
			obj := vm.pop()
			if obj.Type != ValMap {
				return vm.runtimeErr("E300", line, col, "want map")
			}
			m := obj.Data.(map[string]Value)
			if val, ok := m[field]; ok {
				vm.push(val)
			} else {
				vm.push(UnknownValue)
			}

		case compiler.OpDotSet:
			field := chunk.Constants[inst.Operand].(string)
			val := vm.pop()
			obj := vm.pop()
			if obj.Type != ValMap {
				return vm.runtimeErr("E300", line, col, "want map")
			}
			obj.Data.(map[string]Value)[field] = val

		case compiler.OpPrint:
			v := vm.pop()
			fmt.Println(v.String())

		case compiler.OpInput:
			if inst.Operand > 0 {
				prompt := vm.pop()
				fmt.Print(prompt.String())
			}
			var input string
			fmt.Scanln(&input)
			vm.push(StringVal(input))

		case compiler.OpClosure:
			fn := chunk.Constants[inst.Operand].(*compiler.CompiledFn)
			vm.push(FnVal(fn))

		case compiler.OpCall:
			argCount := inst.Operand
			callee := vm.stack[vm.sp-argCount-1]

			if callee.Type != ValFn {
				return vm.runtimeErr("E300", line, col, "want fn")
			}

			fn := callee.Data.(*compiler.CompiledFn)
			if argCount != fn.ParamCount {
				fmt.Printf("DEBUG OpCall: argCount=%d, fn.ParamCount=%d\n", argCount, fn.ParamCount)
				return vm.runtimeErr("E300", line, col, "want %d args", fn.ParamCount)
			}

			frame := callFrame{
				fn:    fn,
				ip:    vm.ip,
				bp:    vm.sp - argCount,
				chunk: chunk,
			}
			vm.frames = append(vm.frames, frame)

			// Execute function chunk
			savedIP := vm.ip
			savedChunk := chunk
			err := vm.execute(fn.Chunk)
			if err != nil {
				// Check try/catch
				if len(vm.tryStack) > 0 {
					tf := vm.tryStack[len(vm.tryStack)-1]
					vm.tryStack = vm.tryStack[:len(vm.tryStack)-1]
					vm.sp = tf.sp
					vm.push(StringVal(err.Error()))
					vm.ip = tf.catchIP
					chunk = savedChunk
					vm.chunk = chunk
					continue
				}
				return err
			}

			// Pop frame
			vm.frames = vm.frames[:len(vm.frames)-1]
			result := vm.pop()

			// Clean up stack: remove function + args
			vm.sp = frame.bp - 1
			if vm.sp < 0 {
				vm.sp = 0
			}
			vm.push(result)

			vm.ip = savedIP
			chunk = savedChunk
			vm.chunk = chunk

		case compiler.OpReturn:
			if len(vm.frames) > 0 {
				return nil // return from function execution
			}
			// Top-level rev = exit
			return nil

		case compiler.OpTry:
			vm.tryStack = append(vm.tryStack, tryFrame{
				catchIP:    inst.Operand,
				sp:         vm.sp,
				frameDepth: len(vm.frames),
			})

		case compiler.OpEndTry:
			if len(vm.tryStack) > 0 {
				vm.tryStack = vm.tryStack[:len(vm.tryStack)-1]
			}

		case compiler.OpThrow:
			v := vm.pop()
			if len(vm.tryStack) > 0 {
				tf := vm.tryStack[len(vm.tryStack)-1]
				vm.tryStack = vm.tryStack[:len(vm.tryStack)-1]
				vm.sp = tf.sp
				vm.push(v)
				vm.ip = tf.catchIP
			} else {
				return vm.runtimeErr("E300", line, col, "%s", v.String())
			}

		case compiler.OpBuiltin:
			if err := vm.executeBuiltin(compiler.BuiltinID(inst.Operand), inst.Operand2, line, col); err != nil {
				return err
			}

		case compiler.OpGetLoopCounter:
			// handled via locals

		case compiler.OpAsync:
			// simplified: just execute normally for now

		case compiler.OpAwait:
			// simplified: value is already on stack
		}
	}

	return nil
}

func (vm *VM) executeBuiltin(id compiler.BuiltinID, argCount int, line, col int) error {
	args := make([]Value, argCount)
	for i := argCount - 1; i >= 0; i-- {
		args[i] = vm.pop()
	}

	var result Value

	switch id {
	case compiler.BuiltinLen:
		if len(args) != 1 {
			return vm.runtimeErr("E300", line, col, "want 1 arg")
		}
		switch args[0].Type {
		case ValString:
			result = NumberVal(float64(len(args[0].Data.(string))))
		case ValArray:
			result = NumberVal(float64(len(args[0].Data.([]Value))))
		case ValMap:
			result = NumberVal(float64(len(args[0].Data.(map[string]Value))))
		default:
			return vm.runtimeErr("E300", line, col, "want string|array|map")
		}

	case compiler.BuiltinGet:
		if len(args) != 1 {
			return vm.runtimeErr("E300", line, col, "want 1 arg")
		}
		result = StringVal(args[0].Type)

	case compiler.BuiltinAdd:
		if len(args) != 2 {
			return vm.runtimeErr("E300", line, col, "want 2 args")
		}
		if args[0].Type != ValArray {
			return vm.runtimeErr("E300", line, col, "want array")
		}
		arr := args[0].Data.([]Value)
		arr = append(arr, args[1])
		// Modify in place via globals would need ref semantics.
		// For now, push new array.
		result = ArrayVal(arr)

	case compiler.BuiltinNum:
		if len(args) != 1 {
			return vm.runtimeErr("E300", line, col, "want 1 arg")
		}
		f, err := strconv.ParseFloat(args[0].String(), 64)
		if err != nil {
			return vm.runtimeErr("E300", line, col, "bad number")
		}
		result = NumberVal(f)

	case compiler.BuiltinStr:
		if len(args) != 1 {
			return vm.runtimeErr("E300", line, col, "want 1 arg")
		}
		result = StringVal(args[0].String())

	case compiler.BuiltinAbs:
		if len(args) != 1 || args[0].Type != ValNumber {
			return vm.runtimeErr("E300", line, col, "want number")
		}
		result = NumberVal(math.Abs(args[0].Data.(float64)))

	case compiler.BuiltinMin:
		if len(args) != 2 || args[0].Type != ValNumber || args[1].Type != ValNumber {
			return vm.runtimeErr("E300", line, col, "want 2 numbers")
		}
		result = NumberVal(math.Min(args[0].Data.(float64), args[1].Data.(float64)))

	case compiler.BuiltinMax:
		if len(args) != 2 || args[0].Type != ValNumber || args[1].Type != ValNumber {
			return vm.runtimeErr("E300", line, col, "want 2 numbers")
		}
		result = NumberVal(math.Max(args[0].Data.(float64), args[1].Data.(float64)))

	case compiler.BuiltinRnd:
		if len(args) != 1 || args[0].Type != ValNumber {
			return vm.runtimeErr("E300", line, col, "want number")
		}
		result = NumberVal(math.Round(args[0].Data.(float64)))

	case compiler.BuiltinCap:
		if len(args) != 1 || args[0].Type != ValString {
			return vm.runtimeErr("E300", line, col, "want string")
		}
		result = StringVal(strings.ToUpper(args[0].Data.(string)))

	case compiler.BuiltinLow:
		if len(args) != 1 || args[0].Type != ValString {
			return vm.runtimeErr("E300", line, col, "want string")
		}
		result = StringVal(strings.ToLower(args[0].Data.(string)))

	case compiler.BuiltinTrm:
		if len(args) != 1 || args[0].Type != ValString {
			return vm.runtimeErr("E300", line, col, "want string")
		}
		result = StringVal(strings.TrimSpace(args[0].Data.(string)))

	case compiler.BuiltinSpl:
		if len(args) != 2 || args[0].Type != ValString || args[1].Type != ValString {
			return vm.runtimeErr("E300", line, col, "want 2 strings")
		}
		parts := strings.Split(args[0].Data.(string), args[1].Data.(string))
		arr := make([]Value, len(parts))
		for i, p := range parts {
			arr[i] = StringVal(p)
		}
		result = ArrayVal(arr)

	case compiler.BuiltinJ:
		if len(args) != 2 || args[0].Type != ValArray || args[1].Type != ValString {
			return vm.runtimeErr("E300", line, col, "want array, string")
		}
		arr := args[0].Data.([]Value)
		strs := make([]string, len(arr))
		for i, v := range arr {
			strs[i] = v.String()
		}
		result = StringVal(strings.Join(strs, args[1].Data.(string)))

	case compiler.BuiltinMod:
		if len(args) != 3 || args[0].Type != ValString || args[1].Type != ValString || args[2].Type != ValString {
			return vm.runtimeErr("E300", line, col, "want 3 strings")
		}
		result = StringVal(strings.ReplaceAll(args[0].Data.(string), args[1].Data.(string), args[2].Data.(string)))

	case compiler.BuiltinHas:
		if len(args) != 2 || args[0].Type != ValString || args[1].Type != ValString {
			return vm.runtimeErr("E300", line, col, "want 2 strings")
		}
		result = BoolVal(strings.Contains(args[0].Data.(string), args[1].Data.(string)))

	case compiler.BuiltinSort:
		if len(args) != 1 || args[0].Type != ValArray {
			return vm.runtimeErr("E300", line, col, "want array")
		}
		arr := make([]Value, len(args[0].Data.([]Value)))
		copy(arr, args[0].Data.([]Value))
		sort.Slice(arr, func(i, j int) bool {
			if arr[i].Type == ValNumber && arr[j].Type == ValNumber {
				return arr[i].Data.(float64) < arr[j].Data.(float64)
			}
			return arr[i].String() < arr[j].String()
		})
		result = ArrayVal(arr)

	case compiler.BuiltinPop:
		if len(args) != 1 || args[0].Type != ValArray {
			return vm.runtimeErr("E300", line, col, "want array")
		}
		arr := args[0].Data.([]Value)
		if len(arr) == 0 {
			return vm.runtimeErr("E300", line, col, "empty array")
		}
		result = arr[len(arr)-1]

	case compiler.BuiltinRm:
		if len(args) != 2 || args[0].Type != ValArray || args[1].Type != ValNumber {
			return vm.runtimeErr("E300", line, col, "want array, number")
		}
		arr := args[0].Data.([]Value)
		idx := int(args[1].Data.(float64))
		if idx < 0 || idx >= len(arr) {
			return vm.runtimeErr("E300", line, col, "bounds %d", idx)
		}
		newArr := make([]Value, 0, len(arr)-1)
		newArr = append(newArr, arr[:idx]...)
		newArr = append(newArr, arr[idx+1:]...)
		result = ArrayVal(newArr)

	case compiler.BuiltinKey:
		if len(args) != 1 || args[0].Type != ValMap {
			return vm.runtimeErr("E300", line, col, "want map")
		}
		m := args[0].Data.(map[string]Value)
		keys := make([]Value, 0, len(m))
		for k := range m {
			keys = append(keys, StringVal(k))
		}
		result = ArrayVal(keys)

	case compiler.BuiltinVal:
		if len(args) != 1 || args[0].Type != ValMap {
			return vm.runtimeErr("E300", line, col, "want map")
		}
		m := args[0].Data.(map[string]Value)
		vals := make([]Value, 0, len(m))
		for _, v := range m {
			vals = append(vals, v)
		}
		result = ArrayVal(vals)

	case compiler.BuiltinRan:
		result = NumberVal(rand.Float64())

	case compiler.BuiltinQ:
		if len(args) != 1 || args[0].Type != ValNumber {
			return vm.runtimeErr("E300", line, col, "want number")
		}
		ms := int(args[0].Data.(float64))
		time.Sleep(time.Duration(ms) * time.Millisecond)
		result = UnknownValue

	case compiler.BuiltinR:
		if len(args) < 1 || args[0].Type != ValString {
			return vm.runtimeErr("E300", line, col, "want string")
		}
		data, err := os.ReadFile(args[0].Data.(string))
		if err != nil {
			return vm.runtimeErr("E300", line, col, "read fail")
		}
		result = StringVal(string(data))

	case compiler.BuiltinW:
		if len(args) != 2 || args[0].Type != ValString {
			return vm.runtimeErr("E300", line, col, "want path, data")
		}
		err := os.WriteFile(args[0].Data.(string), []byte(args[1].String()), 0644)
		if err != nil {
			return vm.runtimeErr("E300", line, col, "write fail")
		}
		result = BoolVal(true)

	default:
		return vm.runtimeErr("E300", line, col, "unknown builtin")
	}

	vm.push(result)
	return nil
}

func valuesEqual(a, b Value) bool {
	if a.Type != b.Type {
		return false
	}
	switch a.Type {
	case ValNumber:
		return a.Data.(float64) == b.Data.(float64)
	case ValString:
		return a.Data.(string) == b.Data.(string)
	case ValBool:
		return a.Data.(bool) == b.Data.(bool)
	case ValUnknown:
		return b.Type == ValUnknown
	}
	return false
}
