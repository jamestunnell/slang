package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type Struct struct {
	Fields []slang.Field `json:"fields"`
}

func NewStruct(fields ...slang.Field) *Expression {
	return NewExpression(slang.ExprSTRUCT, &Struct{Fields: fields})
}

func (s *Struct) IsEqual(other Core) bool {
	s2, ok := other.(*Struct)
	if !ok {
		return false
	}

	if !slices.EqualFunc(s.Fields, s2.Fields, slang.NameTypesEqual) {
		return false
	}

	return true
}
