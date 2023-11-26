package slang

import "encoding/json"

type ExprType int

type Expression interface {
	Type() ExprType
	Equal(Expression) bool
	// Eval(env *objects.Environment) (objects.Object, error)
}

const (
	ExprADD ExprType = iota
	ExprARRAY
	ExprBOOL
	ExprCALL
	ExprDIVIDE
	ExprEQUAL
	ExprFLOAT
	ExprFUNC
	ExprFUNCCALL
	ExprGREATER
	ExprGREATEREQUAL
	ExprIDENTIFIER
	ExprIF
	ExprIFELSE
	ExprINDEX
	ExprINTEGER
	ExprLESS
	ExprLESSEQUAL
	ExprMEMBERACCESS
	ExprMULTIPLY
	ExprNEGATIVE
	ExprNOT
	ExprNOTEQUAL
	ExprSUBTRACT
	ExprSTRING
)

func (et ExprType) MarshalJSON() ([]byte, error) {
	return json.Marshal(et.String())
}

func (et ExprType) String() string {
	var str string

	switch et {
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
	case ExprFUNC:
		str = "FUNC"
	case ExprFUNCCALL:
		str = "FUNCCALL"
	case ExprGREATER:
		str = "GREATER"
	case ExprGREATEREQUAL:
		str = "GREATEREQUAL"
	case ExprIDENTIFIER:
		str = "IDENTIFIER"
	case ExprIF:
		str = "IF"
	case ExprIFELSE:
		str = "IFELSE"
	case ExprINDEX:
		str = "INDEX"
	case ExprINTEGER:
		str = "INTEGER"
	case ExprLESS:
		str = "LESS"
	case ExprLESSEQUAL:
		str = "LESSEQUAL"
	case ExprMEMBERACCESS:
		str = "MEMBERACCESS"
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
