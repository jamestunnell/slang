package slang

type Module interface {
	GetPathParts() []string
	GetStructures() []Structure
	GetFunctions() []Function
	GetVariables() []Variable
	GetConstants() []Constant

	IsEqual(Module) bool
}
