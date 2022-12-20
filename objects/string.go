package objects

import (
	"github.com/jamestunnell/slang"
)

type String struct {
	Value string
}

func NewString(val string) slang.Object {
	return &String{Value: val}
}

func (obj *String) Inspect() string {
	return obj.Value
}

func (obj *String) Truthy() bool {
	return true
}

func (obj *String) Type() slang.ObjectType {
	return slang.ObjectSTRING
}

func (obj *String) Send(method string, args ...slang.Object) (slang.Object, error) {
	switch method {
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

		return obj.sendOne(method, args[0])
	}

	err := slang.NewErrMethodUndefined(method, obj.Type())

	return nil, err
}

func (obj *String) sendOne(method string, arg slang.Object) (slang.Object, error) {
	flt, ok := arg.(*String)
	if !ok {
		return nil, slang.NewErrArgType(slang.ObjectSTRING, arg.Type())
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
