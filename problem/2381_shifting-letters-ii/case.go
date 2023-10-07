package shifting_letters_ii

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	s      string
	shifts [][]int
}

type Output string

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				s:      "abc",
				shifts: [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}},
			},
			Except: "ace",
		},
		{
			Input: Input{
				s:      "dztz",
				shifts: [][]int{{0, 0, 0}, {1, 1, 1}},
			},
			Except: "catz",
		},
	}
}

type Func func(s string, shifts [][]int) string

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.s,
			in.shifts,
		))
	}
}

func AddCases(c func() []tests.Case[Input, Output]) {
	_cases := cases()
	cases = func() []tests.Case[Input, Output] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Equal(x, y Output) bool {

	return reflect.DeepEqual(x, y)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}
