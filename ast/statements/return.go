package statements

import (
	"github.com/jamestunnell/slang"
)

type Return struct {
	*Base
}

func NewReturn() *Return {
	return &Return{
		Base: NewBase(slang.StatementRETURN),
	}
}

func (r *Return) IsEqual(other slang.Statement) bool {
	_, ok := other.(*Return)

	return ok
}

func (st *Return) Eval(env slang.Environment) (slang.Objects, error) {
	return slang.Objects{}, nil
}
