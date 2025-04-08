package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/field"
)

type Struct struct {
	Fields []*field.Field `json:"fields"`
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
