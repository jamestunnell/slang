package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type ReturnStatementParser struct {
	*ParserBase

	ReturnStmt *statements.Return
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

	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	p.ReturnStmt = statements.NewReturn(exprParser.Expr)

	return true
}
