package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/statements"
)

type MethodStatementParser struct {
	*ParserBase

	MethodStmt *statements.Method
}

func NewMethodStatementParser() *MethodStatementParser {
	return &MethodStatementParser{ParserBase: NewParserBase()}
}

func (p *MethodStatementParser) GetStatement() slang.Statement {
	return p.MethodStmt
}

func (p *MethodStatementParser) Run(toks slang.TokenSeq) {
	if !p.ExpectToken(toks.Current(), slang.TokenMETHOD) {
		return
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return
	}

	name := toks.Current().Value()

	toks.Advance()

	sigParser := NewFuncSignatureParser()
	if !p.RunSubParser(toks, sigParser) {
		return
	}

	bodyParser := NewFuncBodyParser()
	if !p.RunSubParser(toks, bodyParser) {
		return
	}

	fn := ast.NewFunction(
		sigParser.Params, sigParser.ReturnTypes, bodyParser.Statements...)

	p.MethodStmt = statements.NewMethod(name, fn)
}
