package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Map struct {
	*Base

	KeyType   slang.Type         `json:"keyType"`
	Keys      []slang.Expression `json:"keys"`
	ValueType slang.Type         `json:"valueType"`
	Values    []slang.Expression `json:"values"`
}

func NewMap(
	keyType slang.Type, keys []slang.Expression,
	valType slang.Type, vals []slang.Expression) *Map {
	return &Map{
		Base:      NewBase(slang.ExprMAP),
		KeyType:   keyType,
		Keys:      keys,
		ValueType: valType,
		Values:    vals,
	}
}

func (m *Map) IsEqual(other slang.Expression) bool {
	m2, ok := other.(*Map)
	if !ok {
		return false
	}

	if !m.KeyType.IsEqual(m2.KeyType) {
		return false
	}

	if !m.ValueType.IsEqual(m2.ValueType) {
		return false
	}

	if !slices.EqualFunc(m.Keys, m2.Keys, expressionsEqual) {
		return false
	}

	return slices.EqualFunc(m.Values, m2.Values, expressionsEqual)
}

func (m *Map) Eval(env slang.Environment) (slang.Object, error) {
	// TODO
}
