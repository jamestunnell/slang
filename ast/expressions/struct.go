package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type Struct struct {
	*Base

	Fields []slang.Field `json:"fields"`
}

func NewStruct(fields ...slang.Field) slang.Expression {
	return &Struct{
		Base:   NewBase(slang.ExprSTRUCT),
		Fields: fields,
	}
}

func (s *Struct) Equal(other slang.Expression) bool {
	s2, ok := other.(*Struct)
	if !ok {
		return false
	}

	if !slices.EqualFunc(s.Fields, s2.Fields, slang.NameTypesEqual) {
		return false
	}

	return true
}
