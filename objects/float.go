package objects

import (
	"math"
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/types"
)

type Float struct {
	*Base

	Value float64
}

const floatClassName = "Float"

var floatClass *BuiltInClass

func init() {
	floatClass = NewBuiltInClass(
		types.NewPrimitiveType(floatClassName),
		map[string]slang.MethodFunc{
			slang.MethodNEG: floatNEG,
			slang.MethodABS: floatABS,
			slang.MethodADD: floatADD,
			slang.MethodSUB: floatSUB,
			slang.MethodMUL: floatMUL,
			slang.MethodDIV: floatDIV,
			slang.MethodEQ:  floatEQ,
			slang.MethodNEQ: floatNEQ,
			slang.MethodLT:  floatLT,
			slang.MethodLEQ: floatLEQ,
			slang.MethodGT:  floatGT,
			slang.MethodGEQ: floatGEQ,
		},
	)
}

func NewFloat(val float64) *Float {
	return &Float{
		Base:  NewBase(floatClass),
		Value: val,
	}
}

func (obj *Float) IsEqual(other slang.Object) bool {
	obj2, ok := other.(*Float)
	if !ok {
		return false
	}

	return obj.Value == obj2.Value
}

func (obj *Float) Inspect() string {
	return strconv.FormatFloat(obj.Value, 'g', -1, 64)
}

func floatNEG(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	if err := CheckArgCount(args, 0); err != nil {
		return nil, err
	}

	return slang.Objects{NewFloat(-obj.(*Float).Value)}, nil
}

func floatABS(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	if err := CheckArgCount(args, 0); err != nil {
		return nil, err
	}

	newValue := math.Abs(obj.(*Float).Value)

	return slang.Objects{NewFloat(newValue)}, nil
}

func floatADD(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Float](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewFloat(obj.(*Float).Value + arg.Value)}, nil
}

func floatSUB(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Float](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewFloat(obj.(*Float).Value - arg.Value)}, nil
}

func floatMUL(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Float](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewFloat(obj.(*Float).Value * arg.Value)}, nil
}

func floatDIV(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Float](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewFloat(obj.(*Float).Value / arg.Value)}, nil
}

func floatEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Float](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Float).Value == arg.Value)}, nil
}

func floatNEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Float](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Float).Value != arg.Value)}, nil
}

func floatLT(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Float](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Float).Value < arg.Value)}, nil
}

func floatLEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Float](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Float).Value <= arg.Value)}, nil
}

func floatGT(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Float](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Float).Value > arg.Value)}, nil
}

func floatGEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Float](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Float).Value >= arg.Value)}, nil
}
