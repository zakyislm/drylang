package compiler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

// compileBinary compiles binary operators and the ?? coalescing special case.
func (c *Compiler) compileBinary(e *ast.BinaryExpr) error {
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
	return nil
}

// compileCall compiles function calls, dispatching builtins to OpBuiltin.
// A user-defined function or global variable with a builtin's name wins over
// the builtin (e.g. `fn add(...)` must not resolve to BuiltinAdd).
func (c *Compiler) compileCall(e *ast.CallExpr) error {
	// Check for built-in function calls
	if ident, ok := e.Callee.(*ast.Ident); ok {
		if bid, found := core.BuiltinNames[ident.Name]; found && !c.globals[ident.Name] && c.resolveLocal(ident.Name) < 0 {
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
	return nil
}

// compileArrowFn compiles an anonymous arrow function into a closure.
func (c *Compiler) compileArrowFn(e *ast.ArrowFnExpr) error {
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
		Chunk:      fnCompiler.chunk,
		Name:       "<arrow>",
		ParamCount: len(e.Params),
	}
	c.fns = append(c.fns, fn)
	fi := c.addConst(fn)
	c.emit(core.OpClosure, fi, e.Line, e.Col)
	return nil
}
