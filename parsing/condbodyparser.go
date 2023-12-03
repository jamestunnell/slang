package parsing

import (
	"github.com/jamestunnell/slang"
)

type CondBodyParser struct {
	*BodyParserBase
}

func NewCondBodyParser() *CondBodyParser {
	p := &CondBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.parseStatement)

	return p
}

func (p *CondBodyParser) parseStatement(toks slang.TokenSeq) bool {
	switch toks.Current().Type() {
	case slang.TokenIF:
		return p.ParseStatement(toks, NewIfStatementParser())
	case slang.TokenRETURN:
		return p.ParseStatement(toks, NewReturnStatementParser())
	}

	return p.ParseStatement(toks, NewExprOrAssignStatementParser())
}
