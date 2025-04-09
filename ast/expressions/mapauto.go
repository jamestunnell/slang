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
