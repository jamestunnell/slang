package statements

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Struct struct {
	*Base

	Name       string            `json:"name"`
	Statements []slang.Statement `json:"statements"`
}

func NewStruct(name string, stmts ...slang.Statement) *Struct {
	return &Struct{
		Base:       NewBase(slang.StatementSTRUCT),
		Name:       name,
		Statements: stmts,
	}
}

func (c *Struct) Equal(other slang.Statement) bool {
	c2, ok := other.(*Struct)
	if !ok {
		return false
	}

	return c.Name == c2.Name && slang.StatementsEqual(c.Statements, c2.Statements)
}

func (s *Struct) GetName() string {
	return s.Name
}

func (s *Struct) GetFieldNames() []string {
	names := []string{}

	for _, s := range s.Statements {
		if f, ok := s.(*StructField); ok {
			names = append(names, f.Names...)
		}
	}

	return names
}

func (s *Struct) GetFieldType(name string) (string, bool) {
	for _, s := range s.Statements {
		if f, ok := s.(*StructField); ok {
			if slices.Contains(f.Names, name) {
				return f.ValueType.String(), true
			}
		}
	}

	return "", false
}
