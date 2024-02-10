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

func (p *ClassBodyParser) parseStatement(toks slang.TokenSeq) slang.Statement {
	var sp StatementParser

	switch toks.Current().Type() {
	case slang.TokenCLASS:
		sp = NewClassStatementParser()
	case slang.TokenCONST:
		sp = NewConstStatementParser()
	case slang.TokenFIELD:
		sp = NewFieldParser()
	case slang.TokenFUNC:
		sp = NewFuncStatementParser()
	case slang.TokenMETHOD:
		sp = NewMethodStatementParser()
	case slang.TokenVAR:
		sp = NewVarStatementParser()
	default:
		p.TokenErr(
			toks.Current(), slang.TokenCLASS, slang.TokenFIELD, slang.TokenFUNC, slang.TokenMETHOD, slang.TokenVAR)

		return nil
	}

	return p.ParseStatement(toks, sp)
}
