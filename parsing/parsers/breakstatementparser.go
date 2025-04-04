package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type BreakStatementParser struct {
	*ParserBase

	Stmt *statements.Statement
}

func NewBreakStatementParser() *BreakStatementParser {
	return &BreakStatementParser{
		ParserBase: NewParserBase(),
	}
}

func (p *BreakStatementParser) GetStatement() *statements.Statement {
	return p.Stmt
}

func (p *BreakStatementParser) Run(
	toks slang.TokenSeq,
	comment string,
) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenBREAK) {
		return false
	}

	toks.Advance()

	p.Stmt = statements.NewBreak()

	p.Stmt.SetComment(comment)

	return true
}
