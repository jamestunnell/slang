package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/tokens"
)

type Break struct{}

func NewBreak() *Statement {
	return NewStatement(slang.StatementBREAK, &Break{})
}

func (b *Break) GetName() (string, bool) {
	return "", false
}

func (b *Break) IsEqual(other Core) bool {
	_, ok := other.(*Break)

	return ok
}

func (b *Break) Render(level int, w slang.CodeWriter) {
	w.WriteString(tokens.StrBREAK)
}

// func (expr *Break) Eval(env *slang.Environment) (slang.Object, error) {
// 	res, err := expr.Condition.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	if res.Truthy() {
// 		return expr.Consequence.Eval(slang.NewEnvironment(env))
// 	}

// 	return objects.NULL(), nil
// }
