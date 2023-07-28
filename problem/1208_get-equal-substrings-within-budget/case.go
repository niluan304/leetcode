package get_equal_substrings_within_budget

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	s       string
	t       string
	maxCost int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				s:       "abcd",
				t:       "bcdf",
				maxCost: 3,
			},
			Except: 3,
		},
		{
			Input: Input{
				s:       "abcd",
				t:       "cdef",
				maxCost: 3,
			},
			Except: 1,
		},
		{
			Input: Input{
				s:       "abcd",
				t:       "acde",
				maxCost: 0,
			},
			Except: 1,
		},
	}
}

type Func func(s string, t string, maxCost int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.s,
			in.t,
			in.maxCost,
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
