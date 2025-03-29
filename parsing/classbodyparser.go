package parsing

import (
	"github.com/jamestunnell/slang"
)

type ClassBodyParser struct {
	*BodyParserBase
}

func NewClassBodyParser() *ClassBodyParser {
	p := &ClassBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.makeStatementParser)

	return p
}

func (p *ClassBodyParser) makeStatementParser(
	cur *slang.Token,
) StatementParser {
	switch cur.Type() {
	case slang.TokenCLASS:
		return NewClassStatementParser()
	case slang.TokenCONST:
		return NewConstStatementParser()
	case slang.TokenFIELD:
		return NewClassFieldParser()
	case slang.TokenFUNC:
		return NewFuncStatementParser()
	case slang.TokenMETHOD:
		return NewMethodStatementParser()
	case slang.TokenVAR:
		return NewVarStatementParser()
	}

	p.TokenErr(
		cur, slang.TokenCLASS, slang.TokenFIELD, slang.TokenFUNC, slang.TokenMETHOD, slang.TokenVAR)

	return nil
}
