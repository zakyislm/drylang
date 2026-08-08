package compiler

// Opcode represents a bytecode instruction.
type Opcode byte

const (
	OpConst Opcode = iota
	OpPop

	// Arithmetic
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpMod
	OpNeg // unary -

	// Comparison
	OpEqual
	OpNotEqual
	OpLess
	OpGreater
	OpLessEq
	OpGreaterEq

	// Logic
	OpAnd
	OpOr
	OpNot

	// Boolean & Unknown
	OpTrue
	OpFalse
	OpUnknown

	// Variables
	OpGetGlobal
	OpSetGlobal
	OpGetLocal
	OpSetLocal

	// Jump
	OpJump
	OpJumpIfFalse
	OpLoop // jump backwards

	// Collections
	OpArray
	OpMap
	OpIndex
	OpSetIndex
	OpDotGet
	OpDotSet

	// Functions
	OpCall
	OpReturn
	OpClosure

	// I/O
	OpPrint
	OpInput

	// String
	OpConcat

	// Async
	OpAsync
	OpAwait

	// Error handling
	OpTry
	OpThrow
	OpEndTry

	// Built-in calls
	OpBuiltin

	// Loop counter
	OpGetLoopCounter
)

// Instruction is a single bytecode instruction with operands.
type Instruction struct {
	Op       Opcode
	Operand  int // for jumps, const index, local slot, builtin id, etc.
	Operand2 int // for 2-operand instructions (e.g., OpMap key count)
}

// Chunk holds compiled bytecode and constants.
type Chunk struct {
	Code      []Instruction
	Constants []interface{} // string, float64, etc.
	Lines     []int         // source line for each instruction
}

func (c *Chunk) Emit(op Opcode, operand int, line int) int {
	idx := len(c.Code)
	c.Code = append(c.Code, Instruction{Op: op, Operand: operand})
	c.Lines = append(c.Lines, line)
	return idx
}

func (c *Chunk) Emit2(op Opcode, operand, operand2 int, line int) int {
	idx := len(c.Code)
	c.Code = append(c.Code, Instruction{Op: op, Operand: operand, Operand2: operand2})
	c.Lines = append(c.Lines, line)
	return idx
}

func (c *Chunk) AddConst(val interface{}) int {
	c.Constants = append(c.Constants, val)
	return len(c.Constants) - 1
}

// CompiledFn represents a compiled function.
type CompiledFn struct {
	Chunk      *Chunk
	Name       string
	ParamCount int
	IsAsync    bool
}
