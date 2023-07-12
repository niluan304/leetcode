package two_sum

import (
	"sort"
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums   []int
	target int
}

type Output []int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		// TODO add question case
		{
			Input:  Input{nums: []int{2, 7, 11, 15}, target: 9},
			Except: []int{0, 1},
		},
		{
			Input:  Input{nums: []int{3, 2, 4}, target: 6},
			Except: []int{1, 2},
		},
		{
			Input:  Input{nums: []int{3, 3}, target: 6},
			Except: []int{0, 1},
		},
	}
}

type Func func(nums []int, target int) []int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.nums,
			in.target,
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

func Equal(a, b Output) bool {
	sort.Ints(a)
	sort.Ints(b)
	return tests.Equal(a, b)
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
