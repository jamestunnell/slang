package compilation

import (
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"golang.org/x/exp/maps"
)

type Base struct {
	scope    string
	stmts    StmtSeq
	symbols  map[string]Symbol
	parent   Compiler
	children map[string]Compiler
}

func NewBase(scope string, stmts StmtSeq, parent Compiler) *Base {
	return &Base{
		scope:    scope,
		stmts:    stmts,
		symbols:  map[string]Symbol{},
		parent:   parent,
		children: map[string]Compiler{},
	}
}

func (base *Base) CollectSymbols() []Symbol {
	symbols := maps.Values(base.symbols)

	for _, child := range base.children {
		symbols = append(symbols, child.CollectSymbols()...)
	}

	return symbols
}

func (base *Base) scopedName(name string) string {
	if base.scope == "" {
		return name
	}

	return base.scope + "." + name
}

func (base *Base) handleClassStmtFirstPass(stmt slang.Statement, parent Compiler) {
	classStmt := stmt.(*statements.Class)
	name := classStmt.Name
	childStmts := NewStmtSeq(classStmt.Statements)
	child := NewClassCompiler(base.scopedName(name), childStmts, parent)

	base.symbols[name] = NewSymbol(base.scope, name, SymbolCLASS)
	base.children[name] = child
}

func (base *Base) handleFuncStmtFirstPass(stmt slang.Statement, parent Compiler) {
	funcStmt := stmt.(*statements.Func)
	name := funcStmt.Name
	childStmts := NewStmtSeq(funcStmt.Function.Statements)
	child := NewFuncCompiler(base.scopedName(name), childStmts, parent)

	base.symbols[name] = NewSymbol(base.scope, name, SymbolFUNC)
	base.children[name] = child
}

func (base *Base) runChildFirstPasses() error {
	for name, child := range base.children {
		if err := child.FirstPass(); err != nil {
			return fmt.Errorf("child %s first pass failed: %w", name, err)
		}
	}

	return nil
}
