package slang

type Interface interface {
	GetName() string
	GetFunctionSpecs() []FunctionSpec
}

type FunctionSpec struct {
	Name            string
	Inputs, Outputs []Param
}
