package parsing

import (
	"github.com/jamestunnell/slang"
)

type StructBodyParser struct {
	*BodyParserBase
}

func NewStructBodyParser() *StructBodyParser {
	p := &StructBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.makeStatementParser)

	return p
}

func (p *StructBodyParser) makeStatementParser(
	cur *slang.Token,
) StatementParser {
	return NewStructFieldParser()
}
