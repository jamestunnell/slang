package ast

import "github.com/jamestunnell/slang"

type Function struct {
	Name    string
	Comment string
	Inputs  []slang.Param
	Outputs []slang.Param
}

func (f *Function) GetName() string {
	return f.Name
}

func (f *Function) GetComment() string {
	return f.Comment
}

func (f *Function) GetInputs() []slang.Param {
	return f.Inputs
}

func (f *Function) GetOutputs() []slang.Param {
	return f.Outputs
}
