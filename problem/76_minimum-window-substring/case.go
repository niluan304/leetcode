package minimum_window_substring

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	s string
	t string
}

type Output string

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			Except: "BANC",
		},
		{
			Input: Input{
				s: "a",
				t: "a",
			},
			Except: "a",
		},
		{
			Input: Input{
				s: "a",
				t: "aa",
			},
			Except: "",
		},
	}
}

type Func func(s string, t string) string

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.s,
			in.t,
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
