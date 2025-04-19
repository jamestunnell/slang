package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
)

type Expression struct {
	Value *expressions.Expression `json:"value"`
}

func NewExpression(val *expressions.Expression) *Statement {
	core := &Expression{Value: val}

	return NewStatement(slang.StatementEXPRESSION, core)
}

func (e *Expression) GetName() (string, bool) {
	return "", false
}

func (e *Expression) IsEqual(other Core) bool {
	e2, ok := other.(*Expression)
	if !ok {
		return false
	}

	return e2.Value.IsEqual(e.Value)
}

func (e *Expression) Render(level int, w slang.CodeWriter) {
	e.Value.Render(level, w)
}

// func (st *Expression) Eval(env *slang.Environment) (slang.Object, error) {
// 	return st.Value.Eval(env)
// }
