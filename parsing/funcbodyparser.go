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

func (p *FuncBodyParser) parseStatement(toks slang.TokenSeq) bool {
	switch toks.Current().Type() {
	case slang.TokenCLASS:
		return p.ParseStatement(toks, NewClassStatementParser())
	case slang.TokenFUNC:
		return p.ParseStatement(toks, NewFuncStatementParser())
	case slang.TokenRETURN:
		return p.ParseStatement(toks, NewReturnStatementParser())
	}

	return p.ParseStatement(toks, NewExprOrAssignStatementParser())
}
