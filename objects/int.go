package objects

import (
	"reflect"
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type Int struct {
	Value int64
}

const ClassINT = "Int"

// var intClass = NewBuiltInClass(ClassINT)

func NewInt(val int64) slang.Object {
	return &Int{Value: val}
}

func (obj *Int) Equal(other slang.Object) bool {
	obj2, ok := other.(*Int)
	if !ok {
		return false
	}

	return obj.Value == obj2.Value
}

func (obj *Int) Inspect() string {
	return strconv.FormatInt(obj.Value, 10)
}

// func (obj *Int) Class() Class {
// 	return intClass
// }

// func (obj *Int) Truthy() bool {
// 	return true
// }

func (obj *Int) Send(method string, args ...slang.Object) (slang.Object, error) {
	// // an added instance method would override a standard one
	// if m, found := intClass.GetInstanceMethod(method); found {
	// 	return m.Run(args)
	// }

	switch method {
	case slang.MethodNEG, slang.MethodABS:
		if err := checkArgCount(args, 0); err != nil {
			return nil, err
		}

		return obj.sendZero(method)
	case slang.MethodADD, slang.MethodSUB, slang.MethodMUL, slang.MethodDIV,
		slang.MethodEQ, slang.MethodNEQ, slang.MethodLT, slang.MethodLEQ,
		slang.MethodGT, slang.MethodGEQ:
		if err := checkArgCount(args, 1); err != nil {
			return nil, err
		}

		if _, isFlt := args[0].(*Float); isFlt {
			return NewFloat(float64(obj.Value)).Send(method, args[0])
		}

		return obj.sendOne(method, args[0])
	}

	err := customerrs.NewErrMethodUndefined(method, ClassINT)

	return nil, err
}

func (obj *Int) sendZero(method string) (slang.Object, error) {
	switch method {
	case slang.MethodNEG:
		return NewInt(-obj.Value), nil
	case slang.MethodABS:
		return NewInt(intAbs(obj.Value)), nil
	}
}

func (obj *Int) sendOne(method string, arg slang.Object) (slang.Object, error) {
	otherInt, ok := arg.(*Int)
	if !ok {
		return nil, customerrs.NewErrArgType(ClassINT, reflect.TypeOf(arg).String())
	}

	var ret slang.Object

	switch method {
	case slang.MethodADD:
		ret = NewInt(obj.Value + otherInt.Value)
	case slang.MethodSUB:
		ret = NewInt(obj.Value - otherInt.Value)
	case slang.MethodMUL:
		ret = NewInt(obj.Value * otherInt.Value)
	case slang.MethodDIV:
		ret = NewInt(obj.Value / otherInt.Value)
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
