package customerrs

import "fmt"

type ErrObjectType struct {
	expected, actual string
}

func NewErrObjectType(expected, actual string) *ErrObjectType {
	return &ErrObjectType{
		expected: expected,
		actual:   actual,
	}
}

func (err *ErrObjectType) Error() string {
	const strFmt = "unexpected object type: wanted %s, got %s"

	return fmt.Sprintf(strFmt, err.expected, err.actual)
}
