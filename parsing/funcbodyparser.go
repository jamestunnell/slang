package parsing

import (
	"github.com/jamestunnell/slang"
)

type FuncBodyParser struct {
	*BodyParser
}

func NewFuncBodyParser() *FuncBodyParser {
	p := &FuncBodyParser{}

	p.BodyParser = NewBodyParser(p.parseStatement)

	return p
}

func (p *FuncBodyParser) parseStatement(toks slang.TokenSeq) bool {
	switch toks.Current().Type() {
	// case slang.TokenFUNC:
	// case slang.TokenCLASS:
	case slang.TokenRETURN:
		return p.ParseReturn(toks)
	case slang.TokenSYMBOL:
		if toks.Next().Is(slang.TokenCOMMA, slang.TokenASSIGN) {
			return p.ParseAssign(toks)
		}
	}

	return p.ParseExpression(toks)
}
