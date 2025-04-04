package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type ContinueStatementParser struct {
	*ParserBase

	Stmt *statements.Statement
}

func NewContinueStatementParser() *ContinueStatementParser {
	return &ContinueStatementParser{
		ParserBase: NewParserBase(),
	}
}

func (p *ContinueStatementParser) GetStatement() *statements.Statement {
	return p.Stmt
}

func (p *ContinueStatementParser) Run(
	toks slang.TokenSeq,
	comment string,
) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenCONTINUE) {
		return false
	}

	toks.Advance()

	p.Stmt = statements.NewContinue()

	p.Stmt.SetComment(comment)

	return true
}
