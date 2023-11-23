package objects

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type String struct {
	Value string
}

const ClassSTRING = "String"

var strClass = NewBuiltInClass(ClassSTRING)

func NewString(val string) Object {
	return &String{Value: val}
}

func (obj *String) Class() Class {
	return strClass
}

func (obj *String) Inspect() string {
	return obj.Value
}

func (obj *String) Truthy() bool {
	return true
}

func (obj *String) Type() ObjectType {
	return ObjectSTRING
}

func (obj *String) Send(methodName string, args ...Object) (Object, error) {
	// an added instance method would override a standard one
	if m, found := strClass.GetInstanceMethod(methodName); found {
		return m.Run(args)
	}

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

	err := customerrs.NewErrMethodUndefined(methodName, ClassSTRING)

	return nil, err
}

func (obj *String) sendOne(method string, arg Object) (Object, error) {
	flt, ok := arg.(*String)
	if !ok {
		return nil, customerrs.NewErrArgType(ClassSTRING, arg.Class().Name())
	}

	var ret Object

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
