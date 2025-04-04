package expressions

import (
	"errors"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/jamestunnell/slang/jsonutil"
	"github.com/tidwall/gjson"
)

var errTypeNotFound = errors.New("type not found")
var errTypeNotString = errors.New("type not a string")

func UnmarshalJSON(d []byte) (slang.Expression, error) {
	result := gjson.GetBytes(d, "type")
	if !result.Exists() {
		return nil, errTypeNotFound
	}

	if result.Type != gjson.String {
		return nil, errTypeNotString
	}

	var expr slang.Expression

	var err error

	switch result.String() {
	case slang.StrExprACCESSMEMBER:
		expr, err = jsonutil.UnmarshalAs[AccessMember](d)
	case slang.StrExprADD:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.StrExprAND:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	// case slang.StrExprARRAY:
	// 	expr, err = jsonutil.UnmarshalAs[Array](d)
	// case slang.StrExprARRAYAUTO:
	// 	expr, err = jsonutil.UnmarshalAs[ArrayAuto](d)
	case slang.StrExprBOOL:
		expr, err = jsonutil.UnmarshalAs[Const[bool]](d)
	case slang.StrExprCONCAT:
		expr, err = jsonutil.UnmarshalAs[Concat](d)
	case slang.StrExprDIVIDE:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.StrExprEQUAL:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.StrExprFLOAT:
		expr, err = jsonutil.UnmarshalAs[Const[float64]](d)
	case slang.StrExprFUNC:
		expr, err = jsonutil.UnmarshalAs[Func](d)
	case slang.StrExprGREATER:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.StrExprGREATEREQUAL:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.StrExprIDENTIFIER:
		expr, err = jsonutil.UnmarshalAs[Identifier](d)
	case slang.StrExprINT:
		expr, err = jsonutil.UnmarshalAs[Const[int64]](d)
	case slang.StrExprINVOKE:
		expr, err = jsonutil.UnmarshalAs[Invoke](d)
	case slang.StrExprLESS:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.StrExprLESSEQUAL:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	// case slang.StrExprMAP:
	// 	expr, err = jsonutil.UnmarshalAs[Map](d)
	// case slang.StrExprMAPAUTO:
	// 	expr, err = jsonutil.UnmarshalAs[MapAuto](d)
	case slang.StrExprMULTIPLY:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.StrExprNEGATIVE:
		expr, err = jsonutil.UnmarshalAs[Negative](d)
	case slang.StrExprNOT:
		expr, err = jsonutil.UnmarshalAs[Not](d)
	case slang.StrExprNOTEQUAL:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.StrExprOR:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.StrExprSUBTRACT:
		expr, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.StrExprSTR:
		expr, err = jsonutil.UnmarshalAs[Const[string]](d)
	case slang.StrExprSTRUCT:
		expr, err = jsonutil.UnmarshalAs[Struct](d)
	}

	if err != nil {
		return nil, err
	}

	if expr == nil {
		return nil, customerrs.NewErrUnknownType(result.String())
	}

	return expr, nil
}
