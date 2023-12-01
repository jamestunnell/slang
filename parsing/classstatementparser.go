package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type ClassStatementParser struct {
	*ParserBase

	ClassStmt *statements.Class
}

func NewClassStatementParser() *ClassStatementParser {
	return &ClassStatementParser{ParserBase: NewParserBase()}
}

func (p *ClassStatementParser) GetStatement() slang.Statement {
	return p.ClassStmt
}

func (p *ClassStatementParser) Run(toks slang.TokenSeq) {
	if !p.ExpectToken(toks.Current(), slang.TokenCLASS) {
		return
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return
	}

	name := toks.Current().Value()

	toks.Advance()

	classParser := NewClassBodyParser()
	if !p.RunSubParser(toks, classParser) {
		return
	}

	p.ClassStmt = statements.NewClass(name, "", classParser.Statements...)
}
