package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexing"
)

type ClassBodyParser struct {
	*BodyParserBase
}

func NewClassBodyParser() *ClassBodyParser {
	p := &ClassBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.parseStatement)

	return p
}

func (p *ClassBodyParser) parseStatement(toks lexing.TokenSeq) slang.Statement {
	var sp StatementParser

	switch toks.Current().Type {
	case lexing.TokenCLASS:
		sp = NewClassStatementParser()
	case lexing.TokenCONST:
		sp = NewConstStatementParser()
	case lexing.TokenFIELD:
		sp = NewFieldParser()
	case lexing.TokenFUNC:
		sp = NewFuncStatementParser()
	case lexing.TokenMETHOD:
		sp = NewMethodStatementParser()
	case lexing.TokenVAR:
		sp = NewVarStatementParser()
	default:
		p.TokenErr(
			toks.Current(), lexing.TokenCLASS, lexing.TokenFIELD, lexing.TokenFUNC, lexing.TokenMETHOD, lexing.TokenVAR)

		return nil
	}

	return p.ParseStatement(toks, sp)
}
