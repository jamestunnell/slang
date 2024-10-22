package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
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

func (p *ClassStatementParser) Run(toks lexing.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), lexing.TokenCLASS) {
		return false
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), lexing.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	toks.Advance()

	classParser := NewClassBodyParser()
	if !p.RunSubParser(toks, classParser) {
		return false
	}

	p.ClassStmt = statements.NewClass(name, "", classParser.Statements...)

	return true
}
