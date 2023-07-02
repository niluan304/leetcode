package plus_one

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, []int] {
	return []tests.Case[[]int, []int]{
		{Input: []int{1, 2, 3}, Except: []int{1, 2, 4}},
		{Input: []int{4, 3, 2, 1}, Except: []int{4, 3, 2, 2}},
		{Input: []int{0}, Except: []int{1}},
	}
}

type Func func(digits []int) []int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in []int) (out []int) {
	return func(in []int) (out []int) {
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
