package parsing

import (
	"github.com/jamestunnell/slang"
)

type ClassBodyParser struct {
	*BodyParserBase
}

func NewClassBodyParser() *ClassBodyParser {
	p := &ClassBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.parseStatement)

	return p
}

func (p *ClassBodyParser) parseStatement(toks slang.TokenSeq) bool {
	var sp StatementParser

	switch toks.Current().Type() {
	case slang.TokenCLASS:
		sp = NewClassStatementParser()
	case slang.TokenFIELD:
		sp = NewFieldStatementParser()
	case slang.TokenFUNC:
		sp = NewFuncStatementParser()
	case slang.TokenMETHOD:
		sp = NewMethodStatementParser()
	default:
		p.TokenErr(
			toks.Current(), slang.TokenCLASS, slang.TokenFIELD, slang.TokenFUNC, slang.TokenMETHOD)
	}

	return p.ParseStatement(toks, sp)
}
