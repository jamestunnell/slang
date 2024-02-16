package slang

type ClassInfo interface {
	GetComment() string

	GetFieldNames() []string
	GetFieldType(name string) (string, bool)

	GetMethodNames() []string
	GetMethodInfo(name string) (FunctionInfo, bool)
}
