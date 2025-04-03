package slang

import "encoding/json"

type ExprType int

type Expression interface {
	GetType() ExprType
	Equal(Expression) bool
	// Eval(env *objects.Environment) (objects.Object, error)
}

const (
	ExprACCESSELEM ExprType = iota
	ExprACCESSMEMBER
	ExprADD
	ExprAND
	ExprARRAY
	ExprARRAYAUTO
	ExprBOOL
	ExprCONCAT
	ExprDIVIDE
	ExprEQUAL
	ExprFLOAT
	ExprFUNC
	ExprINVOKE
	ExprGREATER
	ExprGREATEREQUAL
	ExprIDENTIFIER
	ExprINT
	ExprKEY
	ExprLESS
	ExprLESSEQUAL
	ExprMAP
	ExprMAPAUTO
	ExprMULTIPLY
	ExprNEGATIVE
	ExprNOT
	ExprNOTEQUAL
	ExprOR
	ExprSUBTRACT
	ExprSTR
	ExprSTRUCT
	ExprSTRUCTVAL
)

func ExpressionsEqual(a, b []Expression) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, expr := range a {
		if !expr.Equal(b[idx]) {
			return false
		}
	}

	return true
}

func (et ExprType) MarshalJSON() ([]byte, error) {
	return json.Marshal(et.String())
}

func (et ExprType) String() string {
	var str string

	switch et {
	case ExprACCESSELEM:
		str = "ACCESSELEM"
	case ExprACCESSMEMBER:
		str = "ACCESSMEMBER"
	case ExprADD:
		str = "ADD"
	case ExprAND:
		str = "AND"
	case ExprARRAY:
		str = "ARRAY"
	case ExprBOOL:
		str = "BOOL"
	case ExprCONCAT:
		str = "CONCAT"
	case ExprDIVIDE:
		str = "DIVIDE"
	case ExprEQUAL:
		str = "EQUAL"
	case ExprFLOAT:
		str = "FLOAT"
	case ExprFUNC:
		str = "FUNC"
	case ExprGREATER:
		str = "GREATER"
	case ExprGREATEREQUAL:
		str = "GREATEREQUAL"
	case ExprIDENTIFIER:
		str = "IDENTIFIER"
	case ExprINT:
		str = "INT"
	case ExprINVOKE:
		str = "INVOKE"
	case ExprKEY:
		str = "KEY"
	case ExprLESS:
		str = "LESS"
	case ExprLESSEQUAL:
		str = "LESSEQUAL"
	case ExprMAP:
		str = "MAP"
	case ExprMULTIPLY:
		str = "MULTIPLY"
	case ExprNEGATIVE:
		str = "NEGATIVE"
	case ExprNOT:
		str = "NOT"
	case ExprNOTEQUAL:
		str = "NOTEQUAL"
	case ExprOR:
		str = "OR"
	case ExprSUBTRACT:
		str = "SUBTRACT"
	case ExprSTR:
		str = "STR"
	case ExprSTRUCT:
		str = "STRUCT"
	}

	return str
}
