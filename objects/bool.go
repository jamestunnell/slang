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

type boolBinOp int

const (
	boolClassName = "Bool"

	boolOpEQ boolBinOp = iota
	boolOpNEQ
	boolOpAND
	boolOpOR
)

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
	return slang.Objects{NewBool(!obj.(*Bool).Value)}, nil
}

func boolEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return boolBinaryOp(boolOpEQ, obj, args)
}

func boolNEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return boolBinaryOp(boolOpNEQ, obj, args)
}

func boolAND(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return boolBinaryOp(boolOpAND, obj, args)
}

func boolOR(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return boolBinaryOp(boolOpOR, obj, args)
}

func boolBinaryOp(
	op boolBinOp,
	obj slang.Object,
	args slang.Objects,
) (slang.Objects, error) {
	left := obj.(*Bool)

	if err := checkArgCount(args, 1); err != nil {
		return slang.Objects{}, err
	}

	right, err := CheckType[*Bool](args[0])
	if err != nil {
		return slang.Objects{}, err
	}

	var ret slang.Object

	switch op {
	case boolOpEQ:
		ret = NewBool(left.Value == right.Value)
	case boolOpNEQ:
		ret = NewBool(left.Value != right.Value)
	case boolOpAND:
		ret = NewBool(left.Value && right.Value)
	case boolOpOR:
		ret = NewBool(left.Value || right.Value)
	}

	return slang.Objects{ret}, nil
}
