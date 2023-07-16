package tests

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
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

type (
	Funcs[T, U any] struct {
		solution func(in T) (out U)
		Cases    func() []Case[T, U]
		IsEqual  func(a, b U) bool
	}

	Test[T, U any] struct {
		Solution []function[T, U]
		Cases    func() []Case[T, U]
		IsEqual  func(a, b U) bool
	}
)

func IsFail[T, U any](w io.Writer, s Funcs[T, U]) (fail bool) {
	for _, _case := range s.Cases() {
		var (
			except = _case.Except
			input  = _case.Input
			output = s.solution(input)
		)

		isEqual := s.IsEqual
		if isEqual == nil {
			isEqual = Equal[U]
		}

		equal := isEqual(except, output)

		var (
			colors = textBlack
			bg     = bgDefault
			flag   = flagDefault
		)
		if !equal {
			colors = textRed
			flag = flagUnderline
			fail = true
		}

		list := []struct {
			name  string
			value any
		}{
			{name: "case  ", value: input},
			{name: "except", value: except},
			{name: "output", value: output},
		}

		w.Write([]byte{'\n'})
		for _, item := range list {
			_, err := fmt.Fprintf(w, "\033[%d;%d;%dm%+v: %+v\033[0m\n", flag, bg, colors, item.name, item.value)
			if err != nil {
				panic(err)
			}
		}
	}

	return fail
}

func Unit[T any, U any](t *testing.T, s Test[T, U]) {
	for _, f := range s.Solution {
		t.Run(f.Name(), func(t *testing.T) {
			buf := new(bytes.Buffer)
			fail := IsFail(buf, Funcs[T, U]{
				solution: f.Func(),
				Cases:    s.Cases,
				IsEqual:  s.IsEqual,
			})

			show := t.Log
			if fail {
				show = t.Error
			}

			// TODO 打印函数真实位置
			show(buf)
		})
		fmt.Println()
	}
}

func Bench[T any, U any](b *testing.B, s Test[T, U]) {
	for _, f := range s.Solution {
		var unit func()

		// cache
		cache := s.Cases()
		b.Run(f.Name(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, c := range cache {
					_ = f.Func()(c.Input)
				}
			}
			unit = func() {
				buf := new(bytes.Buffer)
				fail := IsFail(buf, Funcs[T, U]{
					solution: f.Func(),
					Cases:    s.Cases,
					IsEqual:  s.IsEqual,
				})
				if fail {
					b.Error(buf)
				}
			}
		})

		unit()
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
		functions = append(functions,
			newFunc(adaptor(f),
				WithName[T, U](FuncName(f)),
			),
		)
	}
	return functions
}

func Equal[T any](except, output T) bool {
	return fmt.Sprintf("%+v", except) == fmt.Sprintf("%+v", output)

	return reflect.DeepEqual(except, output)
}
