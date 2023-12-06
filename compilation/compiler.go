package compilation

type Compiler struct {
	symbols  map[string]Symbol
	parent   *Compiler
	errors   []error
	children map[string]*Compiler
}

func NewCompiler(parent *Compiler) *Compiler {
	return &Compiler{
		symbols:  map[string]Symbol{},
		parent:   parent,
		errors:   []error{},
		children: map[string]*Compiler{},
	}
}
