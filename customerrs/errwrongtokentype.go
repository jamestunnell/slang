package customerrs

import (
	"fmt"

	"github.com/akrennmair/slice"
	"github.com/jamestunnell/slang"
)

type ErrWrongTokenType struct {
	actualType    slang.TokenType
	expectedTypes []slang.TokenType
}

func NewErrWrongTokenType(
	actualType slang.TokenType,
	expectedTypes []slang.TokenType,
) *ErrWrongTokenType {
	return &ErrWrongTokenType{
		actualType:    actualType,
		expectedTypes: expectedTypes,
	}
}

func (err *ErrWrongTokenType) Error() string {
	const fmtStr = "token type %s did not match any expected types %s"
	expectedStr := slice.Map(err.expectedTypes, func(tokType slang.TokenType) string {
		return tokType.String()
	})

	return fmt.Sprintf(fmtStr, err.actualType.String(), expectedStr)
}
