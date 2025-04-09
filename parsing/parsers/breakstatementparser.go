package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
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
	toks parsing.TokenSeq,
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
