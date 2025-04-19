package statements

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/tokens"
)

type Interface struct {
	Name  string      `json:"name"`
	Specs []*FuncSpec `json:"functionSpecs"`
}

type FuncSpec struct {
	Name    string    `json:"Name"`
	Inputs  field.Seq `json:"Inputs"`
	Outputs field.Seq `json:"Outputs"`
}

func NewInterface(name string, specs ...*FuncSpec) *Statement {
	core := &Interface{Name: name, Specs: specs}

	return NewStatement(slang.StatementINTERFACE, core)
}

func (i *Interface) GetName() (string, bool) {
	return i.Name, true
}

func (i *Interface) IsEqual(other Core) bool {
	i2, ok := other.(*Interface)
	if !ok {
		return false
	}

	if i.Name != i2.Name {
		return false
	}

	return slices.EqualFunc(i.Specs, i2.Specs, funcSpecsEqual)
}

func (i *Interface) Render(level int, w slang.CodeWriter) {
	w.WriteString(tokens.StrINTERFACE)
	w.WriteString(" ")
	w.WriteString(i.Name)
	w.WriteString(" {")
	w.WriteNewline()

	subLevel := level + 1

	for _, spec := range i.Specs {
		w.WriteIndent(subLevel)
		w.WriteString(spec.Name)

		spec.Inputs.Render(level, w)

		if len(spec.Outputs) > 0 {
			w.WriteString(" ")

			spec.Outputs.Render(level, w)
		}

		w.WriteNewline()
	}

	w.WriteIndent(level)
	w.WriteString("}")
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
