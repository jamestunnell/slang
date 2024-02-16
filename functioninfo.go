package slang

type FunctionInfo interface {
	GetComment() string

	GetParamNames() []string
	GetParamType(name string) (Type, bool)

	GetReturnTypes() []Type
}
