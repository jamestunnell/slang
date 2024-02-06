package objects

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/jamestunnell/slang"
// 	"github.com/jamestunnell/slang/customerrs"
// )

// type CallFn func(args ...Object) (Object, error)

// type BuiltInFn struct {
// 	Name   string
// 	Fn     CallFn
// 	Params []string
// }

// const ClassBUILTINFN = "BuiltInFn"

// var (
// 	builtinFns     = map[string]*BuiltInFn{}
// 	builtinFnClass = NewBuiltInClass(ClassBUILTINFN)
// )

// func init() {
// 	builtinFns["puts"] = NewBuiltInFn("puts", puts, "...")
// }

// func FindBuiltInFn(name string) (*BuiltInFn, bool) {
// 	if fn, found := builtinFns[name]; found {
// 		return fn, true
// 	}

// 	return nil, false
// }

// func NewBuiltInFn(name string, fn CallFn, params ...string) *BuiltInFn {
// 	return &BuiltInFn{
// 		Name:   name,
// 		Fn:     fn,
// 		Params: params,
// 	}
// }

// func (obj *BuiltInFn) Class() Class {
// 	return builtinFnClass
// }

// func (obj *BuiltInFn) Inspect() string {
// 	paramsStr := strings.Join(obj.Params, ", ")
// 	return fmt.Sprintf("%s(%s){...}", obj.Name, paramsStr)
// }

// func (obj *BuiltInFn) Truthy() bool {
// 	return true
// }

// func (obj *BuiltInFn) Send(methodName string, args ...Object) (Object, error) {
// 	// an added instance method would override a standard one
// 	if m, found := builtinFnClass.GetInstanceMethod(methodName); found {
// 		return m.Run(args)
// 	}

// 	switch methodName {
// 	case slang.MethodCALL:
// 		return obj.Fn(args...)
// 	}

// 	err := customerrs.NewErrMethodUndefined(methodName, ClassBUILTINFN)

// 	return nil, err
// }

// func puts(args ...Object) (Object, error) {
// 	for _, arg := range args {
// 		fmt.Println(arg.Inspect())
// 	}

// 	return null, nil
// }
