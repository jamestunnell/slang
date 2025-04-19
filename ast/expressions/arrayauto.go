package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type ArrayAuto struct {
	Values []*Expression `json:"values"`
}

func NewArrayAuto(vals ...*Expression) *Expression {
	return NewExpression(slang.ExprARRAYAUTO, &ArrayAuto{Values: vals})
}

func (a *ArrayAuto) IsEqual(other Core) bool {
	a2, ok := other.(*ArrayAuto)
	if !ok {
		return false
	}

	return slices.EqualFunc(a.Values, a2.Values, expressionsEqual)
}

func (a *ArrayAuto) Render(level int, w slang.CodeWriter) {
	switch len(a.Values) {
	case 0:
		w.WriteString("[]")
	case 1:
		w.WriteString("[")

		a.Values[0].Render(level, w)

		w.WriteString("]")
	default:
		w.WriteString("[")

		subLevel := level + 1
		for _, val := range a.Values {
			w.WriteNewline()
			w.WriteIndent(subLevel)

			val.Render(subLevel, w)
		}

		w.WriteIndent(level)
		w.WriteString("]")
	}
}
