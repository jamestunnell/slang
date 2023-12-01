package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/statements"
)

type FuncStatementParser struct {
	*ParserBase

	FuncStmt *statements.Func
}

func NewFuncStatementParser() *FuncStatementParser {
	return &FuncStatementParser{
		ParserBase: NewParserBase(),
	}
}

func (p *FuncStatementParser) GetStatement() slang.Statement {
	return p.FuncStmt
}

func (p *FuncStatementParser) Run(toks slang.TokenSeq) {
	if !p.ExpectToken(toks.Current(), slang.TokenFUNC) {
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
		sigParser.Params, sigParser.ReturnTypes, bodyParser.GetStatements()...)

	p.FuncStmt = statements.NewFunc(name, fn)
}
