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

func (p *CondBodyParser) parseStatement(toks slang.TokenSeq) slang.Statement {
	var sp StatementParser

	switch toks.Current().Type() {
	case slang.TokenBREAK:
		sp = NewBreakStatementParser()
	case slang.TokenCONST:
		sp = NewConstStatementParser()
	case slang.TokenCONTINUE:
		sp = NewContinueStatementParser()
	case slang.TokenIF:
		sp = NewIfStatementParser()
	case slang.TokenFOREACH:
		sp = NewForEachStmtParser()
	case slang.TokenRETURN:
		sp = NewReturnStatementParser()
	case slang.TokenVAR:
		sp = NewVarStatementParser()
	default:
		sp = NewExprOrAssignStatementParser()
	}

	return p.ParseStatement(toks, sp)
}
