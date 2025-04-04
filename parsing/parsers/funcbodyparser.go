package parsers

import (
	"github.com/jamestunnell/slang"
)

type FuncBodyParser struct {
	*BodyParserBase
}

func NewFuncBodyParser() *FuncBodyParser {
	p := &FuncBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.makeStatementParser)

	return p
}

func (p *FuncBodyParser) makeStatementParser(
	cur *slang.Token,
) StatementParser {
	switch cur.Type() {
	case slang.TokenCONST:
		return NewConstStatementParser()
	case slang.TokenIF:
		return NewIfStatementParser()
	case slang.TokenFOREACH:
		return NewForEachStmtParser()
	case slang.TokenRETURN:
		return NewReturnStatementParser()
	case slang.TokenVAR:
		return NewVarStatementParser()
	}

	return NewExprOrAssignStatementParser()
}
