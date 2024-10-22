package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
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

func (p *MethodStatementParser) Run(toks lexing.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), lexing.TokenMETHOD) {
		return false
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), lexing.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value

	toks.Advance()

	sigParser := NewFuncSignatureParser()
	if !p.RunSubParser(toks, sigParser) {
		return false
	}

	bodyParser := NewFuncBodyParser()
	if !p.RunSubParser(toks, bodyParser) {
		return false
	}

	fn := ast.NewFunction(
		sigParser.Params, sigParser.ReturnTypes, bodyParser.Statements...)

	p.MethodStmt = statements.NewMethod(name, fn)

	return true
}
