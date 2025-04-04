package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type ReturnStatementParser struct {
	*ParserBase

	ReturnStmt *statements.Statement
}

func NewReturnStatementParser() *ReturnStatementParser {
	return &ReturnStatementParser{ParserBase: NewParserBase()}
}

func (p *ReturnStatementParser) GetStatement() *statements.Statement {
	return p.ReturnStmt
}

func (p *ReturnStatementParser) Run(
	toks slang.TokenSeq,
	comment string,
) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenRETURN) {
		return false
	}

	toks.Advance()

	if toks.Current().Is(slang.TokenNEWLINE, slang.TokenRBRACE) {
		p.ReturnStmt = statements.NewReturn()

		p.ReturnStmt.SetComment(comment)

		return true
	}

	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	p.ReturnStmt = statements.NewReturnVal(exprParser.Expr)

	p.ReturnStmt.SetComment(comment)

	return true
}
