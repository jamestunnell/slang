package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/ast/types"
)

type Structure struct {
	Name    string
	Comment string
	Fields  []*types.NameType
}

func NewStructure(s *statements.Statement) slang.Structure {
	core, ok := s.Core.(*statements.Struct)
	if !ok {
		return &Structure{
			Name:    "",
			Comment: "",
			Fields:  []*types.NameType{},
		}
	}

	return &Structure{
		Name:    core.Name,
		Comment: s.Comment,
		Fields:  core.Fields,
	}
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
