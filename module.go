package slang

type Module interface {
	GetStructNames() []string
	GetStruct(name string) (Struct, bool)

	GetFunctionNames() []string
	GetFunction(name string) (Function, bool)

	// GetInterfaceNames() []string
	// GetInterface(name string) Interface

	// GetVariableNames() []string
	// GetVariable(name string) Variable
}
