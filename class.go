package slang

type MethodFunc func(obj Object, args Objects) (Objects, error)

type Class interface {
	Environment

	GetType() Type
	GetMethod(name string) (MethodFunc, bool)
}
