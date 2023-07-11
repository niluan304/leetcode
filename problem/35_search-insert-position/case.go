package search_insert_position

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
		{Input: Input{nums: []int{1, 3, 5, 6}, target: 5}, Except: 2},
		{Input: Input{nums: []int{1, 3, 5, 6}, target: 2}, Except: 1},
		{Input: Input{nums: []int{1, 3, 5, 6}, target: 7}, Except: 4},
	}
}

type Func func(nums []int, target int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) (out int) {
	return func(in Input) (out int) {
		return f(in.nums, in.target)
	}
}

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

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[Input, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
