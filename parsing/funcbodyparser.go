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
		return p.ParseReturnStatment(toks)
	}

	return p.ParseExpressionOrAssignStatement(toks)
}
