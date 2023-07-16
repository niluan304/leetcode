package subsets_ii

import (
	"sort"
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums []int
}

type Output [][]int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input:  Input{nums: []int{1, 2, 2}},
			Except: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			Input:  Input{nums: []int{0}},
			Except: [][]int{{}, {0}},
		},
	}
}

type Func func(nums []int) [][]int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.nums,
		))
	}
}

func Sort(x Output) {
	sort.Slice(x, func(i, j int) bool {
		if len(x[i]) != len(x[j]) {
			return len(x[i]) < len(x[j])
		}

		for idx, v := range x[i] {
			if x[j][idx] != v {
				return v < x[j][idx]
			}
		}

		return i < j
	})
}

func Equal(x, y Output) bool {
	if len(x) != len(y) {
		return false
	}

	Sort(x)
	Sort(y)

	for i, v := range x {
		if !tests.Equal(v, y[i]) {
			return false
		}
	}

	return true
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
