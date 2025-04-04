package ast

import "github.com/jamestunnell/slang"

type Structure struct {
	Name    string
	Comment string
	Fields  []slang.Field
}

func (s *Structure) GetName() string {
	return s.Name
}

func (s *Structure) GetComment() string {
	return s.Comment
}

func (s *Structure) GetFields() []slang.Field {
	return s.Fields
}
