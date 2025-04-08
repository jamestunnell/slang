package statements

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/field"
)

type Interface struct {
	Name  string      `json:"name"`
	Specs []*FuncSpec `json:"functionSpecs"`
}

type FuncSpec struct {
	Name    string         `json:"Name"`
	Inputs  []*field.Field `json:"Inputs"`
	Outputs []*field.Field `json:"Outputs"`
}

func NewInterface(name string, specs ...*FuncSpec) *Statement {
	core := &Interface{Name: name, Specs: specs}

	return NewStatement(slang.StatementINTERFACE, core)
}

func (c *Interface) IsEqual(other Core) bool {
	c2, ok := other.(*Interface)
	if !ok {
		return false
	}

	if c.Name != c2.Name {
		return false
	}

	return slices.EqualFunc(c.Specs, c2.Specs, funcSpecsEqual)
}

func funcSpecsEqual(a, b *FuncSpec) bool {
	if a.Name != b.Name {
		return false
	}

	if !slices.EqualFunc(a.Inputs, b.Inputs, fieldsEqual) {
		return false
	}

	if !slices.EqualFunc(a.Outputs, b.Outputs, fieldsEqual) {
		return false
	}

	return true
}
