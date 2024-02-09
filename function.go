package slang

type Function interface {
	GetComment() string

	GetParamNames() []string
	GetParamType(name string) (Type, bool)

	GetReturnTypes() []Type
}
