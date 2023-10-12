package customerrs

import "fmt"

type ErrObjectNotFound struct {
	name string
}

func NewErrObjectNotFound(name string) *ErrObjectNotFound {
	return &ErrObjectNotFound{
		name: name,
	}
}

func (err *ErrObjectNotFound) Error() string {
	return fmt.Sprintf("object %s was not found", err.name)
}
