package objects

import (
	"math"
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type Float struct {
	Value float64
}

const ClassFLOAT = "Float"

var fltClass = NewBuiltInClass(ClassFLOAT)

func NewFloat(val float64) *Float {
	return &Float{Value: val}
}

func (obj *Float) Class() Class {
	return fltClass
}

func (obj *Float) Inspect() string {
	return strconv.FormatFloat(obj.Value, 'g', -1, 64)
}

func (obj *Float) Truthy() bool {
	return true
}

func (obj *Float) Send(methodName string, args ...Object) (Object, error) {
	// an added instance method would override a standard one
	if m, found := fltClass.GetInstanceMethod(methodName); found {
		return m.Run(args)
	}

	switch methodName {
	case slang.MethodNEG:
		return NewFloat(-obj.Value), nil
	case slang.MethodABS:
		return NewFloat(math.Abs(obj.Value)), nil
	case slang.MethodADD, slang.MethodSUB, slang.MethodMUL, slang.MethodDIV,
		slang.MethodEQ, slang.MethodNEQ, slang.MethodLT, slang.MethodLEQ,
		slang.MethodGT, slang.MethodGEQ:
		if err := checkArgCount(args, 1); err != nil {
			return nil, err
		}

		return obj.sendOne(methodName, args[0])
	}

	err := customerrs.NewErrMethodUndefined(methodName, ClassFLOAT)

	return nil, err
}

func (obj *Float) sendOne(method string, arg Object) (Object, error) {
	flt, ok := arg.(*Float)
	if !ok {
		return nil, customerrs.NewErrArgType(ClassFLOAT, arg.Class().Name())
	}

	var ret Object

	switch method {
	case slang.MethodADD:
		ret = NewFloat(obj.Value + flt.Value)
	case slang.MethodSUB:
		ret = NewFloat(obj.Value - flt.Value)
	case slang.MethodMUL:
		ret = NewFloat(obj.Value * flt.Value)
	case slang.MethodDIV:
		ret = NewFloat(obj.Value / flt.Value)
	case slang.MethodEQ:
		ret = NewBool(obj.Value == flt.Value)
	case slang.MethodNEQ:
		ret = NewBool(obj.Value != flt.Value)
	case slang.MethodLT:
		ret = NewBool(obj.Value < flt.Value)
	case slang.MethodLEQ:
		ret = NewBool(obj.Value <= flt.Value)
	case slang.MethodGT:
		ret = NewBool(obj.Value > flt.Value)
	case slang.MethodGEQ:
		ret = NewBool(obj.Value >= flt.Value)
	}

	return ret, nil
}
