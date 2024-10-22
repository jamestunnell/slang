package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
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

func (p *BreakStatementParser) Run(toks lexing.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), lexing.TokenBREAK) {
		return false
	}

	toks.Advance()

	p.Stmt = statements.NewBreak()

	return true
}
