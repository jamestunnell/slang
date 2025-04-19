package expressions

import (
	"github.com/jamestunnell/slang"
)

type Identifier struct {
	Name string `json:"name"`
}

func NewIdentifier(name string) *Expression {
	return NewExpression(slang.ExprIDENTIFIER, &Identifier{
		Name: name,
	})
}

func (i *Identifier) IsEqual(other Core) bool {
	i2, ok := other.(*Identifier)
	if !ok {
		return false
	}

	return i2.Name == i.Name
}

func (i *Identifier) Render(level int, w slang.CodeWriter) {
	w.WriteString(i.Name)
}

// func (expr *Identifier) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, found := env.Get(expr.Name)
// 	if !found {
// 		obj, found = objects.FindBuiltInFn(expr.Name)

// 		if !found {
// 			return nil, customerrs.NewErrObjectNotFound(expr.Name)
// 		}
// 	}

// 	return obj, nil
// }
