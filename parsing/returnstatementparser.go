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

func (p *ReturnStatementParser) Run(toks slang.TokenSeq) {
	if !p.ExpectToken(toks.Current(), slang.TokenRETURN) {
		return
	}

	toks.Advance()

	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return
	}

	p.ReturnStmt = statements.NewReturn(exprParser.Expr)
}
