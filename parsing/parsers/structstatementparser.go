package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type StructStatementParser struct {
	*ParserBase

	StructStmt *statements.Statement
}

func NewStructStatementParser() *StructStatementParser {
	return &StructStatementParser{ParserBase: NewParserBase()}
}

func (p *StructStatementParser) GetStatement() *statements.Statement {
	return p.StructStmt
}

func (p *StructStatementParser) Run(
	toks parsing.TokenSeq,
	comment string,
) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenSTRUCT) {
		return false
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	toks.AdvanceSkip(slang.TokenNEWLINE)

	fieldSeqParser := NewFieldSeqParser()
	if !p.RunSubParser(toks, fieldSeqParser) {
		return false
	}

	p.StructStmt = statements.NewStruct(name, fieldSeqParser.GetFields()...)

	p.StructStmt.SetComment(comment)

	return true
}
