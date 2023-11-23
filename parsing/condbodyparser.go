package parsing

import (
	"github.com/jamestunnell/slang"
)

type CondBodyParser struct {
	*BodyParser
}

func NewCondBodyParser() *CondBodyParser {
	p := &CondBodyParser{}

	p.BodyParser = NewBodyParser(p.parseStatement)

	return p
}

func (p *CondBodyParser) parseStatement(toks slang.TokenSeq) bool {
	switch toks.Current().Type() {
	case slang.TokenRETURN:
		return p.ParseReturn(toks)
	case slang.TokenSYMBOL:
		if toks.Next().Is(slang.TokenCOMMA, slang.TokenASSIGN) {
			return p.ParseAssign(toks)
		}
	}

	return p.ParseExpression(toks)
}
