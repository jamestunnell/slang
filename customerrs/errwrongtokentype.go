package customerrs

import (
	"fmt"

	"github.com/akrennmair/slice"
	"github.com/jamestunnell/slang"
)

type ErrWrongTokenType struct {
	token         *slang.Token
	expectedTypes []slang.TokenType
}

func NewErrWrongTokenType(
	tok *slang.Token,
	expectedTypes []slang.TokenType,
) *ErrWrongTokenType {
	return &ErrWrongTokenType{
		token:         tok,
		expectedTypes: expectedTypes,
	}
}

func (err *ErrWrongTokenType) Error() string {
	const fmtStr = "%s token %s did not match any expected types %s"
	expectedStr := slice.Map(err.expectedTypes, func(tokType slang.TokenType) string {
		return tokType.String()
	})

	return fmt.Sprintf(fmtStr, err.token.Type(), err.token.Value(), expectedStr)
}
