package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type StructStatementParser struct {
	*ParserBase

	StructStmt *statements.Struct
}

func NewStructStatementParser() *StructStatementParser {
	return &StructStatementParser{ParserBase: NewParserBase()}
}

func (p *StructStatementParser) GetStatement() slang.Statement {
	return p.StructStmt
}

func (p *StructStatementParser) Run(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenSTRUCT) {
		return false
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	toks.Advance()

	structParser := NewStructBodyParser()
	if !p.RunSubParser(toks, structParser) {
		return false
	}

	p.StructStmt = statements.NewStruct(name, structParser.Statements...)

	return true
}
