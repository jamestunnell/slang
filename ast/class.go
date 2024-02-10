package ast

import (
	"golang.org/x/exp/maps"

	"github.com/jamestunnell/slang"
)

type Class struct {
	Comment string               `json:"comment"`
	Fields  []*Param             `json:"fields"`
	Methods map[string]*Function `json:"methods"`
}

func NewClass() *Class {
	return &Class{
		Comment: "",
		Fields:  []*Param{},
		Methods: map[string]*Function{},
	}
}

func (c *Class) Equal(other *Class) bool {
	if c.Comment != other.Comment {
		return false
	}

	if len(c.Fields) != len(c.Fields) {
		return false
	}

	for i, param := range c.Fields {
		if !param.Equal(other.Fields[i]) {
			return false
		}
	}

	if len(c.Methods) != len(other.Methods) {
		return false
	}

	for i, method := range c.Methods {
		if !method.Equal(other.Methods[i]) {
			return false
		}
	}

	return true
}

func (c *Class) GetComment() string {
	return c.Comment
}

func (c *Class) GetFieldNames() []string {
	names := make([]string, len(c.Fields))

	for i, param := range c.Fields {
		names[i] = param.Name
	}

	return names
}

func (c *Class) GetFieldType(name string) (slang.Type, bool) {
	for _, field := range c.Fields {
		if name == field.Name {
			return field.Type, true
		}
	}

	return nil, false
}

func (c *Class) GetMethodNames() []string {
	return maps.Keys(c.Methods)
}

func (c *Class) GetMethod(name string) (slang.Function, bool) {
	m, found := c.Methods[name]
	if !found {
		return nil, false
	}

	return m, true
}
