package parser

import (
	"drylang/lexer"
	"testing"
)

func parse(t *testing.T, input string) *Program {
	t.Helper()
	lex := lexer.New(input)
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Fatalf("lex error: %v", err)
	}
	p := New(tokens)
	prog, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	return prog
}

func TestAssignWithEquals(t *testing.T) {
	prog := parse(t, `name = "Zaky"`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	assign, ok := prog.Stmts[0].(*AssignStmt)
	if !ok {
		t.Fatalf("want AssignStmt, got %T", prog.Stmts[0])
	}
	if assign.Name != "name" {
		t.Errorf("want name='name', got %q", assign.Name)
	}
}

func TestAssignWithoutEquals(t *testing.T) {
	prog := parse(t, `name "Zaky"`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	assign, ok := prog.Stmts[0].(*AssignStmt)
	if !ok {
		t.Fatalf("want AssignStmt, got %T", prog.Stmts[0])
	}
	if assign.Name != "name" {
		t.Errorf("want name='name', got %q", assign.Name)
	}
}

func TestAllCapsAutoConst(t *testing.T) {
	prog := parse(t, `MAXLIFE = 5`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	_, ok := prog.Stmts[0].(*ConstDeclStmt)
	if !ok {
		t.Fatalf("want ConstDeclStmt for ALL CAPS, got %T", prog.Stmts[0])
	}
}

func TestConstDecl(t *testing.T) {
	prog := parse(t, `cns pi 3,14`)

	lex := lexer.New(`cns pi 3,14`)
	tokens, _ := lex.Tokenize()
	_ = New(tokens) // just for coverage

	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	c, ok := prog.Stmts[0].(*ConstDeclStmt)
	if !ok {
		t.Fatalf("want ConstDeclStmt, got %T", prog.Stmts[0])
	}
	if c.Name != "pi" {
		t.Errorf("want name='pi', got %q", c.Name)
	}
}

func TestFnDecl(t *testing.T) {
	prog := parse(t, `fn greet(who) {
  rev "Hi " + who
}`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	fn, ok := prog.Stmts[0].(*FnDeclStmt)
	if !ok {
		t.Fatalf("want FnDeclStmt, got %T", prog.Stmts[0])
	}
	if fn.Name != "greet" {
		t.Errorf("want name='greet', got %q", fn.Name)
	}
	if len(fn.Params) != 1 || fn.Params[0] != "who" {
		t.Errorf("want params=[who], got %v", fn.Params)
	}
}

func TestIfElIfEl(t *testing.T) {
	prog := parse(t, `if x = 1 {
  pt "one"
} elif x = 2 {
  pt "two"
} el {
  pt "other"
}`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	ifStmt, ok := prog.Stmts[0].(*IfStmt)
	if !ok {
		t.Fatalf("want IfStmt, got %T", prog.Stmts[0])
	}
	if len(ifStmt.ElIfs) != 1 {
		t.Errorf("want 1 elif, got %d", len(ifStmt.ElIfs))
	}
	if len(ifStmt.Else) == 0 {
		t.Error("want else block")
	}
}

func TestLoop(t *testing.T) {
	prog := parse(t, `lp 5 {
  pt "loop"
}`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	lp, ok := prog.Stmts[0].(*LoopStmt)
	if !ok {
		t.Fatalf("want LoopStmt, got %T", prog.Stmts[0])
	}
	if lp.Limit == nil {
		t.Error("want limit, got nil")
	}
}

func TestOnSwitch(t *testing.T) {
	prog := parse(t, `on(day) {
  1 { pt "senin" }
  2 { pt "selasa" }
}`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	on, ok := prog.Stmts[0].(*OnStmt)
	if !ok {
		t.Fatalf("want OnStmt, got %T", prog.Stmts[0])
	}
	if len(on.Cases) != 2 {
		t.Errorf("want 2 cases, got %d", len(on.Cases))
	}
}

func TestTryErr(t *testing.T) {
	prog := parse(t, `try {
  err "boom"
} err(e) {
  pt e
}`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	tr, ok := prog.Stmts[0].(*TryStmt)
	if !ok {
		t.Fatalf("want TryStmt, got %T", prog.Stmts[0])
	}
	if tr.ErrName != "e" {
		t.Errorf("want err name='e', got %q", tr.ErrName)
	}
}

func TestStructDeclAndInit(t *testing.T) {
	prog := parse(t, `user { name age }
u user { name "Zaky" age 17 }`)
	if len(prog.Stmts) != 2 {
		t.Fatalf("want 2 stmts, got %d", len(prog.Stmts))
	}

	sd, ok := prog.Stmts[0].(*StructDeclStmt)
	if !ok {
		t.Fatalf("want StructDeclStmt, got %T", prog.Stmts[0])
	}
	if sd.Name != "user" || len(sd.Fields) != 2 {
		t.Errorf("want user with 2 fields, got %q with %d", sd.Name, len(sd.Fields))
	}

	assign, ok := prog.Stmts[1].(*AssignStmt)
	if !ok {
		t.Fatalf("want AssignStmt for struct init, got %T", prog.Stmts[1])
	}
	if assign.Name != "u" {
		t.Errorf("want var name='u', got %q", assign.Name)
	}
	_, isInit := assign.Value.(*StructInitExpr)
	if !isInit {
		t.Fatalf("want StructInitExpr, got %T", assign.Value)
	}
}

func TestUnknownBool(t *testing.T) {
	prog := parse(t, `? status`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	ub, ok := prog.Stmts[0].(*UnknownBoolStmt)
	if !ok {
		t.Fatalf("want UnknownBoolStmt, got %T", prog.Stmts[0])
	}
	if ub.Name != "status" {
		t.Errorf("want name='status', got %q", ub.Name)
	}
}

func TestArrayAndMap(t *testing.T) {
	prog := parse(t, `arr = [1, 2, 3]
m = {"key": "val"}`)
	if len(prog.Stmts) != 2 {
		t.Fatalf("want 2 stmts, got %d", len(prog.Stmts))
	}
}

func TestPrint(t *testing.T) {
	prog := parse(t, `pt "hello"`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	_, ok := prog.Stmts[0].(*PrintStmt)
	if !ok {
		t.Fatalf("want PrintStmt, got %T", prog.Stmts[0])
	}
}

func TestUse(t *testing.T) {
	prog := parse(t, `use "helpers"`)
	if len(prog.Stmts) != 1 {
		t.Fatalf("want 1 stmt, got %d", len(prog.Stmts))
	}
	u, ok := prog.Stmts[0].(*UseStmt)
	if !ok {
		t.Fatalf("want UseStmt, got %T", prog.Stmts[0])
	}
	if u.Path != "helpers" {
		t.Errorf("want path='helpers', got %q", u.Path)
	}
}
