package objects

import (
	"reflect"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

func CheckOneArg[T slang.Object](args slang.Objects) (T, error) {
	var t T

	var err error

	if err = CheckArgCount(args, 1); err != nil {
		return t, err
	}

	t, err = CheckType[T](args[0])
	if err != nil {
		return t, err
	}

	return t, nil
}

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

func CheckArgCount(args slang.Objects, count int) error {
	if len(args) != count {
		return customerrs.NewErrArgCount(count, len(args))
	}

	return nil
}
