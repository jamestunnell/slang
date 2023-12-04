package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type ReturnStatementParser struct {
	*ParserBase

	ReturnStmt slang.Statement
}

func NewReturnStatementParser() *ReturnStatementParser {
	return &ReturnStatementParser{ParserBase: NewParserBase()}
}

func (p *ReturnStatementParser) GetStatement() slang.Statement {
	return p.ReturnStmt
}

func (p *ReturnStatementParser) Run(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenRETURN) {
		return false
	}

	toks.Advance()

	if toks.Current().Is(slang.TokenNEWLINE, slang.TokenRBRACE) {
		p.ReturnStmt = statements.NewReturn()

		return true
	}

	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	p.ReturnStmt = statements.NewReturnVal(exprParser.Expr)

	return true
}
