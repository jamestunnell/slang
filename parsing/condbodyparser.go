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
		return p.ParseReturnStatment(toks)
	}

	return p.ParseExpressionOrAssignStatement(toks)
}
