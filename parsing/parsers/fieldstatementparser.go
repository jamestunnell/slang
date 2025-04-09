package parsers

import (
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
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

func (p *RecordFieldStatementParser) Run(toks parsing.TokenSeq, comment string) bool {
	names, typ, ok := p.ParseNamesType(toks)
	if !ok {
		return false
	}

	p.FieldStmt = statements.NewField(typ, names...)

	p.FieldStmt.SetComment(comment)

	return true
}
