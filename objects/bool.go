package objects

import (
	"reflect"
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type Bool struct {
	*slang.EnvBase

	Value bool
}

const ClassBOOL = "Bool"

// var boolClass = NewBuiltInClass(ClassBOOL)

var (
	objFalse = NewBool(false)
	objTrue  = NewBool(true)
)

func NewBool(val bool) slang.Object {
	return &Bool{
		EnvBase: slang.NewEnvBase(nil),
		Value:   val,
	}
}

func TRUE() slang.Object {
	return objTrue
}

func FALSE() slang.Object {
	return objFalse
}

func (obj *Bool) Equal(other slang.Object) bool {
	obj2, ok := other.(*Bool)
	if !ok {
		return false
	}

	return obj.Value == obj2.Value
}

func (obj *Bool) Inspect() string {
	return strconv.FormatBool(obj.Value)
}

func (obj *Bool) Send(methodName string, args ...slang.Object) (slang.Object, error) {
	// // an added instance method would override a standard one
	// if m, found := boolClass.GetInstanceMethod(methodName); found {
	// 	return m.Run(args)
	// }

	switch methodName {
	case slang.MethodNOT:
		return NewBool(!obj.Value), nil
	case slang.MethodEQ, slang.MethodNEQ, slang.MethodAND, slang.MethodOR:
		if err := checkArgCount(args, 1); err != nil {
			return nil, err
		}

		arg, ok := args[0].(*Bool)
		if !ok {
			return nil, customerrs.NewErrArgType(ClassBOOL, reflect.TypeOf(args[0]).String())
		}

		var ret slang.Object
		switch methodName {
		case slang.MethodEQ:
			ret = NewBool(obj.Value == arg.Value)
		case slang.MethodNEQ:
			ret = NewBool(obj.Value != arg.Value)
		case slang.MethodAND:
			ret = NewBool(obj.Value && arg.Value)
		case slang.MethodOR:
			ret = NewBool(obj.Value || arg.Value)
		}

		return ret, nil
	}

	err := customerrs.NewErrMethodUndefined(methodName, ClassBOOL)

	return nil, err
}
