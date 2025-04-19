package expressions

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/jsonutil"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Expression struct {
	Type slang.ExprType `json:"-"`
	Core Core           `json:"-"`
}

type Core interface {
	IsEqual(Core) bool
	Render(level int, w slang.CodeWriter)
}

func NewExpression(typ slang.ExprType, core Core) *Expression {
	return &Expression{
		Type: typ,
		Core: core,
	}
}

func (s *Expression) GetType() slang.ExprType {
	return s.Type
}

func (s *Expression) IsEqual(other slang.Expression) bool {
	s2, ok := other.(*Expression)
	if !ok {
		return false
	}

	if s.Type != s2.Type {
		return false
	}

	return s.Core.IsEqual(s2.Core)
}

func (s *Expression) MarshalJSON() ([]byte, error) {
	d, err := json.Marshal(s.Core)
	if err != nil {
		return []byte{}, err
	}

	d, err = sjson.SetBytes(d, "type", s.Type.String())
	if err != nil {
		return []byte{}, err
	}

	return d, nil
}

func (s *Expression) Render(level int, w slang.CodeWriter) {
	s.Core.Render(level, w)
}

var (
	errTypeNotFound  = errors.New("type not found")
	errTypeNotString = errors.New("type not a string")
)

func (s *Expression) UnmarshalJSON(d []byte) error {
	result := gjson.GetBytes(d, "type")
	if !result.Exists() {
		return errTypeNotFound
	}

	if result.Type != gjson.String {
		return errTypeNotString
	}

	exprType, ok := slang.ParseExprTypeStr(result.String())
	if !ok {
		return fmt.Errorf("invalid expression type string '%s'", result.String())
	}

	var err error

	var core Core

	switch exprType {
	case slang.ExprACCESSMEMBER:
		core, err = jsonutil.UnmarshalAs[AccessMember](d)
	case slang.ExprADD:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.ExprAND:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	// case slang.ExprARRAY:
	// 	core, err = jsonutil.UnmarshalAs[Array](d)
	// case slang.ExprARRAYAUTO:
	// 	core, err = jsonutil.UnmarshalAs[ArrayAuto](d)
	case slang.ExprBOOL:
		core, err = jsonutil.UnmarshalAs[Const[bool]](d)
	case slang.ExprCONCAT:
		core, err = jsonutil.UnmarshalAs[Concat](d)
	case slang.ExprDIVIDE:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.ExprEQUAL:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.ExprFLOAT:
		core, err = jsonutil.UnmarshalAs[Const[float64]](d)
	case slang.ExprLAMBDA:
		core, err = jsonutil.UnmarshalAs[Lambda](d)
	case slang.ExprGREATER:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.ExprGREATEREQUAL:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.ExprIDENTIFIER:
		core, err = jsonutil.UnmarshalAs[Identifier](d)
	case slang.ExprINT:
		core, err = jsonutil.UnmarshalAs[Const[int64]](d)
	case slang.ExprINVOKE:
		core, err = jsonutil.UnmarshalAs[Invoke](d)
	case slang.ExprLESS:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.ExprLESSEQUAL:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	// case slang.ExprMAP:
	// 	core, err = jsonutil.UnmarshalAs[Map](d)
	// case slang.ExprMAPAUTO:
	// 	core, err = jsonutil.UnmarshalAs[MapAuto](d)
	case slang.ExprMULTIPLY:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.ExprNEGATIVE:
		core, err = jsonutil.UnmarshalAs[Negative](d)
	case slang.ExprNOT:
		core, err = jsonutil.UnmarshalAs[Not](d)
	case slang.ExprNOTEQUAL:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.ExprOR:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.ExprSUBTRACT:
		core, err = jsonutil.UnmarshalAs[BinaryOperation](d)
	case slang.ExprSTR:
		core, err = jsonutil.UnmarshalAs[Const[string]](d)
	case slang.ExprSTRUCT:
		core, err = jsonutil.UnmarshalAs[Struct](d)
	}

	if err != nil {
		return fmt.Errorf("failed to unmarshal core: %w", err)
	}

	if core == nil {
		return fmt.Errorf("unhandled expression type %s", result.String())
	}

	s.Core = core
	s.Type = exprType

	return nil
}
