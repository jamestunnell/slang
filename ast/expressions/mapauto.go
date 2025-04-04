package expressions

// import (
// 	"github.com/jamestunnell/slang"
// 	"golang.org/x/exp/slices"
// )

// type MapAuto struct {
// 	*Base

// 	Keys   []slang.Expression `json:"keys"`
// 	Values []slang.Expression `json:"values"`
// }

// func NewMapAuto(
// 	keys []slang.Expression,
// 	vals []slang.Expression) *Map {
// 	return &Map{
// 		Base:   NewBase(slang.ExprMAPAUTO),
// 		Keys:   keys,
// 		Values: vals,
// 	}
// }

// func (m *MapAuto) IsEqual (other Core) bool {
// 	m2, ok := other.(*MapAuto)
// 	if !ok {
// 		return false
// 	}

// 	if !slices.EqualFunc(m.Keys, m2.Keys, expressionsEqual) {
// 		return false
// 	}

// 	return slices.EqualFunc(m.Values, m2.Values, expressionsEqual)
// }

// func expressionsEqual(a, b slang.Expression) bool {
// 	return a.Equal(b)
// }
