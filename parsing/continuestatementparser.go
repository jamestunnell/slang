package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type ContinueStatementParser struct {
	*ParserBase

	Stmt slang.Statement
}

func NewContinueStatementParser() *ContinueStatementParser {
	return &ContinueStatementParser{
		ParserBase: NewParserBase(),
	}
}

func (p *ContinueStatementParser) GetStatement() slang.Statement {
	return p.Stmt
}

func (p *ContinueStatementParser) Run(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenCONTINUE) {
		return false
	}

	toks.Advance()

	p.Stmt = statements.NewContinue()

	return true
}
