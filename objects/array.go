package objects

import (
	"errors"
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/jamestunnell/slang/types"
	"golang.org/x/exp/slices"
)

type Array struct {
	*Base

	Elements []slang.Object
}

const aryClassName = "Array"

var errArrayEmpty = errors.New("array is empty")
var aryClass *BuiltInClass

func init() {
	aryClass = NewBuiltInClass(
		types.NewPrimitiveType(aryClassName),
		map[string]slang.MethodFunc{
			slang.MethodFIRST: aryFIRST,
			slang.MethodLAST:  aryLAST,
			slang.MethodSIZE:  arySIZE,
			slang.MethodELEM:  aryELEM,
		},
	)
}

func NewArray(vals ...slang.Object) slang.Object {
	return &Array{
		Base:     NewBase(aryClass),
		Elements: vals,
	}
}

func (obj *Array) IsEqual(other slang.Object) bool {
	obj2, ok := other.(*Array)
	if !ok {
		return false
	}

	return slices.EqualFunc(obj.Elements, obj2.Elements, func(o1, o2 slang.Object) bool {
		return o1.IsEqual(o2)
	})
}

func (obj *Array) Inspect() string {
	return fmt.Sprintf("[]{%d elements}", len(obj.Elements))
}

func arySIZE(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	if err := CheckArgCount(args, 0); err != nil {
		return nil, err
	}

	return slang.Objects{NewInt(int64(len(obj.(*Array).Elements)))}, nil
}

func aryFIRST(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	if err := CheckArgCount(args, 0); err != nil {
		return nil, err
	}

	ary := obj.(*Array)

	if len(ary.Elements) == 0 {
		return nil, errArrayEmpty
	}

	return slang.Objects{ary.Elements[0]}, nil
}

func aryLAST(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	if err := CheckArgCount(args, 0); err != nil {
		return nil, err
	}

	ary := obj.(*Array)
	n := len(ary.Elements)

	if n == 0 {
		return nil, errArrayEmpty
	}

	return slang.Objects{ary.Elements[n-1]}, nil
}

func aryELEM(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	idx, err := CheckOneArg[*Int](args)
	if err != nil {
		return slang.Objects{}, err
	}

	ary := obj.(*Array)
	n := int64(len(ary.Elements))

	if idx.Value >= n {
		return nil, customerrs.NewErrArrayBounds(idx.Value, n)
	}

	return slang.Objects{ary.Elements[idx.Value]}, nil
}
