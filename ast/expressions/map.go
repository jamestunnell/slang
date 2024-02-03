package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Map struct {
	*Base

	Keys   []slang.Expression `json:"keys"`
	Values []slang.Expression `json:"values"`
}

func NewMap(keys, vals []slang.Expression) *Map {
	return &Map{
		Base:   NewBase(slang.ExprMAP),
		Keys:   keys,
		Values: vals,
	}
}

func (m *Map) Equal(other slang.Expression) bool {
	m2, ok := other.(*Map)
	if !ok {
		return false
	}

	if !slices.EqualFunc(m.Keys, m2.Keys, expressionsEqual) {
		return false
	}

	return slices.EqualFunc(m.Values, m2.Values, expressionsEqual)
}
