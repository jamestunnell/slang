package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/tokens"
)

type Const struct {
	Name  string                  `json:"name"`
	Value *expressions.Expression `json:"value"`
}

func NewConst(name string, val *expressions.Expression) *Statement {
	core := &Const{Name: name, Value: val}

	return NewStatement(slang.StatementCONST, core)
}

func (c *Const) GetName() (string, bool) {
	return c.Name, true
}

func (c *Const) IsEqual(other Core) bool {
	c2, ok := other.(*Const)
	if !ok {
		return false
	}

	if !c.Value.IsEqual(c2.Value) {
		return false
	}

	return c.Name == c2.Name
}

func (c *Const) Render(level int, w slang.CodeWriter) {
	w.WriteString(tokens.StrCONST)
	w.WriteString(" ")
	w.WriteString(c.Name)
	w.WriteString(" ")

	c.Value.Render(level, w)
}
