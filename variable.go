package slang

type Variable interface {
	GetName() string
	GetType() Type
}
