package parsing

import "github.com/jamestunnell/slang/lexing"

type ParseErr struct {
	Error error
	Token *lexing.Token
}

func NewParseError(err error, tok *lexing.Token) *ParseErr {
	return &ParseErr{
		Error: err,
		Token: tok,
	}
}
