package slang

type Method struct {
	Function MethodFn
	Params   []string
}

type MethodFn func(env *Environment) (Object, error)

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

	// binary logical comparison
	MethodEQ  = "eq"
	MethodNEQ = "neq"
	MethodLT  = "lt"
	MethodLEQ = "leq"
	MethodGT  = "gt"
	MethodGEQ = "geq"

	// function methods
	MethodCALL   = "call"
	MethodPARAMS = "params"

	// array methods
	MethodFIRST = "first"
	MethodLAST  = "last"
	MethodSIZE  = "size"
	MethodAT    = "at"
)

func NewMethod(fn MethodFn, params ...string) *Method {
	return &Method{
		Function: fn,
		Params:   params,
	}
}
