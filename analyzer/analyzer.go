package analyzer

import (
	"drylang/ast"
	"fmt"
	"sort"
)

// Analyze validates the AST for dryLang rules.
func Analyze(prog *ast.Program) error {
	a := &analyzer{
		declared: make(map[string]declInfo),
		used:     make(map[string]bool),
		inFn:     false,
		inLoop:   false,
		inAsync:  false,
		privates:       make(map[string]map[string]bool),
		privateClasses: make(map[string]bool),
		instances:      make(map[string]string),
		classMethods:   make(map[string]map[string]bool),
	}

	for _, stmt := range prog.Stmts {
		if err := a.analyzeStmt(stmt); err != nil {
			return err
		}
	}

	// Strict unused check — report the first unused declaration in source order.
	type unusedEntry struct {
		name string
		info declInfo
	}
	var unused []unusedEntry
	for name, info := range a.declared {
		if !a.used[name] && !info.isParam {
			unused = append(unused, unusedEntry{name: name, info: info})
		}
	}
	if len(unused) > 0 {
		sort.Slice(unused, func(i, j int) bool {
			if unused[i].info.line != unused[j].info.line {
				return unused[i].info.line < unused[j].info.line
			}
			return unused[i].info.col < unused[j].info.col
		})
		u := unused[0]
		return fmt.Errorf("%d:%d unused %s", u.info.line, u.info.col, u.name)
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

	// class tracking for pv enforcement
	currentClass   string
	privates       map[string]map[string]bool // className -> private field/method names
	privateClasses map[string]bool            // className is private (pv cl)
	instances      map[string]string          // variable name -> class name
	classMethods   map[string]map[string]bool // className -> method names
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

// analyzeStmt dispatches statement analysis by type.
func (a *analyzer) analyzeStmt(stmt ast.Stmt) error {
	switch s := stmt.(type) {
	case *ast.AssignStmt:
		return a.analyzeAssign(s)
	case *ast.ConstDeclStmt:
		return a.analyzeConstDecl(s)
	case *ast.UnknownBoolStmt:
		a.declare(s.Name, s.Line, s.Col, false, false)
		return nil
	case *ast.AwaitStmt:
		return nil
	case *ast.MulCallStmt:
		return a.analyzeExpr(s.Call)
	case *ast.ExprStmt:
		return a.analyzeExpr(s.Expression)
	case *ast.ReturnStmt:
		return a.analyzeReturn(s)
	case *ast.FnDeclStmt:
		return a.analyzeFnDecl(s)
	case *ast.IfStmt:
		return a.analyzeIf(s)
	case *ast.OnStmt:
		return a.analyzeOn(s)
	case *ast.LoopStmt:
		return a.analyzeLoop(s)
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
		return a.analyzeTry(s)
	case *ast.ThrowStmt:
		return a.analyzeExpr(s.Value)
	case *ast.UseStmt:
		return nil
	case *ast.PrivateStmt:
		return a.analyzeStmt(s.Inner)
	case *ast.StructDeclStmt:
		a.declare(s.Name, s.Line, s.Col, false, false)
		return nil
	case *ast.ClassStmt:
		return a.analyzeClass(s)
	case *ast.IndexAssignStmt:
		return a.analyzeIndexAssign(s)
	case *ast.DotAssignStmt:
		return a.analyzeDotAssign(s)
	}

	return nil
}
