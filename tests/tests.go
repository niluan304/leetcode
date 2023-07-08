package tests

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

// 前景 颜色
//
//	30 黑色
//	31 红色
//	32 绿色
//	33 黄色
//	34 蓝色
//	35 紫红色
//	36 青蓝色
//	37 白色
const (
	textBlack = 30
	textRed   = 31
)

// 背景 颜色
//
//	40 黑色
//	41 红色
//	42 绿色
//	43 黄色
//	44 蓝色
//	45 紫红色
//	46 青蓝色
//	47 白色
const (
	bgDefault = 30
	bgGreen   = 42
	bgWhite   = 47
)

// 代码 意义
//
//	0  终端默认设置
//	1  高亮显示
//	4  使用下划线
//	5  闪烁
//	7  反白显示
//	8  不可见
const (
	flagDefault     = 0
	flagHeightLight = 1
	flagUnderline   = 4
)

func Sign[T any, U any](w io.Writer, f func(in T) (out U), cases ...Case[T, U]) (fail bool) {
	for _, e := range cases {
		w.Write([]byte{'\n'})

		equal := sign(w, f, e)
		if !equal {
			fail = true
		}
	}

	return fail
}

func sign[T any, U any](w io.Writer, f func(in T) (out U), c Case[T, U]) (equal bool) {
	output := f(c.Input)

	out := fmt.Sprintf("%+v", output)
	except := fmt.Sprintf("%+v", c.Except)

	var (
		colors = textBlack
		bg     = bgDefault
		flag   = flagDefault
	)

	equal = out == except
	if !equal {
		colors = textRed
		flag = flagUnderline
	}

	list := []struct {
		name  string
		value any
	}{
		{name: "case  ", value: c.Input},
		{name: "except", value: except},
		{name: "output", value: out},
	}

	for _, item := range list {
		_, err := fmt.Fprintf(w, "\033[%d;%d;%dm%+v: %+v\033[0m\n", flag, bg, colors, item.name, item.value)
		if err != nil {
			panic(err)
		}
	}

	return equal
}

func Msg[T any, U any](f func(in T) (out U), cases []Case[T, U]) (r io.Reader, fail bool) {
	buf := new(bytes.Buffer)
	fail = Sign(buf, f, cases...)

	return buf, fail
}

func Unit[T any, U any](t *testing.T, cases func() []Case[T, U], fs ...function[T, U]) {
	for _, f := range fs {
		t.Run(f.Name(), func(t *testing.T) {
			msg, fail := Msg(f.Func(), cases())

			show := t.Log
			if fail {
				show = t.Error
			}

			show(msg)
		})
		fmt.Println()
	}
}

func Bench[T any, U any](b *testing.B, cases func() []Case[T, U], fs ...function[T, U]) {
	for _, f := range fs {
		// cache
		cache := cases()
		b.Run(f.Name(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, c := range cache {
					_ = f.Func()(c.Input)
				}
			}

			msg, fail := Msg(f.Func(), cases())
			if fail {
				b.Error(msg)
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
