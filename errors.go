package slang

import "fmt"

type ParseErr struct {
	Err   error
	Token *Token
}

func NewParseError(err error, tok *Token) *ParseErr {
	return &ParseErr{
		Err:   err,
		Token: tok,
	}
}

func (err *ParseErr) Error() string {
	return fmt.Sprintf("%s: %v\n", err.Token.Location, err.Err)
}
