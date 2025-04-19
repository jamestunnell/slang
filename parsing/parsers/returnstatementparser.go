package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
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
	toks parsing.TokenSeq,
	comment string,
) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenRETURN) {
		return false
	}

	toks.Advance()

	p.ReturnStmt = statements.NewReturn()

	return true
}
