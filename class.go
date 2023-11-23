package slang

type Class interface {
	GetComment() string

	GetFieldNames() []string
	GetFieldType(name string) (string, bool)

	GetMethodNames() []string
	GetMethod(name string) (Function, bool)
}
