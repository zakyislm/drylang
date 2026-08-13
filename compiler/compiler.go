package compiler

import (
	"drylang/ast"
	"drylang/core"
)

// Compiler compiles AST to bytecode.
type Compiler struct {
	chunk     *core.Chunk
	locals    []local
	maxLocals int
	depth     int
	globals   map[string]bool
	fns       []*core.CompiledFn
	loopCtx   []loopContext
	slotNames map[int]string
}

type local struct {
	name  string
	depth int
}

type loopContext struct {
	start     int
	breaks    []int // jump indices to patch to after-loop
	continues []int // jump indices to patch to increment/loop-top
}

// New creates a new compiler.
func New() *Compiler {
	return &Compiler{
		chunk:     &core.Chunk{},
		globals:   make(map[string]bool),
		slotNames: make(map[int]string),
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
	slot := len(c.locals) - 1
	c.slotNames[slot] = name
	return slot
}

func (c *Compiler) beginScope() {
	c.depth++
}

func (c *Compiler) endScope() {
	for len(c.locals) > 0 && c.locals[len(c.locals)-1].depth == c.depth {
		c.locals = c.locals[:len(c.locals)-1]
	}
	c.depth--
}

// compileStmt dispatches statement compilation by type.
func (c *Compiler) compileStmt(stmt ast.Stmt) error {
	switch s := stmt.(type) {
	case *ast.AssignStmt:
		return c.compileAssign(s)
	case *ast.ConstDeclStmt:
		return c.compileConstDecl(s)
	case *ast.ExprStmt:
		return c.compileExprStmt(s)
	case *ast.UnknownBoolStmt:
		return c.compileUnknownBool(s)
	case *ast.DestructArrayStmt:
		return c.compileDestructArray(s)
	case *ast.DestructMapStmt:
		return c.compileDestructMap(s)
	case *ast.ReturnStmt:
		return c.compileReturn(s)
	case *ast.FnDeclStmt:
		return c.compileFnDecl(s)
	case *ast.IfStmt:
		return c.compileIf(s)
	case *ast.OnStmt:
		return c.compileOn(s)
	case *ast.MulCallStmt:
		return c.compileMulCall(s)
	case *ast.AwaitStmt:
		c.emit(core.OpAwait, 0, s.Line, s.Col)
		return nil
	case *ast.LoopStmt:
		return c.compileLoop(s)
	case *ast.DoneStmt:
		return c.compileDone(s)
	case *ast.ConStmt:
		return c.compileCon(s)
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
		return c.compileStructDecl(s)
	case *ast.ClassStmt:
		return c.compileClass(s)
	case *ast.IndexAssignStmt:
		return c.compileIndexAssign(s)
	case *ast.DotAssignStmt:
		return c.compileDotAssign(s)
	}

	return nil
}
