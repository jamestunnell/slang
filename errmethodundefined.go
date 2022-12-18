package slang

import "fmt"

type ErrMethodUndefined struct {
	methodName string
	objectType fmt.Stringer
}

func NewErrMethodUndefined(methodName string, objectType fmt.Stringer) *ErrMethodUndefined {
	return &ErrMethodUndefined{
		methodName: methodName,
		objectType: objectType,
	}
}

func (err *ErrMethodUndefined) Error() string {
	const strFmt = "method %s is not defined for object type %s"

	return fmt.Sprintf(strFmt, err.methodName, err.objectType)
}
