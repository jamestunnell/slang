package slang

type Module interface {
	GetPathParts() []string
	GetStructures() []Structure
	GetFunctions() []Function
	IsEqual(Module) bool
	// GetInterfaceNames() []string
	// GetInterface(name string) Interface

	// GetVariableNames() []string
	// GetVariable(name string) Variable
}
