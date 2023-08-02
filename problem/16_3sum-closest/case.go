package three_sum_closest

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums   []int
	target int
}

var cases = func() []tests.Case[Input, int] {
	return []tests.Case[Input, int]{
		{
			Input:  Input{nums: []int{-1, 2, 1, -4}, target: 1},
			Except: 2,
		},
		{
			Input:  Input{nums: []int{0, 0, 0}, target: 1},
			Except: 0,
		},
	}
}

type Func func(nums []int, target int) int

func adaptor(f Func) func(in Input) (out int) {
	return func(in Input) (out int) {
		return f(in.nums, in.target)
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[Input, int]) {
	_cases := cases()
	cases = func() []tests.Case[Input, int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[Input, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
