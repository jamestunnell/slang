package statements

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/tokens"
)

type Struct struct {
	Name   string    `json:"name"`
	Fields field.Seq `json:"fields"`
}

func NewStruct(name string, fields ...*field.Field) *Statement {
	core := &Struct{Name: name, Fields: fields}

	return NewStatement(slang.StatementSTRUCT, core)
}

func (s *Struct) GetName() (string, bool) {
	return s.Name, true
}

func (s *Struct) IsEqual(other Core) bool {
	s2, ok := other.(*Struct)
	if !ok {
		return false
	}

	if s.Name != s2.Name {
		return false
	}

	return slices.EqualFunc(s.Fields, s2.Fields, fieldsEqual)
}

func (s *Struct) Render(level int, w slang.CodeWriter) {
	w.WriteString(tokens.StrSTRUCT)
	w.WriteString(" ")
	w.WriteString(s.Name)

	s.Fields.Render(level, w)
}
