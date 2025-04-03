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

	toks.AdvanceSkip(slang.TokenNEWLINE)

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	toks.AdvanceSkip(slang.TokenNEWLINE)

	sigParser := NewDataSignatureParser()
	if !p.RunSubParser(toks, sigParser) {
		return false
	}

	p.StructStmt = statements.NewStruct(name, sigParser.NameTypes...)

	return true
}
