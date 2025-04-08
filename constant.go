package slang

type Constant interface {
	GetName() string
	GetValue() Expression
}
