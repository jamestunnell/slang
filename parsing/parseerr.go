package parsing

import (
	"fmt"

	"github.com/jamestunnell/slang"
)

type ParseErr struct {
	Err   error
	Token *slang.Token
}

func NewParseError(err error, tok *slang.Token) *ParseErr {
	return &ParseErr{
		Err:   err,
		Token: tok,
	}
}

func (err *ParseErr) Error() string {
	return fmt.Sprintf("%s: %v\n", err.Token.Location, err.Err)
}
