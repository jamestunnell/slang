package customerrs

import (
	"fmt"

	"github.com/jamestunnell/slang/lexing"
)

type ErrMissingPrefixParseFn struct {
	Type lexing.TokenType
}

func NewErrMissingPrefixParseFn(typ lexing.TokenType) *ErrMissingPrefixParseFn {
	return &ErrMissingPrefixParseFn{
		Type: typ,
	}
}

func (err *ErrMissingPrefixParseFn) Error() string {
	const strFmt = "missing prefix parse function for token type %s"

	return fmt.Sprintf(strFmt, err.Type)
}
