package slang

import (
	"fmt"
)

type Type interface {
	fmt.Stringer

	// IsEqual(other Type) bool
	// IsArray() bool
	// IsMap() bool
}

const (
	TypeBOOL = "BOOL"
	TypeERR  = "ERR"
	TypeFLT  = "FLT"
	TypeINT  = "INT"
	TypeSTR  = "STR"
)

func TypesEqual(a, b Type) bool {
	return a.String() == b.String()
}
