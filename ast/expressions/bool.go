package expressions

import (
	"github.com/jamestunnell/slang"
)

func NewBool(val bool) *Const[bool] {
	return NewConst(slang.ExprBOOL, val)
}
