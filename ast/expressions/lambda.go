package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/tokens"
	"golang.org/x/exp/slices"
)

type Lambda struct {
	Inputs     field.Seq         `json:"inputs"`
	Outputs    field.Seq         `json:"outputs"`
	Statements []slang.Statement `json:"statements"`
}

func NewLambda(
	inputs, outputs field.Seq,
	statements ...slang.Statement,
) *Expression {
	return NewExpression(slang.ExprLAMBDA, &Lambda{
		Inputs:     inputs,
		Outputs:    outputs,
		Statements: statements,
	})
}

func (f *Lambda) IsEqual(other Core) bool {
	f2, ok := other.(*Lambda)
	if !ok {
		return false
	}

	if !slices.EqualFunc(f.Inputs, f2.Inputs, fieldsEqual) {
		return false
	}

	if !slices.EqualFunc(f.Outputs, f2.Outputs, fieldsEqual) {
		return false
	}

	if !slices.EqualFunc(f.Statements, f2.Statements, slang.StatementsEqual) {
		return false
	}

	return true
}

func (f *Lambda) Render(level int, w slang.CodeWriter) {
	w.WriteString(tokens.StrLAMBDAOP)

	f.Inputs.Render(level, w)

	if len(f.Outputs) > 0 {
		f.Inputs.Render(level, w)
	}
}

func fieldsEqual(a, b *field.Field) bool {
	return a.IsEqual(b)
}

func RenderFunction(
	level int,
	w slang.CodeWriter,
	inputs, outputs field.Seq,
	stmts []slang.Statement,
) {
	inputs.Render(level, w)

	if len(outputs) > 0 {
		w.WriteString(" ")
		outputs.Render(level, w)
	}

	w.WriteString(" {")
	w.WriteNewline()

	subLevel := level + 1

	for _, stmt := range stmts {
		stmt.Render(subLevel, w)
	}

	w.WriteIndent(level)
	w.WriteString("}")
}
