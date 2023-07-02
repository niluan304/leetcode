package tests

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

const (
	textBlack = 30
	textRed   = 31
)

const (
	bgDefault = 30
	bgGreen   = 42
	bgWhite   = 47
)

const (
	flagDefault     = 0
	flagHeightLight = 1
	flagUnderline   = 4
)

func Sign[T any, U any](show bool, f func(in T) (out U), cases ...Case[T, U]) (err error) {
	var errs []error
	for _, e := range cases {
		err = sign(show, f, e)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return JoinError(errs...)
	}

	return nil
}

func sign[T any, U any](show bool, f func(in T) (out U), c Case[T, U]) (err error) {
	output := f(c.Input)
	out, err := json.Marshal(output)
	if err != nil {
		return err
	}
	except, err := json.Marshal(c.Except)
	if err != nil {
		return err
	}

	var (
		colors = textBlack
		bg     = bgDefault
		flag   = flagDefault
	)

	equal := string(out) == string(except)
	if !equal {
		colors = textRed
		flag = flagUnderline
		show = true

		err = fmt.Errorf("\ncase  : %s\nexcept: %s\noutput: %s", c, except, out)
	}

	ret := Result[T, U]{
		Case:   c,
		Output: output,
	}
	if show {
		fmt.Printf("\033[%d;%d;%dm%+v\033[0m\n", flag, bg, colors, ret)
	}

	return err
}

func Once[T any, U any](f func(in T) (out U), cases ...Case[T, U]) (err error) {
	return Sign(true, f, cases...)
}

func Show[T any, U any](f func(in T) (out U), cases ...Case[T, U]) (err error) {
	return Sign(false, f, cases...)
}

func Run[T any, U any](f func(in T) (out U), cases ...Case[T, U]) {
	for _, c := range cases {
		f(c.Input)
	}
}

func Unit[T any, U any](t *testing.T, cases func() []Case[T, U], fs ...function[T, U]) {
	for _, f := range fs {
		t.Run(f.Name(), func(t *testing.T) {
			err := Once(f.Func(), cases()...)
			if err != nil {
				t.Errorf("%+v", err)
			}
		})
		fmt.Println()
	}
}

func Bench[T any, U any](b *testing.B, cases func() []Case[T, U], fs ...function[T, U]) {

	for _, f := range fs {
		err := Show(f.Func(), cases()...)
		if err != nil {
			b.Errorf("%+v", err)
		}

		// cache
		cache := cases()
		b.Run(f.Name(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Run(f.Func(), cache...)
			}
		})
		fmt.Println()
	}
}

func NewFunc[T, U any](fs ...func(in T) (out U)) (functions []function[T, U]) {
	for _, f := range fs {
		functions = append(functions, newFunc(f))
	}
	return functions
}

func NewFuncWithAdaptor[T, U, F any](adaptor func(F) func(in T) (out U),
	fs ...F) (functions []function[T, U]) {
	for _, f := range fs {
		functions = append(functions, newFunc(adaptor(f)))
	}
	return functions
}
