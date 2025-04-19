package expressions

import (
	"strconv"

	"github.com/jamestunnell/slang"
)

type Int struct {
	Value int64
}

func NewInt(val int64) *Expression {
	return NewExpression(slang.ExprINT, &Int{Value: val})
}

func (s *Int) IsEqual(other Core) bool {
	s2, ok := other.(*Int)
	if !ok {
		return false
	}

	return s2.Value == s.Value
}

func (c *Int) Render(level int, w slang.CodeWriter) {
	w.WriteString(strconv.FormatInt(c.Value, 10))
}
