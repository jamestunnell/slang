package objects

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type Null struct {
}

const (
	StrNULL   = "null"
	ClassNULL = "Null"
)

var null = &Null{}

var typeNULL = NewFundType(strNULL)

func NULL() *Null {
	return null
}

func (obj *Null) Get(name string) (slang.Object, bool) {
	return nil, false
}

func (obj *Null) Set(name string, val slang.Object) {
	// ignore set
}

func (obj *Null) Equal(other slang.Object) bool {
	_, ok := other.(*Null)

	return ok
}

func (obj *Null) Inspect() string {
	return StrNULL
}

// func (obj *Null) Truthy() bool {
// 	return false
// }

func (obj *Null) Send(methodName string, args ...slang.Object) (slang.Object, error) {
	// // I guess we could add instance methods to the null class?
	// if m, found := nullClass.GetInstanceMethod(methodName); found {
	// 	return m.Run(args)
	// }

	return nil, customerrs.NewErrMethodUndefined(methodName, ClassNULL)
}
