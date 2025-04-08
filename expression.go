package slang

type ExprType int

type Expression interface {
	GetType() ExprType
	IsEqual(Expression) bool
	// Eval(env *objects.Environment) (objects.Object, error)
}

const (
	ExprACCESSMEMBER ExprType = iota
	ExprADD
	ExprAND
	ExprARRAY
	ExprARRAYAUTO
	ExprBOOL
	ExprCONCAT
	ExprDIVIDE
	ExprEMPTY
	ExprEQUAL
	ExprFLOAT
	ExprFUNC
	ExprINVOKE
	ExprGREATER
	ExprGREATEREQUAL
	ExprIDENTIFIER
	ExprINT
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

	StrExprACCESSMEMBER = "ACCESSMEMBER"
	StrExprADD          = "ADD"
	StrExprAND          = "AND"
	StrExprARRAY        = "ARRAY"
	StrExprARRAYAUTO    = "ARRAYAUTO"
	StrExprBOOL         = "BOOL"
	StrExprCONCAT       = "CONCAT"
	StrExprDIVIDE       = "DIVIDE"
	StrExprEMPTY        = "EMPTY"
	StrExprEQUAL        = "EQUAL"
	StrExprFLOAT        = "FLOAT"
	StrExprFUNC         = "FUNC"
	StrExprGREATER      = "GREATER"
	StrExprGREATEREQUAL = "GREATEREQUAL"
	StrExprIDENTIFIER   = "IDENTIFIER"
	StrExprINT          = "INT"
	StrExprINVOKE       = "INVOKE"
	StrExprLESS         = "LESS"
	StrExprLESSEQUAL    = "LESSEQUAL"
	StrExprMAP          = "MAP"
	StrExprMAPAUTO      = "MAPAUTO"
	StrExprMULTIPLY     = "MULTIPLY"
	StrExprNEGATIVE     = "NEGATIVE"
	StrExprNOT          = "NOT"
	StrExprNOTEQUAL     = "NOTEQUAL"
	StrExprOR           = "OR"
	StrExprSUBTRACT     = "SUBTRACT"
	StrExprSTR          = "STR"
	StrExprSTRUCT       = "STRUCT"
)

// func ExpressionsEqual(a, b []Expression) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	for idx, expr := range a {
// 		if !expr.IsEqual(b[idx]) {
// 			return false
// 		}
// 	}

// 	return true
// }

func ParseExprTypeStr(s string) (ExprType, bool) {
	var et ExprType

	switch s {
	case StrExprACCESSMEMBER:
		et = ExprACCESSMEMBER
	case StrExprADD:
		et = ExprADD
	case StrExprAND:
		et = ExprAND
	case StrExprARRAY:
		et = ExprARRAY
	case StrExprARRAYAUTO:
		et = ExprARRAYAUTO
	case StrExprBOOL:
		et = ExprBOOL
	case StrExprCONCAT:
		et = ExprCONCAT
	case StrExprDIVIDE:
		et = ExprDIVIDE
	case StrExprEMPTY:
		et = ExprEMPTY
	case StrExprEQUAL:
		et = ExprEQUAL
	case StrExprFLOAT:
		et = ExprFLOAT
	case StrExprFUNC:
		et = ExprFUNC
	case StrExprGREATER:
		et = ExprGREATER
	case StrExprGREATEREQUAL:
		et = ExprGREATEREQUAL
	case StrExprIDENTIFIER:
		et = ExprIDENTIFIER
	case StrExprINT:
		et = ExprINT
	case StrExprINVOKE:
		et = ExprINVOKE
	case StrExprLESS:
		et = ExprLESS
	case StrExprLESSEQUAL:
		et = ExprLESSEQUAL
	case StrExprMAP:
		et = ExprMAP
	case StrExprMAPAUTO:
		et = ExprMAPAUTO
	case StrExprMULTIPLY:
		et = ExprMULTIPLY
	case StrExprNEGATIVE:
		et = ExprNEGATIVE
	case StrExprNOT:
		et = ExprNOT
	case StrExprNOTEQUAL:
		et = ExprNOTEQUAL
	case StrExprOR:
		et = ExprOR
	case StrExprSUBTRACT:
		et = ExprSUBTRACT
	case StrExprSTR:
		et = ExprSTR
	case StrExprSTRUCT:
		et = ExprSTRUCT
	default:
		return et, false
	}

	return et, true
}

func (et ExprType) String() string {
	var str string

	switch et {
	case ExprACCESSMEMBER:
		str = StrExprACCESSMEMBER
	case ExprADD:
		str = StrExprADD
	case ExprAND:
		str = StrExprAND
	case ExprARRAY:
		str = StrExprARRAY
	case ExprARRAYAUTO:
		str = StrExprARRAYAUTO
	case ExprBOOL:
		str = StrExprBOOL
	case ExprCONCAT:
		str = StrExprCONCAT
	case ExprDIVIDE:
		str = StrExprDIVIDE
	case ExprEMPTY:
		str = StrExprEMPTY
	case ExprEQUAL:
		str = StrExprEQUAL
	case ExprFLOAT:
		str = StrExprFLOAT
	case ExprFUNC:
		str = StrExprFUNC
	case ExprGREATER:
		str = StrExprGREATER
	case ExprGREATEREQUAL:
		str = StrExprGREATEREQUAL
	case ExprIDENTIFIER:
		str = StrExprIDENTIFIER
	case ExprINT:
		str = StrExprINT
	case ExprINVOKE:
		str = StrExprINVOKE
	case ExprLESS:
		str = StrExprLESS
	case ExprLESSEQUAL:
		str = StrExprLESSEQUAL
	case ExprMAP:
		str = StrExprMAP
	case ExprMAPAUTO:
		str = StrExprMAPAUTO
	case ExprMULTIPLY:
		str = StrExprMULTIPLY
	case ExprNEGATIVE:
		str = StrExprNEGATIVE
	case ExprNOT:
		str = StrExprNOT
	case ExprNOTEQUAL:
		str = StrExprNOTEQUAL
	case ExprOR:
		str = StrExprOR
	case ExprSUBTRACT:
		str = StrExprSUBTRACT
	case ExprSTR:
		str = StrExprSTR
	case ExprSTRUCT:
		str = StrExprSTRUCT
	}

	return str
}
