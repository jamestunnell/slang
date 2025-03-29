package slang

type Function interface {
	GetName() string
	GetComment() string

	GetParamNames() []string
	GetParamType(name string) (Type, bool)

	GetReturnTypes() []Type
}
