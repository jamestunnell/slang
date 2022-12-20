package objects

import (
	"errors"

	"github.com/jamestunnell/slang"
)

type Array struct {
	Elements []slang.Object
}

var errArrayEmpty = errors.New("array is empty")

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
	switch method {
	case slang.MethodFIRST, slang.MethodLAST:
		if err := checkArgCount(args, 0); err != nil {
			return nil, err
		}

		if len(obj.Elements) == 0 {
			return nil, errArrayEmpty
		}

		switch method {
		case slang.MethodFIRST:
			return obj.Elements[0], nil
		case slang.MethodLAST:
			return obj.Elements[len(obj.Elements)], nil
		}
	case slang.MethodINDEX:
		if err := checkArgCount(args, 1); err != nil {
			return nil, err
		}

		return obj.Index(args[0])
	}

	err := slang.NewErrMethodUndefined(method, obj.Type())

	return nil, err
}

func (obj *Array) Index(arg slang.Object) (slang.Object, error) {
	idx, ok := arg.(*Integer)
	if !ok {
		return nil, slang.NewErrArgType(slang.ObjectINTEGER, arg.Type())
	}

	if int(idx.Value) >= len(obj.Elements) {
		return nil, slang.NewErrArrayBounds(idx.Value, int64(len(obj.Elements)))
	}

	return obj.Elements[idx.Value], nil
}
