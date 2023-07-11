package longest_consecutive_sequence

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums []int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{Input: Input{nums: []int{}}, Except: 0},
		{Input: Input{nums: []int{3}}, Except: 1},
		{Input: Input{nums: []int{100, 4, 200, 1, 3, 2}}, Except: 4},
		{Input: Input{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}}, Except: 9},
		{Input: Input{nums: []int{1, 2, 0, 1}}, Except: 3},
	}
}

type Func func(nums []int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.nums,
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

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
