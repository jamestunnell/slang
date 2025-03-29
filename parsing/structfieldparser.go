package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type StructFieldParser struct {
	*ParserBase

	FieldStmt *statements.StructField
}

func NewStructFieldParser() *StructFieldParser {
	return &StructFieldParser{ParserBase: NewParserBase()}
}

func (p *StructFieldParser) GetStatement() slang.Statement {
	return p.FieldStmt
}

func (p *StructFieldParser) Run(toks slang.TokenSeq) bool {
	names, typ, ok := p.ParseNamesType(toks)
	if !ok {
		return false
	}

	p.FieldStmt = statements.NewStructField(names, typ)

	return true
}
