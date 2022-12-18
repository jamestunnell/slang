package objects

import (
	"strconv"

	"github.com/jamestunnell/slang"
)

type Integer struct {
	Value int64
}

func NewInteger(val int64) slang.Object {
	return &Integer{Value: val}
}

func (obj *Integer) Inspect() string {
	return strconv.FormatInt(obj.Value, 10)
}

func (obj *Integer) Truthy() bool {
	return true
}

func (obj *Integer) Type() slang.ObjectType {
	return slang.ObjectINTEGER
}

func (obj *Integer) Send(method string, args ...slang.Object) (slang.Object, error) {
	switch method {
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

		if args[0].Type() == slang.ObjectFLOAT {
			return NewFloat(float64(obj.Value)).Send(method, args...)
		}

		return obj.sendOne(method, args[0])
	}

	err := slang.NewErrMethodUndefined(method, obj.Type())

	return nil, err
}

func (obj *Integer) sendOne(method string, arg slang.Object) (slang.Object, error) {
	flt, ok := arg.(*Integer)
	if !ok {
		return nil, slang.NewErrArgType(slang.ObjectFLOAT, arg.Type())
	}

	var ret slang.Object

	switch method {
	case slang.MethodADD:
		ret = NewInteger(obj.Value + flt.Value)
	case slang.MethodSUB:
		ret = NewInteger(obj.Value - flt.Value)
	case slang.MethodMUL:
		ret = NewInteger(obj.Value * flt.Value)
	case slang.MethodDIV:
		ret = NewInteger(obj.Value / flt.Value)
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

func intAbs(val int64) int64 {
	if val > 0 {
		return val
	}

	return -val
}

func checkArgCount(args []slang.Object, count int) error {
	if len(args) != count {
		return slang.NewErrArgCount(count, len(args))
	}

	return nil
}
