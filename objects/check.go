package objects

import (
	"reflect"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

func CheckType[T slang.Object](obj slang.Object) (T, error) {
	var t T

	var ok bool

	t, ok = obj.(T)
	if !ok {
		err := customerrs.NewErrObjectType(
			reflect.TypeOf(t).String(), reflect.TypeOf(obj).String())

		return t, err
	}

	return t, nil
}
