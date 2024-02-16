package slang

type Object interface {
	Environment

	GetClass() Class

	IsEqual(Object) bool
	Inspect() string
}

type Objects []Object
