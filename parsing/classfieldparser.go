package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type ClassFieldParser struct {
	*ParserBase

	FieldStmt *statements.ClassField
}

func NewClassFieldParser() *ClassFieldParser {
	return &ClassFieldParser{ParserBase: NewParserBase()}
}

func (p *ClassFieldParser) GetStatement() slang.Statement {
	return p.FieldStmt
}

func (p *ClassFieldParser) Run(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenFIELD) {
		return false
	}

	toks.Advance()

	name, typ, ok := p.ParseNameTypePair(toks)
	if !ok {
		return false
	}

	p.FieldStmt = statements.NewClassField(name, typ)

	return true
}
