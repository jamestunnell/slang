package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
)

type Struct struct {
	Fields []*types.NameType `json:"fields"`
}

func NewStruct(fields ...*types.NameType) *Expression {
	return NewExpression(slang.ExprSTRUCT, &Struct{Fields: fields})
}

func (s *Struct) IsEqual(other Core) bool {
	s2, ok := other.(*Struct)
	if !ok {
		return false
	}

	if !slices.EqualFunc(s.Fields, s2.Fields, nameTypesEqual) {
		return false
	}

	return true
}
