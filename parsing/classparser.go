package parsing

import (
	"github.com/jamestunnell/slang"
)

type ClassParser struct {
	*BodyParserBase
}

func NewClassParser() *ClassParser {
	p := &ClassParser{}

	p.BodyParserBase = NewBodyParserBase(p.parseStatement)

	return p
}

func (p *ClassParser) parseStatement(toks slang.TokenSeq) bool {
	var sp StatementParser

	switch toks.Current().Type() {
	case slang.TokenCLASS:
		sp = NewClassStatementParser()
	case slang.TokenFIELD:
		sp = NewFieldParser()
	case slang.TokenFUNC:
		sp = NewFuncStatementParser()
	case slang.TokenMETHOD:
		sp = NewMethodStatementParser()
	default:
		p.TokenErr(
			toks.Current(), slang.TokenCLASS, slang.TokenFIELD, slang.TokenFUNC, slang.TokenMETHOD)

		return false
	}

	return p.ParseStatement(toks, sp)
}
