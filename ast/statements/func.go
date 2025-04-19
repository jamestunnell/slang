package statements

import (
	"slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/sliceutil"
	"github.com/jamestunnell/slang/tokens"
)

type Func struct {
	Name       string       `json:"name"`
	Inputs     field.Seq    `json:"inputs"`
	Outputs    field.Seq    `json:"outputs"`
	Statements []*Statement `json:"statements"`
}

func NewFunc(
	name string,
	inputs, outputs field.Seq,
	statements ...*Statement,
) *Statement {
	core := &Func{
		Name:       name,
		Inputs:     inputs,
		Outputs:    outputs,
		Statements: statements,
	}

	return NewStatement(slang.StatementFUNC, core)
}

func (f *Func) GetName() (string, bool) {
	return f.Name, false
}

func (f *Func) IsEqual(other Core) bool {
	f2, ok := other.(*Func)
	if !ok {
		return false
	}

	if f.Name != f2.Name {
		return false
	}

	if !slices.EqualFunc(f.Inputs, f2.Inputs, fieldsEqual) {
		return false
	}

	if !slices.EqualFunc(f.Outputs, f2.Outputs, fieldsEqual) {
		return false
	}

	if !slices.EqualFunc(f.Statements, f2.Statements, statementsEqual) {
		return false
	}

	return true
}

func (f *Func) Render(level int, w slang.CodeWriter) {
	w.WriteString(tokens.StrFUNC)
	w.WriteString(" ")
	w.WriteString(f.Name)

	stmts := sliceutil.Map(f.Statements, func(s *Statement) slang.Statement { return s })

	expressions.RenderFunction(level, w, f.Inputs, f.Outputs, stmts)
}

func (f *Func) GetInputs() []slang.Param {
	params := make([]slang.Field, len(f.Inputs))

	for i, param := range f.Inputs {
		params[i] = param
	}

	return params
}

func (f *Func) GetOutputs() []slang.Param {
	params := make([]slang.Field, len(f.Outputs))

	for i, param := range f.Outputs {
		params[i] = param
	}

	return params
}

func fieldsEqual(a, b *field.Field) bool {
	return a.IsEqual(b)
}
