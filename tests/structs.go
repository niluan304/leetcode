package tests

import (
	"reflect"
	"runtime"
)

type Case[T any, U any] struct {
	Input  T
	Except U
}

type Result[T any, U any] struct {
	//Equal bool
	Case[T, U]
	Output U
}

func newFunc[T any, U any](f func(in T) (out U)) function[T, U] {
	return function[T, U]{
		function: f,
		name:     FuncName(f),
	}
}

type function[T any, U any] struct {
	function func(in T) (out U)
	name     string
}

func (f function[T, U]) Name() (name string) {
	return f.name
}

func (f function[T, U]) Func() func(in T) (out U) {
	return f.function
}

func FuncName(entityFunc any) string {
	return runtime.FuncForPC(reflect.ValueOf(entityFunc).Pointer()).Name()
}
