package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
)

type Identifier struct {
	Name string
}

func NewIdentifier(name string) *Identifier {
	return &Identifier{Name: name}
}

func (i *Identifier) Type() slang.ExprType { return slang.ExprIDENTIFIER }

func (i *Identifier) Equal(other slang.Expression) bool {
	i2, ok := other.(*Identifier)
	if !ok {
		return false
	}

	return i2.Name == i.Name
}

func (expr *Identifier) Eval(env *slang.Environment) (slang.Object, error) {
	obj, found := env.Get(expr.Name)
	if !found {
		obj, found = objects.FindBuiltInFn(expr.Name)

		if !found {
			return nil, slang.NewErrObjectNotFound(expr.Name)
		}
	}

	return obj, nil
}
