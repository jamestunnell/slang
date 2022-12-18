package slang

import "fmt"

type ErrArgType struct {
	expected, actual fmt.Stringer
}

func NewErrArgType(expected, actual fmt.Stringer) *ErrArgType {
	return &ErrArgType{
		expected: expected,
		actual:   actual,
	}
}

func (err *ErrArgType) Error() string {
	const strFmt = "unexpected argument type: wanted %s, got %s"

	return fmt.Sprintf(strFmt, err.expected, err.actual)
}
