package permutations_ii

import (
	"reflect"
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
		// TODO add question case
		{
			Input:  Input{nums: []int{1, 1, 2}},
			Except: Output{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			Input:  Input{nums: []int{1, 2, 3}},
			Except: Output{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
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

func AddCases(c func() []tests.Case[Input, Output]) {
	_cases := cases()
	cases = func() []tests.Case[Input, Output] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Sort(x Output) {
	sort.Slice(x, func(i, j int) bool {
		xi, xj := x[i], x[j]
		for k, v := range xi {
			if v != xj[k] {
				return v < xj[k]
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
