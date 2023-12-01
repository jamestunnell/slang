package objects

import (
	"reflect"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type String struct {
	Value string
}

// const ClassSTRING = "String"

// var strClass = NewBuiltInClass(ClassSTRING)

func NewString(val string) slang.Object {
	return &String{Value: val}
}

func (obj *String) Equal(other slang.Object) bool {
	obj2, ok := other.(*String)
	if !ok {
		return false
	}

	return obj.Value == obj2.Value
}

func (obj *String) Inspect() string {
	return obj.Value
}

// func (obj *String) Class() Class {
// 	return strClass
// }

// func (obj *String) Truthy() bool {
// 	return true
// }

func (obj *String) Send(methodName string, args ...slang.Object) (slang.Object, error) {
	// // an added instance method would override a standard one
	// if m, found := strClass.GetInstanceMethod(methodName); found {
	// 	return m.Run(args)
	// }

	switch methodName {
	case slang.MethodSIZE:
		sz := NewInteger(int64(len(obj.Value)))
		return sz, nil
	case slang.MethodADD,
		slang.MethodEQ, slang.MethodNEQ,
		slang.MethodLT, slang.MethodLEQ,
		slang.MethodGT, slang.MethodGEQ:

		if err := checkArgCount(args, 1); err != nil {
			return nil, err
		}

		return obj.sendOne(methodName, args[0])
	}

	err := customerrs.NewErrMethodUndefined(methodName, "String")

	return nil, err
}

func (obj *String) sendOne(method string, arg slang.Object) (slang.Object, error) {
	flt, ok := arg.(*String)
	if !ok {
		return nil, customerrs.NewErrArgType("String", reflect.TypeOf(arg).String())
	}

	var ret slang.Object

	switch method {
	case slang.MethodADD:
		ret = NewString(obj.Value + flt.Value)
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
