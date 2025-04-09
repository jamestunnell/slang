package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type VarStatementParser struct {
	*ParserBase

	VarStmt *statements.Statement
}

func NewVarStatementParser() *VarStatementParser {
	return &VarStatementParser{ParserBase: NewParserBase()}
}

func (p *VarStatementParser) GetStatement() *statements.Statement {
	return p.VarStmt
}

func (p *VarStatementParser) Run(toks slang.TokenSeq, comment string) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenVAR) {
		return false
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	toks.Advance()

	exprParser := NewExprParser(parsing.PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	p.VarStmt = statements.NewVar(name, exprParser.Expr)

	p.VarStmt.SetComment(comment)

	return true
}
