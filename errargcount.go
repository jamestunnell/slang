package slang

import "fmt"

type ErrArgCount struct {
	expected, actual int
}

func NewErrArgCount(expected, actual int) *ErrArgCount {
	return &ErrArgCount{
		expected: expected,
		actual:   actual,
	}
}

func (err *ErrArgCount) Error() string {
	const strFmt = "wrong number of arguments: wanted %d, got %d"

	return fmt.Sprintf(strFmt, err.expected, err.actual)
}
