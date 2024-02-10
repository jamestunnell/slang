package expressions

import (
	"github.com/jamestunnell/slang"
)

func NewString(val string) *Const[string] {
	return NewConst(slang.ExprSTRING, val)
}
