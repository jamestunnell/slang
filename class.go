package slang

type Class interface {
	Name() string
	// GetInstanceMethod(string) (Method, bool)
	AddInstanceMethod(string, Method)
}

type Method interface {
	ParamNames() []string
	Run(args []Object) (Object, error)
}
