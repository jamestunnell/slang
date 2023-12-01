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
	switch toks.Current().Type() {
	case slang.TokenCLASS:
		return p.ParseStatement(toks, NewClassStatementParser())
	case slang.TokenFIELD:
		return p.ParseStatement(toks, NewFieldStatementParser())
	case slang.TokenFUNC:
		return p.ParseStatement(toks, NewFuncStatementParser())
	case slang.TokenMETHOD:
		return p.ParseStatement(toks, NewMethodStatementParser())
	}

	p.TokenErr(
		toks.Current(), slang.TokenCLASS, slang.TokenFIELD, slang.TokenFUNC, slang.TokenMETHOD)

	return false
}
