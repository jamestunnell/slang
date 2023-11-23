package parsing

import "github.com/jamestunnell/slang"

type ParseErr struct {
	Error error
	Token *slang.Token
}

func NewParseError(err error, tok *slang.Token) *ParseErr {
	return &ParseErr{
		Error: err,
		Token: tok,
	}
}
