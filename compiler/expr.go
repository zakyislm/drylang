package compiler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
	"strconv"
)

// compileExpr compiles a single expression.
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
		if err := c.compileBinary(e); err != nil {
			return err
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
		if err := c.compileCall(e); err != nil {
			return err
		}

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
		if err := c.compileArrowFn(e); err != nil {
			return err
		}

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
		for fname, fval := range e.Fields {
			fni := c.addConst(fname)
			c.emit(core.OpConst, fni, e.Line, e.Col)
			if err := c.compileExpr(fval); err != nil {
				return err
			}
		}
		// Also inject the __struct__ field
		c.emit(core.OpConst, c.addConst("__struct__"), e.Line, e.Col)
		c.emit(core.OpConst, c.addConst(e.TypeName), e.Line, e.Col)

		c.emit(core.OpMap, len(e.Fields)+1, e.Line, e.Col)
	}

	return nil
}
