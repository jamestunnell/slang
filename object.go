package slang

type Object interface {
	Equal(Object) bool
	Inspect() string
	Send(name string, args ...Object) (Object, error)
}
