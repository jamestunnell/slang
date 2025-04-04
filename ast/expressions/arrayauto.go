package expressions

// import (
// 	"github.com/jamestunnell/slang"
// 	"golang.org/x/exp/slices"
// )

// type ArrayAuto struct {
// 	Values []slang.Expression `json:"values"`
// }

// func NewArrayAuto(vals ...slang.Expression) slang.Expression {
// 	return &Array{
// 		Base:   NewBase(slang.ExprARRAYAUTO),
// 		Values: vals,
// 	}
// }

// func (a *ArrayAuto) IsEqual (other Core) bool {
// 	a2, ok := other.(*ArrayAuto)
// 	if !ok {
// 		return false
// 	}

// 	return slices.EqualFunc(a.Values, a2.Values, expressionsEqual)
// }
