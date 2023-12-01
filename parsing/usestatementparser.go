package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type UseStatementParser struct {
	*ParserBase

	UseStmt *statements.Use
}

func NewUseStatementParser() *UseStatementParser {
	return &UseStatementParser{ParserBase: NewParserBase()}
}

func (p *UseStatementParser) GetStatement() slang.Statement {
	return p.UseStmt
}

func (p *UseStatementParser) Run(toks slang.TokenSeq) {
	if !p.ExpectToken(toks.Current(), slang.TokenUSE) {
		return
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), slang.TokenSTRING) {
		return
	}

	path := toks.Current().Value()

	toks.Advance()

	p.UseStmt = statements.NewUse(path)
}
