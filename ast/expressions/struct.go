package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/tokens"
)

type Struct struct {
	Fields field.Seq `json:"fields"`
}

func NewStruct(fields ...*field.Field) *Expression {
	return NewExpression(slang.ExprSTRUCT, &Struct{Fields: fields})
}

func (s *Struct) IsEqual(other Core) bool {
	s2, ok := other.(*Struct)
	if !ok {
		return false
	}

	if !slices.EqualFunc(s.Fields, s2.Fields, fieldsEqual) {
		return false
	}

	return true
}

func (s *Struct) Render(level int, w slang.CodeWriter) {
	w.WriteString(tokens.StrSTRUCT)

	s.Fields.Render(level, w)
}
