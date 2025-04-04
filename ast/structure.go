package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
)

type Structure struct {
	Name    string
	Comment string
	Fields  []*types.NameType
}

func (s *Structure) GetName() string {
	return s.Name
}

func (s *Structure) GetComment() string {
	return s.Comment
}

func (s *Structure) GetFields() []slang.Field {
	fields := make([]slang.Field, len(s.Fields))

	for i, f := range s.Fields {
		fields[i] = f
	}

	return fields
}
