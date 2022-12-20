package objects

import (
	"github.com/jamestunnell/slang"
)

type Array struct {
	Elements []slang.Object
}

func NewArray(vals ...slang.Object) slang.Object {
	return &Array{Elements: vals}
}

func (obj *Array) Inspect() string {
	return "[...]"
}

func (obj *Array) Truthy() bool {
	return true
}

func (obj *Array) Type() slang.ObjectType {
	return slang.ObjectARRAY
}

func (obj *Array) Send(method string, args ...slang.Object) (slang.Object, error) {
	// switch method {
	// case slang.MethodNOT:
	// 	return NewArray(!obj.Value), nil
	// case slang.MethodEQ, slang.MethodNEQ:
	// 	if err := checkArgCount(args, 1); err != nil {
	// 		return nil, err
	// 	}

	// 	arg, ok := args[0].(*Array)
	// 	if !ok {
	// 		return nil, slang.NewErrArgType(slang.ObjectArray, args[0].Type())
	// 	}

	// 	var ret slang.Object
	// 	switch method {
	// 	case slang.MethodEQ:
	// 		ret = NewArray(obj.Value == arg.Value)
	// 	case slang.MethodNEQ:
	// 		ret = NewArray(obj.Value != arg.Value)
	// 	}

	// 	return ret, nil
	// }

	err := slang.NewErrMethodUndefined(method, obj.Type())

	return nil, err
}
