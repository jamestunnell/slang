package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Map struct {
	*Base

	KeyType   string             `json:"keyType"`
	Keys      []slang.Expression `json:"keys"`
	ValueType string             `json:"valueType"`
	Values    []slang.Expression `json:"values"`
}

func NewMap(
	keyType string, keys []slang.Expression,
	valType string, vals []slang.Expression) *Map {
	return &Map{
		Base:      NewBase(slang.ExprMAP),
		KeyType:   keyType,
		Keys:      keys,
		ValueType: valType,
		Values:    vals,
	}
}

func (m *Map) Equal(other slang.Expression) bool {
	m2, ok := other.(*Map)
	if !ok {
		return false
	}

	if m.KeyType != m2.KeyType {
		return false
	}

	if m.ValueType != m2.ValueType {
		return false
	}

	if !slices.EqualFunc(m.Keys, m2.Keys, expressionsEqual) {
		return false
	}

	return slices.EqualFunc(m.Values, m2.Values, expressionsEqual)
}
