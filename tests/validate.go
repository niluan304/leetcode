package tests

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/sourcegraph/conc"
)

func Validate[T, U any](input T, except U, fs ...func(T) U) (err error) {
	wg := conc.NewWaitGroup()
	defer wg.Wait()

	for _, f := range fs {
		wg.Go(func() {
			actual := f(input)
			if !reflect.DeepEqual(actual, except) {
				msg := fmt.Sprintf("f:%s, input:%+v, except:%v, actual:%v",
					FuncName(f), input, except, actual)
				err = errors.Join(err, errors.New(msg))
			}
		})
	}
	return err
}

func Validate2[T1, T2, U any](input1 T1, input2 T2, except U, fs ...func(T1, T2) U) (err error) {
	for _, f := range fs {
		e := Validate([]any{input1, input2}, except, func(_ []any) U {
			return f(input1, input2)
		})
		if e != nil {
			e = errors.Join(errors.New("funcName: "+FuncName(f)), e)
			err = errors.Join(err, e)
		}
	}

	return err
}

func Validate3[T1, T2, T3, U any](input1 T1, input2 T2, input3 T3, except U, fs ...func(T1, T2, T3) U) (err error) {
	for _, f := range fs {
		e := Validate([]any{input1, input2, input3}, except, func(_ []any) U {
			return f(input1, input2, input3)
		})
		if e != nil {
			e = errors.Join(errors.New("funcName: "+FuncName(f)), e)
			err = errors.Join(err, e)
		}
	}

	return err
}
