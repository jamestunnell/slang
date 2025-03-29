package ast

import (
	"github.com/jamestunnell/slang"
)

type Struct struct {
	Name    string            `json:"name"`
	Comment string            `json:"comment"`
	Fields  []*slang.NameType `json:"fields"`
}

func NewStruct() *Struct {
	return &Struct{
		Comment: "",
		Fields:  []*slang.NameType{},
	}
}

func (s *Struct) Equal(other *Struct) bool {
	if s.Comment != other.Comment {
		return false
	}

	if len(s.Fields) != len(other.Fields) {
		return false
	}

	for i, field := range s.Fields {
		if !field.Equal(other.Fields[i]) {
			return false
		}
	}

	return true
}

func (s *Struct) GetName() string {
	return s.Name
}

func (s *Struct) GetComment() string {
	return s.Comment
}

func (s *Struct) GetFieldNames() []string {
	names := make([]string, len(s.Fields))

	for i, field := range s.Fields {
		names[i] = field.Name
	}

	return names
}

func (s *Struct) GetFieldType(name string) (slang.Type, bool) {
	for _, field := range s.Fields {
		if name == field.Name {
			return field.Type, true
		}
	}

	return nil, false
}
