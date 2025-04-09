package slang

type Variable interface {
	GetName() string
	GetInitialValue() Expression
}
