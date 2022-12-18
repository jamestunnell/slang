package objects

import "github.com/jamestunnell/slang"

type Null struct {
}

const StrNULL = "null"

var null = &Null{}

func NULL() *Null {
	return null
}

func (obj *Null) Inspect() string {
	return StrNULL
}

func (obj *Null) Truthy() bool {
	return false
}

func (obj *Null) Type() slang.ObjectType {
	return slang.ObjectNULL
}

func (obj *Null) Send(method string, arg ...slang.Object) (slang.Object, error) {
	return nil, slang.NewErrMethodUndefined(method, slang.ObjectNULL)
}
