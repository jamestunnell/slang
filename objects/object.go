package objects

import "github.com/jamestunnell/slang"

type Object struct {
	Type slang.ObjectType
	Core Core
}

type Core interface {
	IsEqual(Core) bool
	Inspect() string
	Send(name string, args ...slang.Object) (slang.Object, error)
}

func (obj *Object) GetType() slang.ObjectType {
	return obj.Type
}

func (obj *Object) IsEqual(other slang.Object) bool {
	obj2, ok := other.(*Object)
	if !ok {
		return false
	}

	if obj.Type != obj2.Type {
		return false
	}

	return obj.Core.IsEqual(obj2.Core)
}

func (obj *Object) Inspect() string {
	return obj.Core.Inspect()
}

func (obj *Object) Send(name string, args ...slang.Object) (slang.Object, error) {
	return obj.Core.Send(name, args...)
}

// type (
// 	Object interface {
// 		Class() Class
// 		Inspect() string
// 		Truthy() bool
// 		// Type() ObjectType
// 		Send(method string, args ...Object) (Object, error)
// 	}

// )
