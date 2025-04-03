package slang

type Function interface {
	GetName() string
	GetComment() string

	GetInputParams() []Param
	GetOutputParams() []Param
}
