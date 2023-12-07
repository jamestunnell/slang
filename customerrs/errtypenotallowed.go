package customerrs

import "fmt"

type ErrTypeNotAllowed struct {
	Thing, Context string
}

func NewErrTypeNotAllowed(thing, context string) *ErrTypeNotAllowed {
	return &ErrTypeNotAllowed{
		Thing:   thing,
		Context: context,
	}
}

func (err *ErrTypeNotAllowed) Error() string {
	return fmt.Sprintf("%s is not allowed in a %s context", err.Thing, err.Context)
}
