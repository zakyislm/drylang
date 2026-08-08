package compiler

import (
	"drylang/lexer"
	"drylang/parser"
	"fmt"
	"strconv"
)

// BuiltinID identifies built-in functions.
type BuiltinID int

const (
	BuiltinLen BuiltinID = iota
	BuiltinGet
	BuiltinAdd
	BuiltinNum
	BuiltinStr
	BuiltinAbs
	BuiltinMin
	BuiltinMax
	BuiltinRnd
	BuiltinCap
	BuiltinLow
	BuiltinTrm
	BuiltinSpl
	BuiltinJ
	BuiltinMod
	BuiltinHas
	BuiltinSort
	BuiltinPop
	BuiltinRm
	BuiltinKey
	BuiltinVal
	BuiltinRan
	BuiltinQ
	BuiltinR
	BuiltinW
)

// BuiltinNames maps function names to builtin IDs.
var BuiltinNames = map[string]BuiltinID{
	"len":  BuiltinLen,
	"get":  BuiltinGet,
	"add":  BuiltinAdd,
	"num":  BuiltinNum,
	"str":  BuiltinStr,
	"abs":  BuiltinAbs,
	"min":  BuiltinMin,
	"max":  BuiltinMax,
	"rnd":  BuiltinRnd,
	"cap":  BuiltinCap,
	"low":  BuiltinLow,
	"trm":  BuiltinTrm,
	"spl":  BuiltinSpl,
	"j":    BuiltinJ,
	"mod":  BuiltinMod,
	"has":  BuiltinHas,
	"sort": BuiltinSort,
	"pop":  BuiltinPop,
	"rm":   BuiltinRm,
	"key":  BuiltinKey,
	"val":  BuiltinVal,
	"ran":  BuiltinRan,
	"q":    BuiltinQ,
	"r":    BuiltinR,
	"w":    BuiltinW,
}

// Compiler compiles AST to bytecode.
type Compiler struct {
	chunk   *Chunk
	locals  []local
	depth   int
	fns     []*CompiledFn
	loopCtx []loopContext
}

type local struct {
	name  string
	depth int
}

type loopContext struct {
	start  int
	breaks []int // indices to patch
}

// New creates a new compiler.
func New() *Compiler {
	return &Compiler{
		chunk: &Chunk{},
	}
}

// Compile compiles a program AST to bytecode.
func (c *Compiler) Compile(prog *parser.Program) (*Chunk, []*CompiledFn, error) {
	for _, stmt := range prog.Stmts {
		if err := c.compileStmt(stmt); err != nil {
			return nil, nil, err
		}
	}
	return c.chunk, c.fns, nil
}

func (c *Compiler) emit(op Opcode, operand int, line int) int {
	return c.chunk.Emit(op, operand, line)
}

func (c *Compiler) emit2(op Opcode, operand, operand2 int, line int) int {
	return c.chunk.Emit2(op, operand, operand2, line)
}

func (c *Compiler) addConst(val interface{}) int {
	return c.chunk.AddConst(val)
}

func (c *Compiler) errorf(line, col int, format string, args ...interface{}) error {
	return fmt.Errorf("%d:%d %s", line, col, fmt.Sprintf(format, args...))
}

func (c *Compiler) resolveLocal(name string) int {
	for i := len(c.locals) - 1; i >= 0; i-- {
		if c.locals[i].name == name {
			return i
		}
	}
	return -1
}

func (c *Compiler) addLocal(name string) int {
	c.locals = append(c.locals, local{name: name, depth: c.depth})
	return len(c.locals) - 1
}

func (c *Compiler) beginScope() {
	c.depth++
}

func (c *Compiler) endScope() {
	for len(c.locals) > 0 && c.locals[len(c.locals)-1].depth == c.depth {
		c.locals = c.locals[:len(c.locals)-1]
		c.emit(OpPop, 0, 0)
	}
	c.depth--
}

func (c *Compiler) compileStmt(stmt parser.Stmt) error {
	switch s := stmt.(type) {
	case *parser.AssignStmt:
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		idx := c.resolveLocal(s.Name)
		if idx >= 0 {
			c.emit(OpSetLocal, idx, s.Line)
		} else if c.depth > 0 {
			slot := c.addLocal(s.Name)
			c.emit(OpSetLocal, slot, s.Line)
		} else {
			ci := c.addConst(s.Name)
			c.emit(OpSetGlobal, ci, s.Line)
		}
		return nil

	case *parser.ConstDeclStmt:
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		if c.depth > 0 {
			c.addLocal(s.Name)
		} else {
			ci := c.addConst(s.Name)
			c.emit(OpSetGlobal, ci, s.Line)
		}
		return nil

	case *parser.UnknownBoolStmt:
		c.emit(OpUnknown, 0, s.Line)
		if c.depth > 0 {
			c.addLocal(s.Name)
		} else {
			ci := c.addConst(s.Name)
			c.emit(OpSetGlobal, ci, s.Line)
		}
		return nil

	case *parser.PrintStmt:
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		c.emit(OpPrint, 0, s.Line)
		return nil

	case *parser.ExprStmt:
		if err := c.compileExpr(s.Expression); err != nil {
			return err
		}
		c.emit(OpPop, 0, s.Line)
		return nil

	case *parser.ReturnStmt:
		if s.Value != nil {
			if err := c.compileExpr(s.Value); err != nil {
				return err
			}
		} else {
			c.emit(OpUnknown, 0, s.Line) // default return unknown
		}
		c.emit(OpReturn, 0, s.Line)
		return nil

	case *parser.FnDeclStmt:
		return c.compileFnDecl(s)

	case *parser.IfStmt:
		return c.compileIf(s)

	case *parser.OnStmt:
		return c.compileOn(s)

	case *parser.LoopStmt:
		return c.compileLoop(s)

	case *parser.DoneStmt:
		if len(c.loopCtx) == 0 {
			return c.errorf(s.Line, s.Col, "stray done")
		}
		ctx := &c.loopCtx[len(c.loopCtx)-1]
		jmp := c.emit(OpJump, 0, s.Line)
		ctx.breaks = append(ctx.breaks, jmp)
		return nil

	case *parser.ConStmt:
		if len(c.loopCtx) == 0 {
			return c.errorf(s.Line, s.Col, "stray con")
		}
		ctx := &c.loopCtx[len(c.loopCtx)-1]
		c.emit(OpLoop, ctx.start, s.Line)
		return nil

	case *parser.TryStmt:
		return c.compileTry(s)

	case *parser.ThrowStmt:
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		c.emit(OpThrow, 0, s.Line)
		return nil

	case *parser.UseStmt:
		// Module loading handled externally
		return nil

	case *parser.PrivateStmt:
		return c.compileStmt(s.Inner)

	case *parser.StructDeclStmt:
		// Store struct definition as a constant for runtime
		ci := c.addConst(StructDef{Name: s.Name, Fields: s.Fields})
		nameIdx := c.addConst(s.Name)
		c.emit(OpConst, ci, s.Line)
		c.emit(OpSetGlobal, nameIdx, s.Line)
		return nil

	case *parser.IndexAssignStmt:
		if err := c.compileExpr(s.Object); err != nil {
			return err
		}
		if err := c.compileExpr(s.Index); err != nil {
			return err
		}
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		c.emit(OpSetIndex, 0, s.Line)
		return nil

	case *parser.DotAssignStmt:
		if err := c.compileExpr(s.Object); err != nil {
			return err
		}
		fi := c.addConst(s.Field)
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		c.emit(OpDotSet, fi, s.Line)
		return nil
	}

	return nil
}

func (c *Compiler) compileExpr(expr parser.Expr) error {
	switch e := expr.(type) {
	case *parser.NumberLit:
		v, err := strconv.ParseFloat(e.Value, 64)
		if err != nil {
			return c.errorf(e.Line, e.Col, "bad number")
		}
		ci := c.addConst(v)
		c.emit(OpConst, ci, e.Line)

	case *parser.StringLit:
		ci := c.addConst(e.Value)
		c.emit(OpConst, ci, e.Line)

	case *parser.RawStringLit:
		ci := c.addConst(e.Value)
		c.emit(OpConst, ci, e.Line)

	case *parser.BoolLit:
		if e.Value {
			c.emit(OpTrue, 0, e.Line)
		} else {
			c.emit(OpFalse, 0, e.Line)
		}

	case *parser.UnknownLit:
		c.emit(OpUnknown, 0, e.Line)

	case *parser.Ident:
		idx := c.resolveLocal(e.Name)
		if idx >= 0 {
			c.emit(OpGetLocal, idx, e.Line)
		} else {
			ci := c.addConst(e.Name)
			c.emit(OpGetGlobal, ci, e.Line)
		}

	case *parser.BinaryExpr:
		if err := c.compileExpr(e.Left); err != nil {
			return err
		}
		if err := c.compileExpr(e.Right); err != nil {
			return err
		}
		switch e.Op {
		case lexer.TOKEN_PLUS:
			c.emit(OpAdd, 0, e.Line)
		case lexer.TOKEN_MINUS:
			c.emit(OpSub, 0, e.Line)
		case lexer.TOKEN_STAR:
			c.emit(OpMul, 0, e.Line)
		case lexer.TOKEN_SLASH:
			c.emit(OpDiv, 0, e.Line)
		case lexer.TOKEN_PERCENT:
			c.emit(OpMod, 0, e.Line)
		case lexer.TOKEN_ASSIGN:
			c.emit(OpEqual, 0, e.Line)
		case lexer.TOKEN_NOT_EQ:
			c.emit(OpNotEqual, 0, e.Line)
		case lexer.TOKEN_LT:
			c.emit(OpLess, 0, e.Line)
		case lexer.TOKEN_GT:
			c.emit(OpGreater, 0, e.Line)
		case lexer.TOKEN_LT_EQ:
			c.emit(OpLessEq, 0, e.Line)
		case lexer.TOKEN_GT_EQ:
			c.emit(OpGreaterEq, 0, e.Line)
		case lexer.TOKEN_AND:
			c.emit(OpAnd, 0, e.Line)
		case lexer.TOKEN_OR:
			c.emit(OpOr, 0, e.Line)
		}

	case *parser.UnaryExpr:
		if err := c.compileExpr(e.Operand); err != nil {
			return err
		}
		switch e.Op {
		case lexer.TOKEN_MINUS:
			c.emit(OpNeg, 0, e.Line)
		case lexer.TOKEN_NOT:
			c.emit(OpNot, 0, e.Line)
		}

	case *parser.CallExpr:
		// Check for built-in function calls
		if ident, ok := e.Callee.(*parser.Ident); ok {
			if bid, found := BuiltinNames[ident.Name]; found {
				for _, arg := range e.Args {
					if err := c.compileExpr(arg); err != nil {
						return err
					}
				}
				c.emit2(OpBuiltin, int(bid), len(e.Args), e.Line)
				return nil
			}
		}

		if err := c.compileExpr(e.Callee); err != nil {
			return err
		}
		for _, arg := range e.Args {
			if err := c.compileExpr(arg); err != nil {
				return err
			}
		}
		c.emit(OpCall, len(e.Args), e.Line)

	case *parser.IndexExpr:
		if err := c.compileExpr(e.Object); err != nil {
			return err
		}
		if err := c.compileExpr(e.Index); err != nil {
			return err
		}
		c.emit(OpIndex, 0, e.Line)

	case *parser.DotExpr:
		if err := c.compileExpr(e.Object); err != nil {
			return err
		}
		fi := c.addConst(e.Field)
		c.emit(OpDotGet, fi, e.Line)

	case *parser.ArrayLit:
		for _, item := range e.Items {
			if err := c.compileExpr(item); err != nil {
				return err
			}
		}
		c.emit(OpArray, len(e.Items), e.Line)

	case *parser.MapLit:
		for i := range e.Keys {
			if err := c.compileExpr(e.Keys[i]); err != nil {
				return err
			}
			if err := c.compileExpr(e.Values[i]); err != nil {
				return err
			}
		}
		c.emit(OpMap, len(e.Keys), e.Line)

	case *parser.ArrowFnExpr:
		fnCompiler := New()
		fnCompiler.depth = 1
		for _, p := range e.Params {
			fnCompiler.addLocal(p)
		}
		for _, s := range e.Body {
			if err := fnCompiler.compileStmt(s); err != nil {
				return err
			}
		}
		fnCompiler.emit(OpUnknown, 0, e.Line) // default return
		fnCompiler.emit(OpReturn, 0, e.Line)

		fn := &CompiledFn{
			Chunk:      fnCompiler.chunk,
			Name:       "<arrow>",
			ParamCount: len(e.Params),
		}
		c.fns = append(c.fns, fn)
		fi := c.addConst(fn)
		c.emit(OpClosure, fi, e.Line)

	case *parser.InputExpr:
		if e.Prompt != nil {
			if err := c.compileExpr(e.Prompt); err != nil {
				return err
			}
			c.emit(OpInput, 1, e.Line)
		} else {
			c.emit(OpInput, 0, e.Line)
		}

	case *parser.StringInterp:
		for i, part := range e.Parts {
			if err := c.compileExpr(part); err != nil {
				return err
			}
			if i > 0 {
				c.emit(OpConcat, 0, e.Line)
			}
		}

	case *parser.StructInitExpr:
		// Push field values and create struct instance
		ci := c.addConst(e.TypeName)
		c.emit(OpConst, ci, e.Line)
		for fname, fval := range e.Fields {
			fni := c.addConst(fname)
			c.emit(OpConst, fni, e.Line)
			if err := c.compileExpr(fval); err != nil {
				return err
			}
		}
		c.emit(OpMap, len(e.Fields), e.Line)

	case *parser.AwaitExpr:
		if err := c.compileExpr(e.Value); err != nil {
			return err
		}
		c.emit(OpAwait, 0, e.Line)
	}

	return nil
}

func (c *Compiler) compileFnDecl(s *parser.FnDeclStmt) error {
	fnCompiler := New()
	fnCompiler.depth = 1
	for _, p := range s.Params {
		fnCompiler.addLocal(p)
	}
	for _, stmt := range s.Body {
		if err := fnCompiler.compileStmt(stmt); err != nil {
			return err
		}
	}
	fnCompiler.emit(OpUnknown, 0, s.Line) // default return unknown
	fnCompiler.emit(OpReturn, 0, s.Line)

	fn := &CompiledFn{
		Chunk:      fnCompiler.chunk,
		Name:       s.Name,
		ParamCount: len(s.Params),
		IsAsync:    s.IsAsync,
	}

	c.fns = append(c.fns, fn)
	fi := c.addConst(fn)
	c.emit(OpClosure, fi, s.Line)

	if c.depth > 0 {
		c.addLocal(s.Name)
	} else {
		ci := c.addConst(s.Name)
		c.emit(OpSetGlobal, ci, s.Line)
	}

	return nil
}

func (c *Compiler) compileIf(s *parser.IfStmt) error {
	if err := c.compileExpr(s.Condition); err != nil {
		return err
	}

	// Jump to elif/else if false
	falseJump := c.emit(OpJumpIfFalse, 0, s.Line)

	c.beginScope()
	for _, stmt := range s.Body {
		if err := c.compileStmt(stmt); err != nil {
			return err
		}
	}
	c.endScope()

	// Jump over elif/else blocks
	var endJumps []int
	endJump := c.emit(OpJump, 0, s.Line)
	endJumps = append(endJumps, endJump)

	// Patch false jump
	c.chunk.Code[falseJump].Operand = len(c.chunk.Code)

	// Compile elif chains
	for _, elif := range s.ElIfs {
		if err := c.compileExpr(elif.Condition); err != nil {
			return err
		}
		elifFalseJump := c.emit(OpJumpIfFalse, 0, s.Line)

		c.beginScope()
		for _, stmt := range elif.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()

		ej := c.emit(OpJump, 0, s.Line)
		endJumps = append(endJumps, ej)
		c.chunk.Code[elifFalseJump].Operand = len(c.chunk.Code)
	}

	// Compile else
	if len(s.Else) > 0 {
		c.beginScope()
		for _, stmt := range s.Else {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()
	}

	// Patch all end jumps
	for _, ej := range endJumps {
		c.chunk.Code[ej].Operand = len(c.chunk.Code)
	}

	return nil
}

func (c *Compiler) compileOn(s *parser.OnStmt) error {
	if err := c.compileExpr(s.Value); err != nil {
		return err
	}

	var endJumps []int

	for _, cas := range s.Cases {
		// Duplicate the switch value
		c.emit(OpConst, c.addConst("__dup__"), s.Line) // placeholder
		if err := c.compileExpr(cas.Value); err != nil {
			return err
		}
		c.emit(OpEqual, 0, s.Line)
		falseJump := c.emit(OpJumpIfFalse, 0, s.Line)

		c.beginScope()
		for _, stmt := range cas.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()

		ej := c.emit(OpJump, 0, s.Line)
		endJumps = append(endJumps, ej)
		c.chunk.Code[falseJump].Operand = len(c.chunk.Code)
	}

	for _, ej := range endJumps {
		c.chunk.Code[ej].Operand = len(c.chunk.Code)
	}

	c.emit(OpPop, 0, s.Line) // pop switch value

	return nil
}

func (c *Compiler) compileLoop(s *parser.LoopStmt) error {
	loopStart := len(c.chunk.Code)
	c.loopCtx = append(c.loopCtx, loopContext{start: loopStart})

	if s.Limit != nil {
		// Counted loop: emit counter init and check
		ci := c.addConst(float64(0))
		c.emit(OpConst, ci, s.Line) // push 0 as counter
		c.addLocal("i")

		loopStart = len(c.chunk.Code) // actual loop start
		c.loopCtx[len(c.loopCtx)-1].start = loopStart

		// Check: i < limit
		idx := c.resolveLocal("i")
		c.emit(OpGetLocal, idx, s.Line)
		if err := c.compileExpr(s.Limit); err != nil {
			return err
		}
		c.emit(OpLess, 0, s.Line)
		exitJump := c.emit(OpJumpIfFalse, 0, s.Line)

		c.beginScope()
		for _, stmt := range s.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()

		// Increment i
		c.emit(OpGetLocal, idx, s.Line)
		c.emit(OpConst, c.addConst(float64(1)), s.Line)
		c.emit(OpAdd, 0, s.Line)
		c.emit(OpSetLocal, idx, s.Line)

		c.emit(OpLoop, loopStart, s.Line)
		c.chunk.Code[exitJump].Operand = len(c.chunk.Code)

		// Remove counter local
		c.locals = c.locals[:len(c.locals)-1]
		c.emit(OpPop, 0, s.Line)
	} else {
		// Infinite loop
		c.beginScope()
		for _, stmt := range s.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()
		c.emit(OpLoop, loopStart, s.Line)
	}

	// Patch breaks
	ctx := c.loopCtx[len(c.loopCtx)-1]
	for _, b := range ctx.breaks {
		c.chunk.Code[b].Operand = len(c.chunk.Code)
	}
	c.loopCtx = c.loopCtx[:len(c.loopCtx)-1]

	return nil
}

func (c *Compiler) compileTry(s *parser.TryStmt) error {
	tryJump := c.emit(OpTry, 0, s.Line)

	c.beginScope()
	for _, stmt := range s.Body {
		if err := c.compileStmt(stmt); err != nil {
			return err
		}
	}
	c.endScope()

	endJump := c.emit(OpJump, 0, s.Line)
	c.emit(OpEndTry, 0, s.Line)

	// Patch try jump to catch block
	c.chunk.Code[tryJump].Operand = len(c.chunk.Code)

	c.beginScope()
	c.addLocal(s.ErrName) // error variable
	for _, stmt := range s.Catch {
		if err := c.compileStmt(stmt); err != nil {
			return err
		}
	}
	c.endScope()

	c.chunk.Code[endJump].Operand = len(c.chunk.Code)

	return nil
}

// StructDef represents a struct type definition at runtime.
type StructDef struct {
	Name   string
	Fields []string
}
