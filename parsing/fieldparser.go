package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type FieldParser struct {
	*ParserBase

	FieldStmt *statements.Field
}

func NewFieldParser() *FieldParser {
	return &FieldParser{ParserBase: NewParserBase()}
}

func (p *FieldParser) GetStatement() slang.Statement {
	return p.FieldStmt
}

func (p *FieldParser) Run(toks slang.TokenSeq) bool {
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
