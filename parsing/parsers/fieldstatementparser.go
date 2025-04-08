package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type RecordFieldStatementParser struct {
	*ParserBase

	FieldStmt *statements.Statement
}

func NewFieldStatementParser() *RecordFieldStatementParser {
	return &RecordFieldStatementParser{ParserBase: NewParserBase()}
}

func (p *RecordFieldStatementParser) GetStatement() *statements.Statement {
	return p.FieldStmt
}

func (p *RecordFieldStatementParser) Run(toks slang.TokenSeq, comment string) bool {
	names, typ, ok := p.ParseNamesType(toks)
	if !ok {
		return false
	}

	p.FieldStmt = statements.NewField(typ, names...)

	p.FieldStmt.SetComment(comment)

	return true
}
