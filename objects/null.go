package objects

import "github.com/jamestunnell/slang"

type Null struct {
}

const (
	StrNULL   = "null"
	ClassNULL = "Null"
)

var null = &Null{}

var nullClass = NewBuiltInClass(ClassNULL)

func NULL() *Null {
	return null
}

func (obj *Null) Class() slang.Class {
	return nullClass
}

func (obj *Null) Inspect() string {
	return StrNULL
}

func (obj *Null) Truthy() bool {
	return false
}

func (obj *Null) Send(methodName string, args ...slang.Object) (slang.Object, error) {
	// I guess we could add instance methods to the null class?
	if m, found := nullClass.GetInstanceMethod(methodName); found {
		return m.Run(args)
	}

	return nil, slang.NewErrMethodUndefined(methodName, ClassNULL)
}
