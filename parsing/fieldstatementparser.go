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

func (p *FieldStatementParser) Run(toks slang.TokenSeq) {
	if !p.ExpectToken(toks.Current(), slang.TokenFIELD) {
		return
	}

	toks.Advance()

	name, typ, ok := p.ParseNameTypePair(toks)
	if !ok {
		return
	}

	p.FieldStmt = statements.NewField(name, typ)
}
