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

type strBinOp int

const (
	strClassName = "Bool"

	strOpADD strBinOp = iota
	strOpEQ
	strOpNEQ
	strOpLT
	strOpLEQ
	strOpGT
	strOpGEQ
)

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
	return slang.Objects{NewInt(int64(len(obj.(*String).Value)))}, nil
}

func strADD(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return strBinaryOp(strOpADD, obj, args)
}

func strEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return strBinaryOp(strOpEQ, obj, args)
}

func strNEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return strBinaryOp(strOpNEQ, obj, args)
}

func strLT(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return strBinaryOp(strOpLT, obj, args)
}

func strLEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return strBinaryOp(strOpLEQ, obj, args)
}

func strGT(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return strBinaryOp(strOpGT, obj, args)
}

func strGEQ(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return strBinaryOp(strOpGEQ, obj, args)
}

func strBinaryOp(
	op strBinOp,
	obj slang.Object,
	args slang.Objects,
) (slang.Objects, error) {
	left := obj.(*String)

	if err := checkArgCount(args, 1); err != nil {
		return slang.Objects{}, err
	}

	right, err := CheckType[*String](args[0])
	if err != nil {
		return slang.Objects{}, err
	}

	var ret slang.Object

	switch op {
	case strOpADD:
		ret = NewString(left.Value + right.Value)
	case strOpEQ:
		ret = NewBool(left.Value == right.Value)
	case strOpNEQ:
		ret = NewBool(left.Value != right.Value)
	case strOpLT:
		ret = NewBool(left.Value < right.Value)
	case strOpLEQ:
		ret = NewBool(left.Value <= right.Value)
	case strOpGT:
		ret = NewBool(left.Value > right.Value)
	case strOpGEQ:
		ret = NewBool(left.Value >= right.Value)
	}

	return slang.Objects{ret}, nil
}
