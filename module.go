package slang

type Module interface {
	GetPathParts() []string

	GetConstants() []Constant
	GetFunctions() []Function
	GetImports() []Import
	GetInterfaces() []Interface
	GetStructures() []Structure
	GetVariables() []Variable

	IsEqual(Module) bool
}
