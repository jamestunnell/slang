package customerrs

import (
	"fmt"

	"github.com/akrennmair/slice"
	"github.com/jamestunnell/slang/lexing"
)

type ErrWrongTokenType struct {
	token         *lexing.Token
	expectedTypes []lexing.TokenType
}

func NewErrWrongTokenType(
	tok *lexing.Token,
	expectedTypes ...lexing.TokenType,
) *ErrWrongTokenType {
	return &ErrWrongTokenType{
		token:         tok,
		expectedTypes: expectedTypes,
	}
}

func (err *ErrWrongTokenType) Error() string {
	const fmtStr = "%s token %s did not match any expected types %s"
	expectedStr := slice.Map(err.expectedTypes, func(tokType lexing.TokenType) string {
		return tokType.String()
	})

	return fmt.Sprintf(fmtStr, err.token.Type, err.token.Value, expectedStr)
}
