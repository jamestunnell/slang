package slang

type Module interface {
	GetClassNames() []string
	GetClass(name string) (Class, bool)

	GetFunctionNames() []string
	GetFunction(name string) (Function, bool)

	// GetInterfaceNames() []string
	// GetInterface(name string) Interface

	// GetVariableNames() []string
	// GetVariable(name string) Variable
}
