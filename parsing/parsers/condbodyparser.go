package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/jamestunnell/slang/parsing"
)

func NewCondBodyParser() BodyParser {
	return NewBodyParser(
		func(toks parsing.TokenSeq) error {
			if !toks.Current().Is(slang.TokenLBRACE) {
				return customerrs.NewErrWrongTokenType(toks.Current(), slang.TokenLBRACE)
			}

			toks.Advance()

			return nil
		},
		slang.TokenRBRACE,
		MakeCondBodyStmtParser,
	)
}

func MakeCondBodyStmtParser(
	cur *slang.Token,
) (StatementParser, error) {
	var sp StatementParser

	switch cur.Type() {
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

	return sp, nil
}
