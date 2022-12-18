package objects

import (
	"strconv"

	"github.com/jamestunnell/slang"
)

type Bool struct {
	Value bool
}

func NewBool(val bool) slang.Object {
	return &Bool{Value: val}
}

func (obj *Bool) Inspect() string {
	return strconv.FormatBool(obj.Value)
}

func (obj *Bool) Truthy() bool {
	return obj.Value
}

func (obj *Bool) Type() slang.ObjectType {
	return slang.ObjectBOOL
}

func (obj *Bool) Send(method string, args ...slang.Object) (slang.Object, error) {
	switch method {
	case slang.MethodNOT:
		return NewBool(!obj.Value), nil
	case slang.MethodEQ, slang.MethodNEQ:
		if err := checkArgCount(args, 1); err != nil {
			return nil, err
		}

		arg, ok := args[0].(*Bool)
		if !ok {
			return nil, slang.NewErrArgType(slang.ObjectBOOL, args[0].Type())
		}

		var ret slang.Object
		switch method {
		case slang.MethodEQ:
			ret = NewBool(obj.Value == arg.Value)
		case slang.MethodNEQ:
			ret = NewBool(obj.Value != arg.Value)
		}

		return ret, nil
	}

	err := slang.NewErrMethodUndefined(method, obj.Type())

	return nil, err
}
