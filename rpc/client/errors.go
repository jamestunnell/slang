package client

import "fmt"

func newErrMethodFailed(method string, err error) *errMethodFailed {
	return &errMethodFailed{
		method: method,
		err:    err,
	}
}

type errMethodFailed struct {
	method string
	err    error
}

func (err *errMethodFailed) Error() string {
	return fmt.Sprintf("method %s failed: %v", err.method, err.err)
}
