package parsing

import (
	"github.com/jamestunnell/slang"
)

type StructBodyParser struct {
	*BodyParserBase
}

func NewStructBodyParser() *StructBodyParser {
	p := &StructBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.parseStatement)

	return p
}

func (p *StructBodyParser) parseStatement(toks slang.TokenSeq) slang.Statement {
	return p.ParseStatement(toks, NewStructFieldParser())
}
