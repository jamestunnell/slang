package slang

const (
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

	MethodCALL = "call"

	// array methods
	MethodSIZE  = "size"
	MethodINDEX = "index"
	MethodFIRST = "first"
	MethodLAST  = "last"
)
