package slang

type Module interface {
	GetPath() string
	GetStructs() []Structure
	GetFunctions() []Function

	// GetInterfaceNames() []string
	// GetInterface(name string) Interface

	// GetVariableNames() []string
	// GetVariable(name string) Variable
}
