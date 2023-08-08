package median_of_two_sorted_arrays

import (
	"testing"

	"leetcode/tests"
)

type Input struct{ nums1, nums2 []int }

type Output float64

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			Except: 2.0,
		},
		{
			Input: Input{
				nums1: []int{2, 3},
				nums2: []int{}},
			Except: 2.5,
		},
	}
}

type Func func(nums1 []int, nums2 []int) float64

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) (out Output) {
	return func(in Input) (out Output) {
		return Output(f(in.nums1, in.nums2))
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

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
