package parsing

import (
	"github.com/jamestunnell/slang"
)

type CondBodyParser struct {
	*BodyParserBase
}

func NewCondBodyParser() *CondBodyParser {
	p := &CondBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.makeStatementParser)

	return p
}

func (p *CondBodyParser) makeStatementParser(
	cur *slang.Token,
) StatementParser {
	switch cur.Type() {
	case slang.TokenBREAK:
		return NewBreakStatementParser()
	case slang.TokenCONST:
		return NewConstStatementParser()
	case slang.TokenCONTINUE:
		return NewContinueStatementParser()
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
