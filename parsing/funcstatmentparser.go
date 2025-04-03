package parsing

import (
	"github.com/jamestunnell/slang"
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

func (p *FuncStatementParser) Run(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenFUNC) {
		return false
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	toks.AdvanceSkip(slang.TokenNEWLINE)

	sigParser := NewFuncSignatureParser()
	if !p.RunSubParser(toks, sigParser) {
		return false
	}

	bodyParser := NewFuncBodyParser()
	if !p.RunSubParser(toks, bodyParser) {
		return false
	}

	p.FuncStmt = statements.NewFunc(
		name,
		sigParser.InParams,
		sigParser.OutParams,
		bodyParser.GetStatements()...,
	)

	return true
}
