package customerrs

import "fmt"

type ErrMethodUndefined struct {
	method, class string
}

func NewErrMethodUndefined(method string, class string) *ErrMethodUndefined {
	return &ErrMethodUndefined{
		method: method,
		class:  class,
	}
}

func (err *ErrMethodUndefined) Error() string {
	const strFmt = "method %s is not defined for class %s"

	return fmt.Sprintf(strFmt, err.method, err.class)
}
