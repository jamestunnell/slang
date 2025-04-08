package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
	"golang.org/x/exp/slices"
)

type Field struct {
	Names []string
	Type  *types.Type
}

func NewField(typ *types.Type, names ...string) *Statement {
	core := &Field{
		Names: names,
		Type:  typ,
	}

	return NewStatement(slang.StatementFIELD, core)
}

func (f *Field) IsEqual(other Core) bool {
	f2, ok := other.(*Field)
	if !ok {
		return false
	}

	if !slices.Equal(f.Names, f2.Names) {
		return false
	}

	return f.Type.IsEqual(f2.Type)
}

// func (expr *Continue) Eval(env *slang.Environment) (slang.Object, error) {
// 	res, err := expr.Condition.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	if res.Truthy() {
// 		return expr.Consequence.Eval(slang.NewEnvironment(env))
// 	}

// 	return objects.NULL(), nil
// }
