package slang

type ExprType int

type Expression interface {
	Type() ExprType
	Equal(Expression) bool
	Eval(env *Environment) (Object, error)
}

const (
	ExprADD ExprType = iota
	ExprARRAY
	ExprBOOL
	ExprDIVIDE
	ExprEQUAL
	ExprFLOAT
	ExprFUNCTIONCALL
	ExprFUNCTIONLITERAL
	ExprGREATER
	ExprGREATEREQUAL
	ExprIDENTIFIER
	ExprIF
	ExprINDEX
	ExprINTEGER
	ExprLESS
	ExprLESSEQUAL
	ExprMETHODCALL
	ExprMULTIPLY
	ExprNEGATIVE
	ExprNOT
	ExprNOTEQUAL
	ExprSUBTRACT
	ExprSTRING
)

func (st ExprType) String() string {
	var str string

	switch st {
	case ExprADD:
		str = "ADD"
	case ExprARRAY:
		str = "ARRAY"
	case ExprBOOL:
		str = "BOOL"
	case ExprDIVIDE:
		str = "DIVIDE"
	case ExprEQUAL:
		str = "EQUAL"
	case ExprFLOAT:
		str = "FLOAT"
	case ExprFUNCTIONCALL:
		str = "FUNCTIONCALL"
	case ExprFUNCTIONLITERAL:
		str = "FUNCTIONLITERAL"
	case ExprGREATER:
		str = "GREATER"
	case ExprGREATEREQUAL:
		str = "GREATEREQUAL"
	case ExprIDENTIFIER:
		str = "IDENTIFIER"
	case ExprIF:
		str = "IF"
	case ExprINDEX:
		str = "INDEX"
	case ExprINTEGER:
		str = "INTEGER"
	case ExprLESS:
		str = "LESS"
	case ExprLESSEQUAL:
		str = "LESSEQUAL"
	case ExprMETHODCALL:
		str = "METHODCALL"
	case ExprMULTIPLY:
		str = "MULTIPLY"
	case ExprNEGATIVE:
		str = "NEGATIVE"
	case ExprNOT:
		str = "NOT"
	case ExprNOTEQUAL:
		str = "NOTEQUAL"
	case ExprSUBTRACT:
		str = "SUBTRACT"
	case ExprSTRING:
		str = "STRING"
	}

	return str
}
