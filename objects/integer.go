package objects

import (
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type Integer struct {
	Value int64
}

const (
	ClassINTEGER = "Integer"
	ParamX       = "x"
)

var intClass = NewBuiltInClass(ClassINTEGER)

func NewInteger(val int64) slang.Object {
	return &Integer{Value: val}
}

func (obj *Integer) Class() slang.Class {
	return intClass
}

func (obj *Integer) Inspect() string {
	return strconv.FormatInt(obj.Value, 10)
}

func (obj *Integer) Truthy() bool {
	return true
}

func (obj *Integer) Send(methodName string, args ...slang.Object) (slang.Object, error) {
	// an added instance method would override a standard one
	if m, found := intClass.GetInstanceMethod(methodName); found {
		return m.Run(args)
	}

	switch methodName {
	case slang.MethodNEG:
		return NewInteger(-obj.Value), nil
	case slang.MethodABS:
		return NewInteger(intAbs(obj.Value)), nil
	case slang.MethodADD, slang.MethodSUB, slang.MethodMUL, slang.MethodDIV,
		slang.MethodEQ, slang.MethodNEQ, slang.MethodLT, slang.MethodLEQ,
		slang.MethodGT, slang.MethodGEQ:
		if err := checkArgCount(args, 1); err != nil {
			return nil, err
		}

		if args[0].Class().Name() == ClassFLOAT {
			return NewFloat(float64(obj.Value)).Send(methodName, args...)
		}

		return obj.sendOne(methodName, args[0])
	}

	err := customerrs.NewErrMethodUndefined(methodName, ClassINTEGER)

	return nil, err
}

func (obj *Integer) sendOne(method string, arg slang.Object) (slang.Object, error) {
	otherInt, ok := arg.(*Integer)
	if !ok {
		return nil, customerrs.NewErrArgType(ClassFLOAT, arg.Class().Name())
	}

	var ret slang.Object

	switch method {
	case slang.MethodADD:
		ret = NewInteger(obj.Value + otherInt.Value)
	case slang.MethodSUB:
		ret = NewInteger(obj.Value - otherInt.Value)
	case slang.MethodMUL:
		ret = NewInteger(obj.Value * otherInt.Value)
	case slang.MethodDIV:
		ret = NewInteger(obj.Value / otherInt.Value)
	case slang.MethodEQ:
		ret = NewBool(obj.Value == otherInt.Value)
	case slang.MethodNEQ:
		ret = NewBool(obj.Value != otherInt.Value)
	case slang.MethodLT:
		ret = NewBool(obj.Value < otherInt.Value)
	case slang.MethodLEQ:
		ret = NewBool(obj.Value <= otherInt.Value)
	case slang.MethodGT:
		ret = NewBool(obj.Value > otherInt.Value)
	case slang.MethodGEQ:
		ret = NewBool(obj.Value >= otherInt.Value)
	}

	return ret, nil
}

func intAbs(val int64) int64 {
	if val > 0 {
		return val
	}

	return -val
}

func checkArgCount(args []slang.Object, count int) error {
	if len(args) != count {
		return customerrs.NewErrArgCount(count, len(args))
	}

	return nil
}
