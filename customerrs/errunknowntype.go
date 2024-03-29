package customerrs

import "fmt"

type ErrUnknownType struct {
	Type string
}

func NewErrUnknownType(typ string) *ErrUnknownType {
	return &ErrUnknownType{Type: typ}
}

func (err *ErrUnknownType) Error() string {
	return fmt.Sprintf("unknown type %s", err.Type)
}
