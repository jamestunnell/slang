package slang

import (
	"fmt"
)

type Type interface {
	fmt.Stringer

	IsEqual(other Type) bool
	IsBool() bool
	IsArray() bool
	IsMap() bool
}
