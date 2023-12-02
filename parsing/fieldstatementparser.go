package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type FieldStatementParser struct {
	*ParserBase

	FieldStmt *statements.Field
}

func NewFieldStatementParser() *FieldStatementParser {
	return &FieldStatementParser{ParserBase: NewParserBase()}
}

func (p *FieldStatementParser) GetStatement() slang.Statement {
	return p.FieldStmt
}

func (p *FieldStatementParser) Run(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenFIELD) {
		return false
	}

	toks.Advance()

	name, typ, ok := p.ParseNameTypePair(toks)
	if !ok {
		return false
	}

	p.FieldStmt = statements.NewField(name, typ)

	return true
}
