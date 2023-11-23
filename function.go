package slang

type Function interface {
	GetComment() string

	GetParamNames() []string
	GetParamType(name string) (string, bool)

	GetReturnTypes() []string
}
