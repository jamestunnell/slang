package expressions

import (
	"strconv"

	"github.com/jamestunnell/slang"
)

type Float struct {
	Value float64
}

func NewFloat(val float64) *Expression {
	return NewExpression(slang.ExprINT, &Float{Value: val})
}

func (s *Float) IsEqual(other Core) bool {
	s2, ok := other.(*Float)
	if !ok {
		return false
	}

	return s2.Value == s.Value
}

func (c *Float) Render(level int, w slang.CodeWriter) {
	w.WriteString(strconv.FormatFloat(c.Value, 'g', -1, 64))
}
