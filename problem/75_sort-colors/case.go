package sort_colors

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, []int] {
	return []tests.Case[[]int, []int]{
		{Input: []int{2, 0, 2, 1, 1, 0}, Except: []int{0, 0, 1, 1, 2, 2}},
		{Input: []int{2, 0, 1}, Except: []int{0, 1, 2}},
	}
}

type Func func(nums []int)

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in []int) (out []int) {
	return func(in []int) (out []int) {
		out = make([]int, len(in))
		copy(out, in)
		f(out)
		return
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
	tests.Unit(t, tests.Test[[]int, []int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[[]int, []int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
