package objects

import (
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/types"
)

type Bool struct {
	*Base

	Value bool
}

const boolClassName = "Bool"

var boolClass *BuiltInClass

func init() {
	boolClass = NewBuiltInClass(
		types.NewPrimitiveType(boolClassName),
		map[string]slang.MethodFunc{
			slang.MethodNOT: boolNOT,
			slang.MethodEQ:  boolEQ,
			slang.MethodNEQ: boolNEQ,
			slang.MethodAND: boolAND,
			slang.MethodOR:  boolOR,
		},
	)
}

var (
	objFalse = NewBool(false)
	objTrue  = NewBool(true)
)

func NewBool(val bool) slang.Object {
	return &Bool{
		Base:  NewBase(boolClass),
		Value: val,
	}
}

func TRUE() slang.Object {
	return objTrue
}

func FALSE() slang.Object {
	return objFalse
}

func (obj *Bool) IsEqual(other slang.Object) bool {
	obj2, ok := other.(*Bool)
	if !ok {
		return false
	}

	return obj.Value == obj2.Value
}

func (obj *Bool) Inspect() string {
	return strconv.FormatBool(obj.Value)
}

func boolNOT(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	if err := CheckArgCount(args, 0); err != nil {
		return nil, err
	}

	return slang.Objects{NewBool(!obj.(*Bool).Value)}, nil
}

func boolEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Bool](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Bool).Value == arg.Value)}, nil
}

func boolNEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Bool](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Bool).Value != arg.Value)}, nil
}

func boolAND(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Bool](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Bool).Value && arg.Value)}, nil
}

func boolOR(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	arg, err := CheckOneArg[*Bool](args)
	if err != nil {
		return slang.Objects{}, err
	}

	return slang.Objects{NewBool(obj.(*Bool).Value || arg.Value)}, nil
}
