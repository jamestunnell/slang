package expressions

import (
	"fmt"

	"github.com/jamestunnell/slang"
)

type Str struct {
	Value string
}

func NewStr(val string) *Expression {
	return NewExpression(slang.ExprSTR, &Str{Value: val})
}

func (s *Str) IsEqual(other Core) bool {
	s2, ok := other.(*Str)
	if !ok {
		return false
	}

	return s2.Value == s.Value
}

func (c *Str) Render(level int, w slang.CodeWriter) {
	w.WriteString(fmt.Sprintf(`"%v"`, c.Value))
}
