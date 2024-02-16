package objects

import (
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/types"
)

type String struct {
	*Base

	Value string
}

const strClassName = "Bool"

var strClass *BuiltInClass

func init() {
	strClass = NewBuiltInClass(
		types.NewPrimitiveType(strClassName),
		map[string]slang.MethodFunc{
			slang.MethodSIZE: strSIZE,
			slang.MethodADD:  strADD,
			slang.MethodEQ:   strEQ,
			slang.MethodNEQ:  strNEQ,
			slang.MethodLT:   strLT,
			slang.MethodLEQ:  strLEQ,
			slang.MethodGT:   strGT,
			slang.MethodGEQ:  strGEQ,
		},
	)
}

func NewString(val string) slang.Object {
	return &String{
		Base:  NewBase(strClass),
		Value: val,
	}
}

func (obj *String) Inspect() string {
	return fmt.Sprintf("\"%s\"", obj.Value)
}

func (obj *String) IsEqual(other slang.Object) bool {
	obj2, ok := other.(*String)
	if !ok {
		return false
	}

	return obj.Value == obj2.Value
}

func strSIZE(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	if err := CheckArgCount(args, 0); err != nil {
		return nil, err
	}

	return slang.Objects{NewInt(int64(len(obj.(*String).Value)))}, nil
}

func strADD(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*String](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewString(obj.(*String).Value + arg.Value)}, nil
}

func strEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*String](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*String).Value == arg.Value)}, nil
}

func strNEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*String](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*String).Value != arg.Value)}, nil
}

func strLT(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*String](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*String).Value < arg.Value)}, nil
}

func strLEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*String](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*String).Value <= arg.Value)}, nil
}

func strGT(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*String](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*String).Value > arg.Value)}, nil
}

func strGEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*String](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*String).Value >= arg.Value)}, nil
}
