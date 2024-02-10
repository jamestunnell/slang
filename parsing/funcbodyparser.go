package parsing

import (
	"github.com/jamestunnell/slang"
)

type FuncBodyParser struct {
	*BodyParserBase
}

func NewFuncBodyParser() *FuncBodyParser {
	p := &FuncBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.parseStatement)

	return p
}

func (p *FuncBodyParser) parseStatement(toks slang.TokenSeq) slang.Statement {
	var sp StatementParser

	switch toks.Current().Type() {
	case slang.TokenCONST:
		sp = NewConstStatementParser()
	case slang.TokenIF:
		sp = NewIfStatementParser()
	case slang.TokenRETURN:
		sp = NewReturnStatementParser()
	case slang.TokenVAR:
		sp = NewVarStatementParser()
	default:
		sp = NewExprOrAssignStatementParser()
	}

	return p.ParseStatement(toks, sp)
}
