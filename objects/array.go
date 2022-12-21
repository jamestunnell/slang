package objects

import (
	"errors"

	"github.com/jamestunnell/slang"
)

type Array struct {
	Elements []slang.Object
	methods  map[string]*slang.Method
}

const ParamINDEX = "index"

var errArrayEmpty = errors.New("array is empty")

func NewArray(vals ...slang.Object) slang.Object {
	ary := &Array{
		Elements: vals,
		methods:  map[string]*slang.Method{},
	}

	ary.methods[slang.MethodFIRST] = slang.NewMethod(ary.first)
	ary.methods[slang.MethodLAST] = slang.NewMethod(ary.last)
	ary.methods[slang.MethodSIZE] = slang.NewMethod(ary.size)
	ary.methods[slang.MethodAT] = slang.NewMethod(ary.at, ParamINDEX)

	return ary
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

func (obj *Array) Methods() map[string]*slang.Method {
	return obj.methods
}

func (obj *Array) size(env *slang.Environment) (slang.Object, error) {
	return NewInteger(int64(len(obj.Elements))), nil
}

func (obj *Array) first(env *slang.Environment) (slang.Object, error) {

	if len(obj.Elements) == 0 {
		return nil, errArrayEmpty
	}

	return obj.Elements[0], nil
}

func (obj *Array) last(env *slang.Environment) (slang.Object, error) {
	if len(obj.Elements) == 0 {
		return nil, errArrayEmpty
	}

	return obj.Elements[len(obj.Elements)-1], nil
}

func (obj *Array) at(env *slang.Environment) (slang.Object, error) {
	idx, err := GetInt(env, ParamINDEX)
	if err != nil {
		return NULL(), err
	}

	n := int64(len(obj.Elements))
	if idx.Value >= n {
		return nil, slang.NewErrArrayBounds(idx.Value, n)
	}

	return obj.Elements[idx.Value], nil
}
