package objects

import (
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/types"
)

type Int struct {
	*Base

	Value int64
}

const intClassName = "Int"

var intClass *BuiltInClass

func init() {
	intClass = NewBuiltInClass(
		types.NewPrimitiveType(intClassName),
		map[string]slang.MethodFunc{
			slang.MethodNEG: intNEG,
			slang.MethodABS: intABS,
			slang.MethodADD: intADD,
			slang.MethodSUB: intSUB,
			slang.MethodMUL: intMUL,
			slang.MethodDIV: intDIV,
			slang.MethodEQ:  intEQ,
			slang.MethodNEQ: intNEQ,
			slang.MethodLT:  intLT,
			slang.MethodLEQ: intLEQ,
			slang.MethodGT:  intGT,
			slang.MethodGEQ: intGEQ,
		},
	)
}

func NewInt(val int64) slang.Object {
	return &Int{
		Base:  NewBase(intClass),
		Value: val,
	}
}

func (obj *Int) IsEqual(other slang.Object) bool {
	obj2, ok := other.(*Int)
	if !ok {
		return false
	}

	return obj.Value == obj2.Value
}

func (obj *Int) Inspect() string {
	return strconv.FormatInt(obj.Value, 10)
}

func intNEG(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	if err := CheckArgCount(args, 0); err != nil {
		return nil, err
	}

	return slang.Objects{NewInt(-obj.(*Int).Value)}, nil
}

func intABS(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	if err := CheckArgCount(args, 0); err != nil {
		return nil, err
	}

	newValue := intAbs(obj.(*Int).Value)

	return slang.Objects{NewInt(newValue)}, nil
}

func intADD(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewInt(obj.(*Int).Value + arg.Value)}, nil
}

func intSUB(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewInt(obj.(*Int).Value - arg.Value)}, nil
}

func intMUL(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewInt(obj.(*Int).Value * arg.Value)}, nil
}

func intDIV(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewInt(obj.(*Int).Value / arg.Value)}, nil
}

func intEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Int).Value == arg.Value)}, nil
}

func intNEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Int).Value != arg.Value)}, nil
}

func intLT(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Int).Value < arg.Value)}, nil
}

func intLEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Int).Value <= arg.Value)}, nil
}

func intGT(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Int).Value > arg.Value)}, nil
}

func intGEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Int).Value >= arg.Value)}, nil
}

// 		if _, isFlt := args[0].(*Float); isFlt {
// 			return NewFloat(float64(obj.Value)).Send(method, args[0])
// 		}

func intAbs(val int64) int64 {
	if val > 0 {
		return val
	}

	return -val
}
