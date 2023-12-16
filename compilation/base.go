package compilation

import (
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type Base struct {
	symbol   *slang.Symbol
	stmts    StmtSeq
	parent   Compiler
	children map[string]Compiler
}

func NewBase(symbol *slang.Symbol, stmts StmtSeq, parent Compiler) *Base {
	return &Base{
		symbol:   symbol,
		stmts:    stmts,
		parent:   parent,
		children: map[string]Compiler{},
	}
}

func (base *Base) Symbol() *slang.Symbol {
	return base.symbol
}

func (base *Base) ChildSymbols() []*slang.Symbol {
	symbols := []*slang.Symbol{}

	for _, child := range base.children {
		symbols = append(symbols, child.Symbol())
		symbols = append(symbols, child.ChildSymbols()...)
	}

	return symbols
}

func (base *Base) handleClassStmtFirstPass(stmt slang.Statement, parent Compiler) {
	classStmt := stmt.(*statements.Class)
	name := classStmt.Name
	childSymbol := slang.NewChildSymbol(name, slang.SymbolCLASS, base.symbol)
	childStmts := NewStmtSeq(classStmt.Statements)

	base.children[name] = NewClassCompiler(childSymbol, childStmts, parent)
}

func (base *Base) handleFuncStmtFirstPass(stmt slang.Statement, parent Compiler) {
	funcStmt := stmt.(*statements.Func)
	name := funcStmt.Name
	childSymbol := slang.NewChildSymbol(name, slang.SymbolFUNC, base.symbol)
	childStmts := NewStmtSeq(funcStmt.Function.Statements)

	base.children[name] = NewFuncCompiler(childSymbol, childStmts, parent)
}

func (base *Base) runChildFirstPasses() error {
	for name, child := range base.children {
		if err := child.FirstPass(); err != nil {
			return fmt.Errorf("child %s first pass failed: %w", name, err)
		}
	}

	return nil
}
