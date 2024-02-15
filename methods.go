package slang

const (
	// callable methods
	MethodCALL = "call"

	// unary methods
	MethodNEG = "neg"
	MethodABS = "abs"
	MethodNOT = "not"

	// binary arithmetic
	MethodADD = "add"
	MethodSUB = "sub"
	MethodMUL = "mul"
	MethodDIV = "div"

	// binary equality
	MethodEQ  = "eq"
	MethodNEQ = "neq"

	// binary relational
	MethodLT  = "lt"
	MethodLEQ = "leq"
	MethodGT  = "gt"
	MethodGEQ = "geq"

	// binary logical
	MethodAND = "and"
	MethodOR  = "or"

	// string methods
	MethodCONCAT = "concat"

	// container methods
	MethodELEM = "elem"
	MethodSIZE = "size"

	// array methods
	MethodINDEX = "index"
	MethodFIRST = "first"
	MethodLAST  = "last"
)
