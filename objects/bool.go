package objects

import (
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type Bool struct {
	Value bool
}

const ClassBOOL = "Bool"

var boolClass = NewBuiltInClass(ClassBOOL)

func NewBool(val bool) slang.Object {
	return &Bool{Value: val}
}

func (obj *Bool) Class() slang.Class {
	return boolClass
}

func (obj *Bool) Inspect() string {
	return strconv.FormatBool(obj.Value)
}

func (obj *Bool) Truthy() bool {
	return obj.Value
}

func (obj *Bool) Send(methodName string, args ...slang.Object) (slang.Object, error) {
	// an added instance method would override a standard one
	if m, found := boolClass.GetInstanceMethod(methodName); found {
		return m.Run(args)
	}

	switch methodName {
	case slang.MethodNOT:
		return NewBool(!obj.Value), nil
	case slang.MethodEQ, slang.MethodNEQ:
		if err := checkArgCount(args, 1); err != nil {
			return nil, err
		}

		arg, ok := args[0].(*Bool)
		if !ok {
			return nil, customerrs.NewErrArgType(ClassBOOL, args[0].Class().Name())
		}

		var ret slang.Object
		switch methodName {
		case slang.MethodEQ:
			ret = NewBool(obj.Value == arg.Value)
		case slang.MethodNEQ:
			ret = NewBool(obj.Value != arg.Value)
		}

		return ret, nil
	}

	err := customerrs.NewErrMethodUndefined(methodName, ClassBOOL)

	return nil, err
}
