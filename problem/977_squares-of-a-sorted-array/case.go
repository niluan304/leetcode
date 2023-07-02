package squares_of_a_sorted_array

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, []int] {
	return []tests.Case[[]int, []int]{
		{Input: []int{-4, -1, 0, 3, 10}, Except: []int{0, 1, 9, 16, 100}},
		{Input: []int{-7, -3, 2, 3, 11}, Except: []int{4, 9, 9, 49, 121}},
	}
}

type Func func([]int) []int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in []int) (out []int) {
	return func(in []int) (out []int) {
		if !checkResult {
			return f(in)
		}

		out = make([]int, len(in))
		copy(out, in)
		return f(out)
	}
}

func AddCases(c func() []tests.Case[[]int, []int]) {
	_cases := cases()
	cases = func() []tests.Case[[]int, []int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
