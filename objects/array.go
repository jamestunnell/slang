package objects

import (
	"errors"

	"github.com/jamestunnell/slang"
)

type Array struct {
	Elements []slang.Object
}

const ClassARRAY = "Array"

var (
	errArrayEmpty = errors.New("array is empty")
	aryClass      = NewBuiltInClass(ClassARRAY)
)

func NewArray(vals ...slang.Object) slang.Object {
	return &Array{Elements: vals}
}

func (obj *Array) Class() slang.Class {
	return aryClass
}

func (obj *Array) Inspect() string {
	return "[...]"
}

func (obj *Array) Truthy() bool {
	return true
}

func (obj *Array) Send(methodName string, args ...slang.Object) (slang.Object, error) {
	// an added instance method would override a standard one
	if m, found := aryClass.GetInstanceMethod(methodName); found {
		return m.Run(args)
	}

	switch methodName {
	case slang.MethodFIRST, slang.MethodLAST, slang.MethodSIZE:
		if err := checkArgCount(args, 0); err != nil {
			return nil, err
		}

		if len(obj.Elements) == 0 {
			return nil, errArrayEmpty
		}

		switch methodName {
		case slang.MethodFIRST:
			return obj.Elements[0], nil
		case slang.MethodLAST:
			return obj.Elements[len(obj.Elements)-1], nil
		case slang.MethodSIZE:
			return NewInteger(int64(len(obj.Elements))), nil
		}
	case slang.MethodINDEX:
		if err := checkArgCount(args, 1); err != nil {
			return nil, err
		}

		return obj.Index(args[0])
	}

	err := slang.NewErrMethodUndefined(methodName, ClassARRAY)

	return nil, err
}

func (obj *Array) Index(arg slang.Object) (slang.Object, error) {
	idx, ok := arg.(*Integer)
	if !ok {
		return nil, slang.NewErrArgType(ClassINTEGER, arg.Class().Name())
	}

	if int(idx.Value) >= len(obj.Elements) {
		return nil, slang.NewErrArrayBounds(idx.Value, int64(len(obj.Elements)))
	}

	return obj.Elements[idx.Value], nil
}
