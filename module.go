package slang

type Module interface {
	GetPath() string

	GetStructNames() []string
	GetStruct(name string) (Structure, bool)

	GetFunctionNames() []string
	GetFunction(name string) (Function, bool)

	// GetInterfaceNames() []string
	// GetInterface(name string) Interface

	// GetVariableNames() []string
	// GetVariable(name string) Variable
}
