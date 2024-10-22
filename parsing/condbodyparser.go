package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexing"
)

type CondBodyParser struct {
	*BodyParserBase
}

func NewCondBodyParser() *CondBodyParser {
	p := &CondBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.parseStatement)

	return p
}

func (p *CondBodyParser) parseStatement(toks lexing.TokenSeq) slang.Statement {
	var sp StatementParser

	switch toks.Current().Type {
	case lexing.TokenBREAK:
		sp = NewBreakStatementParser()
	case lexing.TokenCONST:
		sp = NewConstStatementParser()
	case lexing.TokenCONTINUE:
		sp = NewContinueStatementParser()
	case lexing.TokenIF:
		sp = NewIfStatementParser()
	case lexing.TokenFOREACH:
		sp = NewForEachStmtParser()
	case lexing.TokenRETURN:
		sp = NewReturnStatementParser()
	case lexing.TokenVAR:
		sp = NewVarStatementParser()
	default:
		sp = NewExprOrAssignStatementParser()
	}

	return p.ParseStatement(toks, sp)
}
