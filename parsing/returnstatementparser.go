package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
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

func (p *ReturnStatementParser) Run(toks lexing.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), lexing.TokenRETURN) {
		return false
	}

	toks.Advance()

	if toks.Current().Is(lexing.TokenNEWLINE, lexing.TokenRBRACE) {
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
