package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type BreakStatementParser struct {
	*ParserBase

	Stmt slang.Statement
}

func NewBreakStatementParser() *BreakStatementParser {
	return &BreakStatementParser{
		ParserBase: NewParserBase(),
	}
}

func (p *BreakStatementParser) GetStatement() slang.Statement {
	return p.Stmt
}

func (p *BreakStatementParser) Run(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenBREAK) {
		return false
	}

	toks.Advance()

	p.Stmt = statements.NewBreak()

	return true
}
