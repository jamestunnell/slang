package slang

type Function interface {
	GetName() string
	GetComment() string

	GetInputs() []Param
	GetOutputs() []Param
}
