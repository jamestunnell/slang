package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type InterfaceStatementParser struct {
	*ParserBase

	InterfaceStmt *statements.Statement
}

func NewInterfaceStatementParser() *InterfaceStatementParser {
	return &InterfaceStatementParser{ParserBase: NewParserBase()}
}

func (p *InterfaceStatementParser) GetStatement() *statements.Statement {
	return p.InterfaceStmt
}

func (p *InterfaceStatementParser) Run(
	toks slang.TokenSeq,
	comment string,
) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenINTERFACE) {
		return false
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	ifcName := toks.Current().Value()

	toks.AdvanceSkip()

	if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
		return false
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

	specs := []*statements.FuncSpec{}

	for !toks.Current().Is(slang.TokenRBRACE) {
		if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
			return false
		}

		funcName := toks.Current().Value()

		toks.Advance()

		sigParser := NewFuncSignatureParser()
		if !p.RunSubParser(toks, sigParser) {
			return false
		}

		spec := &statements.FuncSpec{
			Name:    funcName,
			Inputs:  sigParser.Inputs,
			Outputs: sigParser.Outputs,
		}

		specs = append(specs, spec)
	}

	p.InterfaceStmt = statements.NewInterface(ifcName, specs...)

	p.InterfaceStmt.SetComment(comment)

	return true
}
