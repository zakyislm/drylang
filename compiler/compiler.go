package compiler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/errfmt"
	"drylang/lexer"
	"drylang/compiler/varhandler"
	"drylang/compiler/exprhandler"
	"fmt"
	"strconv"
)

// Compiler compiles AST to bytecode.
type Compiler struct {
	chunk     *core.Chunk
	locals    []local
	depth     int
	maxLocals int
	fns       []*core.CompiledFn
	loopCtx   []loopContext
	globals   map[string]bool
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
		chunk:   &core.Chunk{},
		globals: make(map[string]bool),
	}
}

// Compile compiles a program AST to bytecode.
func (c *Compiler) Compile(prog *ast.Program) (*core.Chunk, []*core.CompiledFn, error) {
	for _, stmt := range prog.Stmts {
		if err := c.compileStmt(stmt); err != nil {
			return nil, nil, err
		}
	}
	c.chunk.LocalsCount = c.maxLocals
	return c.chunk, c.fns, nil
}

func (c *Compiler) emit(op core.Opcode, operand int, line, col int) int {
	return c.chunk.Emit(op, operand, line, col)
}

func (c *Compiler) emit2(op core.Opcode, operand, operand2 int, line, col int) int {
	return c.chunk.Emit2(op, operand, operand2, line, col)
}

func (c *Compiler) addConst(val interface{}) int {
	return c.chunk.AddConst(val)
}

func (c *Compiler) errorf(code string, line, col int, format string, args ...interface{}) error {
	return errfmt.Format(code, line, col, fmt.Sprintf(format, args...))
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
	if len(c.locals) > c.maxLocals {
		c.maxLocals = len(c.locals)
	}
	return len(c.locals) - 1
}

func (c *Compiler) beginScope() {
	c.depth++
}

func (c *Compiler) endScope() {
	for len(c.locals) > 0 && c.locals[len(c.locals)-1].depth == c.depth {
		c.locals = c.locals[:len(c.locals)-1]
		c.emit(core.OpPop, 0, 0, 0)
	}
	c.depth--
}

func (c *Compiler) compileStmt(stmt ast.Stmt) error {
	switch s := stmt.(type) {
	case *ast.AssignStmt:
		return varhandler.CompileAssign(c, s)
	case *ast.ExprStmt:
		return exprhandler.CompileExprStmt(c, s)
	case *ast.UnknownBoolStmt:
		return varhandler.CompileUnknownBool(c, s)

	case *ast.DestructArrayStmt:
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		
		if c.depth > 0 {
			tmpSlot := c.addLocal("")
			c.emit(core.OpSetLocal, tmpSlot, s.Line, s.Col)

			for i, name := range s.Names {
				c.emit(core.OpGetLocal, tmpSlot, s.Line, s.Col)
				c.emit(core.OpConst, c.addConst(float64(i)), s.Line, s.Col)
				c.emit(core.OpIndex, 0, s.Line, s.Col)
				
				idx := c.resolveLocal(name)
				if idx >= 0 {
					c.emit(core.OpSetLocal, idx, s.Line, s.Col)
				} else {
					slot := c.addLocal(name)
					c.emit(core.OpSetLocal, slot, s.Line, s.Col)
				}
				c.emit(core.OpPop, 0, s.Line, s.Col)
			}
			c.locals = c.locals[:len(c.locals)-1]
		} else {
			ci := c.addConst("__destruct__")
			c.emit(core.OpSetGlobal, ci, s.Line, s.Col)

			for i, name := range s.Names {
				c.emit(core.OpGetGlobal, ci, s.Line, s.Col)
				c.emit(core.OpConst, c.addConst(float64(i)), s.Line, s.Col)
				c.emit(core.OpIndex, 0, s.Line, s.Col)
				
				c.globals[name] = true
				nameConst := c.addConst(name)
				c.emit(core.OpSetGlobal, nameConst, s.Line, s.Col)
				c.emit(core.OpPop, 0, s.Line, s.Col)
			}
		}
		c.emit(core.OpPop, 0, s.Line, s.Col)
		return nil

		case *ast.DestructMapStmt:
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		
		if c.depth > 0 {
			tmpSlot := c.addLocal("")
			c.emit(core.OpSetLocal, tmpSlot, s.Line, s.Col)

			for _, key := range s.Keys {
				c.emit(core.OpGetLocal, tmpSlot, s.Line, s.Col)
				c.emit(core.OpConst, c.addConst(key), s.Line, s.Col)
				c.emit(core.OpIndex, 0, s.Line, s.Col)
				
				idx := c.resolveLocal(key)
				if idx >= 0 {
					c.emit(core.OpSetLocal, idx, s.Line, s.Col)
				} else {
					slot := c.addLocal(key)
					c.emit(core.OpSetLocal, slot, s.Line, s.Col)
				}
			}
			c.locals = c.locals[:len(c.locals)-1]
		} else {
			ci := c.addConst("__destruct__")
			c.emit(core.OpSetGlobal, ci, s.Line, s.Col)

			for _, key := range s.Keys {
				c.emit(core.OpGetGlobal, ci, s.Line, s.Col)
				c.emit(core.OpConst, c.addConst(key), s.Line, s.Col)
				c.emit(core.OpIndex, 0, s.Line, s.Col)
				
				c.globals[key] = true
				nameConst := c.addConst(key)
				c.emit(core.OpSetGlobal, nameConst, s.Line, s.Col)
			}
		}
case *ast.ReturnStmt:
		if s.Value != nil {
			if err := c.compileExpr(s.Value); err != nil {
				return err
			}
		} else {
			c.emit(core.OpUnknown, 0, s.Line, s.Col) // default return unknown
		}
		c.emit(core.OpReturn, 0, s.Line, s.Col)
		return nil

	case *ast.FnDeclStmt:
		return c.compileFnDecl(s)

	case *ast.IfStmt:
		return c.compileIf(s)

	case *ast.OnStmt:
		return c.compileOn(s)

	case *ast.LoopStmt:
		return c.compileLoop(s)

	case *ast.DoneStmt:
		if len(c.loopCtx) == 0 {
			return c.errorf("E203", s.Line, s.Col, "stray done")
		}
		ctx := &c.loopCtx[len(c.loopCtx)-1]
		jmp := c.emit(core.OpJump, 0, s.Line, s.Col)
		ctx.breaks = append(ctx.breaks, jmp)
		return nil

	case *ast.ConStmt:
		if len(c.loopCtx) == 0 {
			return c.errorf("E204", s.Line, s.Col, "stray con")
		}
		ctx := &c.loopCtx[len(c.loopCtx)-1]
		c.emit(core.OpLoop, ctx.start, s.Line, s.Col)
		return nil

	case *ast.TryStmt:
		return c.compileTry(s)

	case *ast.ThrowStmt:
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		c.emit(core.OpThrow, 0, s.Line, s.Col)
		return nil

	case *ast.UseStmt:
		// Module loading handled externally
		return nil

	case *ast.PrivateStmt:
		return c.compileStmt(s.Inner)

	case *ast.StructDeclStmt:
		// Store struct definition as a constant for runtime
		ci := c.addConst(core.StructDef{Name: s.Name, Fields: s.Fields, Visibility: s.Visibility})
		nameIdx := c.addConst(s.Name)
		c.emit(core.OpConst, ci, s.Line, s.Col)
		c.globals[s.Name] = true
		c.emit(core.OpSetGlobal, nameIdx, s.Line, s.Col)
		return nil

	case *ast.ClassStmt:
		methods := make(map[string]core.ClassMethod)
		for _, m := range s.Methods {
			// Compile each method as a core.Chunk
			sub := New()
			sub.depth = 1 // Inside a method, we have local scope + `this`

			// Pre-declare `this` as local 0, which corresponds to the callee slot
			sub.addLocal("this")

			// Pre-declare parameters starting from local 1
			for _, param := range m.Params {
				sub.addLocal(param)
			}

			for _, bStmt := range m.Body {
				if err := sub.compileStmt(bStmt); err != nil {
					return err
				}
			}
			// Implicit return unknown if no return statement
			if len(m.Body) == 0 || !isReturn(m.Body[len(m.Body)-1]) {
				sub.emit(core.OpUnknown, 0, m.Line, m.Col)
				sub.emit(core.OpReturn, 0, m.Line, m.Col)
			}
			sub.chunk.LocalsCount = sub.maxLocals

			methods[m.Name] = core.ClassMethod{
				Chunk: sub.chunk,
				Visibility: m.Visibility,
				IsAsync:    m.IsAsync,
			}
		}

		ci := c.addConst(core.ClassDef{Name: s.Name, ParentNames: s.Extends, Fields: s.Fields, Methods: methods, Visibility: s.Visibility})
		nameIdx := c.addConst(s.Name)
		c.emit(core.OpConst, ci, s.Line, s.Col)
		c.globals[s.Name] = true
		c.emit(core.OpSetGlobal, nameIdx, s.Line, s.Col)
		return nil

	case *ast.IndexAssignStmt:
		if err := c.compileExpr(s.Object); err != nil {
			return err
		}
		if err := c.compileExpr(s.Index); err != nil {
			return err
		}
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		c.emit(core.OpSetIndex, 0, s.Line, s.Col)
		return nil

	case *ast.DotAssignStmt:
		if err := c.compileExpr(s.Object); err != nil {
			return err
		}
		fi := c.addConst(s.Field)
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
		c.emit(core.OpDotSet, fi, s.Line, s.Col)
		return nil
	}

	return nil
}

func (c *Compiler) compileExpr(expr ast.Expr) error {
	switch e := expr.(type) {
	case *ast.NumberLit:
		v, err := strconv.ParseFloat(e.Value, 64)
		if err != nil {
			return c.errorf("E110", e.Line, e.Col, "bad number")
		}
		ci := c.addConst(v)
		c.emit(core.OpConst, ci, e.Line, e.Col)

	case *ast.StringLit:
		ci := c.addConst(e.Value)
		c.emit(core.OpConst, ci, e.Line, e.Col)

	case *ast.RawStringLit:
		ci := c.addConst(e.Value)
		c.emit(core.OpConst, ci, e.Line, e.Col)

	case *ast.BoolLit:
		if e.Value {
			c.emit(core.OpTrue, 0, e.Line, e.Col)
		} else {
			c.emit(core.OpFalse, 0, e.Line, e.Col)
		}

	case *ast.UnknownLit:
		c.emit(core.OpUnknown, 0, e.Line, e.Col)

	case *ast.Ident:
		idx := c.resolveLocal(e.Name)
		if idx >= 0 {
			c.emit(core.OpGetLocal, idx, e.Line, e.Col)
		} else {
			ci := c.addConst(e.Name)
			c.emit(core.OpGetGlobal, ci, e.Line, e.Col)
		}

	case *ast.BinaryExpr:
		if e.Op == lexer.TOKEN_QQ {
			if err := c.compileExpr(e.Left); err != nil {
				return err
			}
			jump := c.emit(core.OpJumpIfNotUnknown, 0, e.Line, e.Col)
			c.emit(core.OpPop, 0, e.Line, e.Col) // pop the unknown value
			if err := c.compileExpr(e.Right); err != nil {
				return err
			}
			c.chunk.Code[jump].Operand = len(c.chunk.Code)
			return nil
		}

		if err := c.compileExpr(e.Left); err != nil {
			return err
		}
		if err := c.compileExpr(e.Right); err != nil {
			return err
		}
		switch e.Op {
		case lexer.TOKEN_PLUS:
			c.emit(core.OpAdd, 0, e.Line, e.Col)
		case lexer.TOKEN_MINUS:
			c.emit(core.OpSub, 0, e.Line, e.Col)
		case lexer.TOKEN_STAR:
			c.emit(core.OpMul, 0, e.Line, e.Col)
		case lexer.TOKEN_SLASH:
			c.emit(core.OpDiv, 0, e.Line, e.Col)
		case lexer.TOKEN_PERCENT:
			c.emit(core.OpMod, 0, e.Line, e.Col)
		case lexer.TOKEN_ASSIGN:
			c.emit(core.OpEqual, 0, e.Line, e.Col)
		case lexer.TOKEN_NOT_EQ:
			c.emit(core.OpNotEqual, 0, e.Line, e.Col)
		case lexer.TOKEN_LT:
			c.emit(core.OpLess, 0, e.Line, e.Col)
		case lexer.TOKEN_GT:
			c.emit(core.OpGreater, 0, e.Line, e.Col)
		case lexer.TOKEN_LT_EQ:
			c.emit(core.OpLessEq, 0, e.Line, e.Col)
		case lexer.TOKEN_GT_EQ:
			c.emit(core.OpGreaterEq, 0, e.Line, e.Col)
		case lexer.TOKEN_AND:
			c.emit(core.OpAnd, 0, e.Line, e.Col)
		case lexer.TOKEN_OR:
			c.emit(core.OpOr, 0, e.Line, e.Col)
		}

	case *ast.UnaryExpr:
		if err := c.compileExpr(e.Operand); err != nil {
			return err
		}
		switch e.Op {
		case lexer.TOKEN_MINUS:
			c.emit(core.OpNeg, 0, e.Line, e.Col)
		case lexer.TOKEN_NOT:
			c.emit(core.OpNot, 0, e.Line, e.Col)
		}

	case *ast.CallExpr:
		// Check for built-in function calls
		if ident, ok := e.Callee.(*ast.Ident); ok {
			if bid, found := core.BuiltinNames[ident.Name]; found {
				for _, arg := range e.Args {
					if err := c.compileExpr(arg); err != nil {
						return err
					}
				}
				c.emit2(core.OpBuiltin, int(bid), len(e.Args), e.Line, e.Col)
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
		c.emit(core.OpCall, len(e.Args), e.Line, e.Col)

	case *ast.IndexExpr:
		if err := c.compileExpr(e.Object); err != nil {
			return err
		}
		if err := c.compileExpr(e.Index); err != nil {
			return err
		}
		c.emit(core.OpIndex, 0, e.Line, e.Col)

	case *ast.DotExpr:
		if err := c.compileExpr(e.Object); err != nil {
			return err
		}
		fi := c.addConst(e.Field)
		if e.Optional {
			c.emit(core.OpOptDotGet, fi, e.Line, e.Col)
		} else {
			c.emit(core.OpDotGet, fi, e.Line, e.Col)
		}

	case *ast.ArrayLit:
		for _, item := range e.Items {
			if err := c.compileExpr(item); err != nil {
				return err
			}
		}
		c.emit(core.OpArray, len(e.Items), e.Line, e.Col)

	case *ast.MapLit:
		for i := range e.Keys {
			if err := c.compileExpr(e.Keys[i]); err != nil {
				return err
			}
			if err := c.compileExpr(e.Values[i]); err != nil {
				return err
			}
		}
		c.emit(core.OpMap, len(e.Keys), e.Line, e.Col)

	case *ast.ArrowFnExpr:
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
		fnCompiler.emit(core.OpUnknown, 0, e.Line, e.Col) // default return
		fnCompiler.emit(core.OpReturn, 0, e.Line, e.Col)
		fnCompiler.chunk.LocalsCount = fnCompiler.maxLocals

		fn := &core.CompiledFn{
			Chunk: fnCompiler.chunk,
			Name:       "<arrow>",
			ParamCount: len(e.Params),
		}
		c.fns = append(c.fns, fn)
		fi := c.addConst(fn)
		c.emit(core.OpClosure, fi, e.Line, e.Col)
	case *ast.StringInterp:
		for i, part := range e.Parts {
			if err := c.compileExpr(part); err != nil {
				return err
			}
			if i > 0 {
				c.emit(core.OpConcat, 0, e.Line, e.Col)
			}
		}

	case *ast.StructInitExpr:
		// Push field values and create struct instance
		ci := c.addConst(e.TypeName)
		c.emit(core.OpConst, ci, e.Line, e.Col)
		for fname, fval := range e.Fields {
			fni := c.addConst(fname)
			c.emit(core.OpConst, fni, e.Line, e.Col)
			if err := c.compileExpr(fval); err != nil {
				return err
			}
		}
		c.emit(core.OpMap, len(e.Fields), e.Line, e.Col)

	case *ast.AwaitExpr:
		if err := c.compileExpr(e.Value); err != nil {
			return err
		}
		c.emit(core.OpAwait, 0, e.Line, e.Col)
	}

	return nil
}

func (c *Compiler) compileFnDecl(s *ast.FnDeclStmt) error {
	fnCompiler := New()
	fnCompiler.globals = c.globals
	fnCompiler.depth = 1
	for _, p := range s.Params {
		fnCompiler.addLocal(p)
	}
	for _, stmt := range s.Body {
		if err := fnCompiler.compileStmt(stmt); err != nil {
			return err
		}
	}
	fnCompiler.emit(core.OpUnknown, 0, s.Line, s.Col) // default return unknown
	fnCompiler.emit(core.OpReturn, 0, s.Line, s.Col)
	fnCompiler.chunk.LocalsCount = fnCompiler.maxLocals

	fn := &core.CompiledFn{
		Chunk: fnCompiler.chunk,
		Name:       s.Name,
		ParamCount: len(s.Params),
		IsAsync:    s.IsAsync,
	}

	c.fns = append(c.fns, fn)
	fi := c.addConst(fn)
	c.emit(core.OpClosure, fi, s.Line, s.Col)

	if c.depth > 0 {
		c.addLocal(s.Name)
	} else {
		c.globals[s.Name] = true
		ci := c.addConst(s.Name)
		c.emit(core.OpSetGlobal, ci, s.Line, s.Col)
	}

	return nil
}

func (c *Compiler) compileIf(s *ast.IfStmt) error {
	if err := c.compileExpr(s.Condition); err != nil {
		return err
	}

	// Jump to elif/else if false
	falseJump := c.emit(core.OpJumpIfFalse, 0, s.Line, s.Col)

	c.beginScope()
	for _, stmt := range s.Body {
		if err := c.compileStmt(stmt); err != nil {
			return err
		}
	}
	c.endScope()

	// Jump over elif/else blocks
	var endJumps []int
	endJump := c.emit(core.OpJump, 0, s.Line, s.Col)
	endJumps = append(endJumps, endJump)

	// Patch false jump
	c.chunk.Code[falseJump].Operand = len(c.chunk.Code)

	// Compile elif chains
	for _, elif := range s.ElIfs {
		if err := c.compileExpr(elif.Condition); err != nil {
			return err
		}
		elifFalseJump := c.emit(core.OpJumpIfFalse, 0, s.Line, s.Col)

		c.beginScope()
		for _, stmt := range elif.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()

		ej := c.emit(core.OpJump, 0, s.Line, s.Col)
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

func (c *Compiler) compileOn(s *ast.OnStmt) error {
	if err := c.compileExpr(s.Value); err != nil {
		return err
	}

	var endJumps []int

	for _, cas := range s.Cases {
		// Duplicate the switch value
		c.emit(core.OpConst, c.addConst("__dup__"), s.Line, s.Col) // placeholder
		if err := c.compileExpr(cas.Value); err != nil {
			return err
		}
		c.emit(core.OpEqual, 0, s.Line, s.Col)
		falseJump := c.emit(core.OpJumpIfFalse, 0, s.Line, s.Col)

		c.beginScope()
		for _, stmt := range cas.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()

		ej := c.emit(core.OpJump, 0, s.Line, s.Col)
		endJumps = append(endJumps, ej)
		c.chunk.Code[falseJump].Operand = len(c.chunk.Code)
	}

	for _, ej := range endJumps {
		c.chunk.Code[ej].Operand = len(c.chunk.Code)
	}

	c.emit(core.OpPop, 0, s.Line, s.Col) // pop switch value

	return nil
}

func (c *Compiler) compileLoop(s *ast.LoopStmt) error {
	loopStart := len(c.chunk.Code)
	c.loopCtx = append(c.loopCtx, loopContext{start: loopStart})

	if s.Limit != nil {
		// Counted loop: emit counter init and check
		ci := c.addConst(float64(0))
		c.emit(core.OpConst, ci, s.Line, s.Col) // push 0 as counter
		iSlot := c.addLocal("i")
		c.emit(core.OpSetLocal, iSlot, s.Line, s.Col)

		loopStart = len(c.chunk.Code) // actual loop start
		c.loopCtx[len(c.loopCtx)-1].start = loopStart

		// Check: i < limit
		idx := c.resolveLocal("i")
		c.emit(core.OpGetLocal, idx, s.Line, s.Col)
		if err := c.compileExpr(s.Limit); err != nil {
			return err
		}
		c.emit(core.OpLess, 0, s.Line, s.Col)
		exitJump := c.emit(core.OpJumpIfFalse, 0, s.Line, s.Col)

		c.beginScope()
		for _, stmt := range s.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()

		// Increment i
		c.emit(core.OpGetLocal, idx, s.Line, s.Col)
		c.emit(core.OpConst, c.addConst(float64(1)), s.Line, s.Col)
		c.emit(core.OpAdd, 0, s.Line, s.Col)
		c.emit(core.OpSetLocal, idx, s.Line, s.Col)

		c.emit(core.OpLoop, len(c.chunk.Code)+1-loopStart, s.Line, s.Col)
		c.chunk.Code[exitJump].Operand = len(c.chunk.Code)

		// Remove counter local
		c.locals = c.locals[:len(c.locals)-1]
		c.emit(core.OpPop, 0, s.Line, s.Col)
	} else {
		// Infinite loop
		c.beginScope()
		for _, stmt := range s.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()
		c.emit(core.OpLoop, len(c.chunk.Code)+1-loopStart, s.Line, s.Col)
	}

	// Patch breaks
	ctx := c.loopCtx[len(c.loopCtx)-1]
	for _, b := range ctx.breaks {
		c.chunk.Code[b].Operand = len(c.chunk.Code)
	}
	c.loopCtx = c.loopCtx[:len(c.loopCtx)-1]

	return nil
}

func (c *Compiler) compileTry(s *ast.TryStmt) error {
	tryJump := c.emit(core.OpTry, 0, s.Line, s.Col)

	c.beginScope()
	for _, stmt := range s.Body {
		if err := c.compileStmt(stmt); err != nil {
			return err
		}
	}
	c.endScope()

	c.emit(core.OpEndTry, 0, s.Line, s.Col)
	endJump := c.emit(core.OpJump, 0, s.Line, s.Col)

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

// core.StructDef represents a struct type definition at runtime.

// core.ClassDef represents a class type definition at runtime.

func isReturn(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.ReturnStmt)
	return ok
}
