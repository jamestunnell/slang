package slang

import "fmt"

type ErrObjectType struct {
	expected, actual fmt.Stringer
}

func NewErrObjectType(expected, actual fmt.Stringer) *ErrObjectType {
	return &ErrObjectType{
		expected: expected,
		actual:   actual,
	}
}

func (err *ErrObjectType) Error() string {
	const strFmt = "unexpected object type: wanted %s, got %s"

	return fmt.Sprintf(strFmt, err.expected, err.actual)
}
