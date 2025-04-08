package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

func NewFileParser() BodyParser {
	return NewBodyParser(
		func(slang.TokenSeq) error { return nil },
		slang.TokenEOF,
		MakeFileBodyStmtParser,
	)
}

func MakeFileBodyStmtParser(cur *slang.Token) (StatementParser, error) {
	switch cur.Type() {
	case slang.TokenCONST:
		return NewConstStatementParser(), nil
	case slang.TokenFUNC:
		return NewFuncStatementParser(), nil
	case slang.TokenSTRUCT:
		return NewStructStatementParser(), nil
	case slang.TokenVAR:
		return NewVarStatementParser(), nil
	case slang.TokenUSE:
		return NewUseStatementParser(), nil
	}

	return nil, customerrs.NewErrWrongTokenType(
		cur, slang.TokenCONST, slang.TokenFUNC, slang.TokenSTRUCT, slang.TokenUSE, slang.TokenVAR)
}
