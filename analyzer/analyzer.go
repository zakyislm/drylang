package analyzer

import (
	"drylang/ast"
	"fmt"
)

// Analyze validates the AST for dryLang rules.
func Analyze(prog *ast.Program) error {
	a := &analyzer{
		declared: make(map[string]declInfo),
		used:     make(map[string]bool),
		inFn:     false,
		inLoop:   false,
		inAsync:  false,
	}

	for _, stmt := range prog.Stmts {
		if err := a.analyzeStmt(stmt); err != nil {
			return err
		}
	}

	// Strict unused check
	for name, info := range a.declared {
		if !a.used[name] && !info.isParam {
			return fmt.Errorf("%d:%d unused %s", info.line, info.col, name)
		}
	}

	return nil
}

type declInfo struct {
	line    int
	col     int
	isConst bool
	isParam bool
}

type analyzer struct {
	declared map[string]declInfo
	used     map[string]bool
	consts   map[string]bool
	inFn     bool
	inLoop   bool
	inAsync  bool
}

func (a *analyzer) declare(name string, line, col int, isConst, isParam bool) {
	a.declared[name] = declInfo{line: line, col: col, isConst: isConst, isParam: isParam}
}

func (a *analyzer) markUsed(name string) {
	a.used[name] = true
}

func (a *analyzer) errorf(line, col int, format string, args ...interface{}) error {
	return fmt.Errorf("%d:%d %s", line, col, fmt.Sprintf(format, args...))
}

func (a *analyzer) analyzeStmt(stmt ast.Stmt) error {
	switch s := stmt.(type) {
	case *ast.AssignStmt:
		a.declare(s.Name, s.Line, s.Col, false, false)
		return a.analyzeExpr(s.Value)

	case *ast.ConstDeclStmt:
		if info, exists := a.declared[s.Name]; exists && info.isConst {
			return a.errorf(s.Line, s.Col, "locked %s", s.Name)
		}
		a.declare(s.Name, s.Line, s.Col, true, false)
		return a.analyzeExpr(s.Value)

	case *ast.UnknownBoolStmt:
		a.declare(s.Name, s.Line, s.Col, false, false)
		return nil

	case *ast.PrintStmt:
		return a.analyzeExpr(s.Value)

	case *ast.ExprStmt:
		return a.analyzeExpr(s.Expression)

	case *ast.ReturnStmt:
		if !a.inFn {
			return a.errorf(s.Line, s.Col, "stray rev")
		}
		if s.Value != nil {
			return a.analyzeExpr(s.Value)
		}
		return nil

	case *ast.FnDeclStmt:
		a.declare(s.Name, s.Line, s.Col, false, false)
		return a.analyzeFnBody(s.Params, s.Body, s.IsAsync)

	case *ast.IfStmt:
		if err := a.analyzeExpr(s.Condition); err != nil {
			return err
		}
		for _, stmt := range s.Body {
			if err := a.analyzeStmt(stmt); err != nil {
				return err
			}
		}
		for _, elif := range s.ElIfs {
			if err := a.analyzeExpr(elif.Condition); err != nil {
				return err
			}
			for _, stmt := range elif.Body {
				if err := a.analyzeStmt(stmt); err != nil {
					return err
				}
			}
		}
		for _, stmt := range s.Else {
			if err := a.analyzeStmt(stmt); err != nil {
				return err
			}
		}
		return nil

	case *ast.OnStmt:
		if err := a.analyzeExpr(s.Value); err != nil {
			return err
		}
		for _, c := range s.Cases {
			if err := a.analyzeExpr(c.Value); err != nil {
				return err
			}
			for _, stmt := range c.Body {
				if err := a.analyzeStmt(stmt); err != nil {
					return err
				}
			}
		}
		return nil

	case *ast.LoopStmt:
		if s.Limit != nil {
			if err := a.analyzeExpr(s.Limit); err != nil {
				return err
			}
		}
		prev := a.inLoop
		a.inLoop = true
		// Loop counter 'i' is implicitly declared
		a.declare("i", s.Line, s.Col, false, true)
		for _, stmt := range s.Body {
			if err := a.analyzeStmt(stmt); err != nil {
				return err
			}
		}
		a.inLoop = prev
		return nil

	case *ast.DoneStmt:
		if !a.inLoop {
			return a.errorf(s.Line, s.Col, "stray done")
		}
		return nil

	case *ast.ConStmt:
		if !a.inLoop {
			return a.errorf(s.Line, s.Col, "stray con")
		}
		return nil

	case *ast.TryStmt:
		for _, stmt := range s.Body {
			if err := a.analyzeStmt(stmt); err != nil {
				return err
			}
		}
		a.declare(s.ErrName, s.Line, s.Col, false, true)
		for _, stmt := range s.Catch {
			if err := a.analyzeStmt(stmt); err != nil {
				return err
			}
		}
		return nil

	case *ast.ThrowStmt:
		return a.analyzeExpr(s.Value)

	case *ast.UseStmt:
		return nil

	case *ast.PrivateStmt:
		return a.analyzeStmt(s.Inner)

	case *ast.StructDeclStmt:
		a.declare(s.Name, s.Line, s.Col, false, false)
		return nil

	case *ast.IndexAssignStmt:
		if err := a.analyzeExpr(s.Object); err != nil {
			return err
		}
		if err := a.analyzeExpr(s.Index); err != nil {
			return err
		}
		return a.analyzeExpr(s.Value)

	case *ast.DotAssignStmt:
		if err := a.analyzeExpr(s.Object); err != nil {
			return err
		}
		return a.analyzeExpr(s.Value)
	}

	return nil
}

func (a *analyzer) analyzeExpr(expr ast.Expr) error {
	switch e := expr.(type) {
	case *ast.Ident:
		a.markUsed(e.Name)
	case *ast.BinaryExpr:
		if err := a.analyzeExpr(e.Left); err != nil {
			return err
		}
		return a.analyzeExpr(e.Right)
	case *ast.UnaryExpr:
		return a.analyzeExpr(e.Operand)
	case *ast.CallExpr:
		if err := a.analyzeExpr(e.Callee); err != nil {
			return err
		}
		for _, arg := range e.Args {
			if err := a.analyzeExpr(arg); err != nil {
				return err
			}
		}
	case *ast.IndexExpr:
		if err := a.analyzeExpr(e.Object); err != nil {
			return err
		}
		return a.analyzeExpr(e.Index)
	case *ast.DotExpr:
		return a.analyzeExpr(e.Object)
	case *ast.ArrayLit:
		for _, item := range e.Items {
			if err := a.analyzeExpr(item); err != nil {
				return err
			}
		}
	case *ast.MapLit:
		for i := range e.Keys {
			if err := a.analyzeExpr(e.Keys[i]); err != nil {
				return err
			}
			if err := a.analyzeExpr(e.Values[i]); err != nil {
				return err
			}
		}
	case *ast.ArrowFnExpr:
		return a.analyzeFnBody(e.Params, e.Body, false)
	case *ast.StringInterp:
		for _, part := range e.Parts {
			if err := a.analyzeExpr(part); err != nil {
				return err
			}
		}
	case *ast.InputExpr:
		if e.Prompt != nil {
			return a.analyzeExpr(e.Prompt)
		}
	case *ast.StructInitExpr:
		a.markUsed(e.TypeName)
		for _, v := range e.Fields {
			if err := a.analyzeExpr(v); err != nil {
				return err
			}
		}
	case *ast.AwaitExpr:
		if !a.inAsync {
			return a.errorf(e.Line, e.Col, "stray awt")
		}
		return a.analyzeExpr(e.Value)
	}
	return nil
}

func (a *analyzer) analyzeFnBody(params []string, body []ast.Stmt, isAsync bool) error {
	prevFn := a.inFn
	prevAsync := a.inAsync
	a.inFn = true
	a.inAsync = isAsync

	for _, p := range params {
		a.declare(p, 0, 0, false, true)
	}

	for _, stmt := range body {
		if err := a.analyzeStmt(stmt); err != nil {
			return err
		}
	}

	a.inFn = prevFn
	a.inAsync = prevAsync
	return nil
}
