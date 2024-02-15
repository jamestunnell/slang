package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type Identifier struct {
	*Base

	Name string `json:"name"`
}

func NewIdentifier(name string) *Identifier {
	return &Identifier{
		Base: NewBase(slang.ExprIDENTIFIER),
		Name: name,
	}
}

func (i *Identifier) Equal(other slang.Expression) bool {
	i2, ok := other.(*Identifier)
	if !ok {
		return false
	}

	return i2.Name == i.Name
}

func (expr *Identifier) Eval(env slang.Environment) (slang.Object, error) {
	obj, found := env.Get(expr.Name)
	if !found {
		return nil, customerrs.NewErrObjectNotFound(expr.Name)
	}

	return obj, nil
}
