package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type MapAuto struct {
	Keys   []*Expression `json:"keys"`
	Values []*Expression `json:"values"`
}

func NewMapAuto(
	keys []*Expression,
	vals []*Expression) *Expression {
	core := &MapAuto{Keys: keys, Values: vals}

	return NewExpression(slang.ExprMAPAUTO, core)
}

func (m *MapAuto) IsEqual(other Core) bool {
	m2, ok := other.(*MapAuto)
	if !ok {
		return false
	}

	if !slices.EqualFunc(m.Keys, m2.Keys, expressionsEqual) {
		return false
	}

	return slices.EqualFunc(m.Values, m2.Values, expressionsEqual)
}

func (m *MapAuto) Render(level int, w slang.CodeWriter) {
	switch len(m.Values) {
	case 0:
		w.WriteString("[]")
	case 1:
		w.WriteString("[")

		m.Keys[0].Render(level, w)
		w.WriteString(":")
		m.Values[0].Render(level, w)

		w.WriteString("]")
	default:
		w.WriteString("[")

		subLevel := level + 1
		for i, val := range m.Values {
			w.WriteNewline()
			w.WriteIndent(subLevel)

			m.Keys[i].Render(level, w)
			w.WriteString(":")
			val.Render(level, w)
		}

		w.WriteIndent(level)
		w.WriteString("]")
	}
}
