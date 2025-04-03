package expressions

import (
	"github.com/jamestunnell/slang"
)

func NewStr(val string) *Const[string] {
	return NewConst(slang.ExprSTR, val)
}
