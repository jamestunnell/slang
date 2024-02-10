package slang

import (
	"fmt"
)

type Type interface {
	fmt.Stringer

	IsEqual(other Type) bool
	IsArray() bool
	IsMap() bool
}
