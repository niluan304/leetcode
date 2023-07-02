package median_of_two_sorted_arrays

import (
	"testing"

	"leetcode/tests"
)

type Input struct{ Nums1, Nums2 []int }

var cases = func() []tests.Case[Input, float64] {
	return []tests.Case[Input, float64]{
		{Input: Input{Nums1: []int{1, 3}, Nums2: []int{2}}, Except: 2.0},
		{Input: Input{Nums1: []int{1, 2}, Nums2: []int{3, 4}}, Except: 2.5},

		{Input: Input{Nums1: []int{2}, Nums2: []int{}}, Except: 2.0},
		{Input: Input{Nums1: []int{2, 3}, Nums2: []int{}}, Except: 2.5},

		{Input: Input{Nums1: []int{}, Nums2: []int{2}}, Except: 2.0},
		{Input: Input{Nums1: []int{}, Nums2: []int{2, 3}}, Except: 2.5},

		{Input: Input{Nums1: []int{1, 2, 3, 4, 5, 6}, Nums2: []int{7}}, Except: 4.0},
		{Input: Input{Nums1: []int{1, 2, 3, 4, 5, 6, 7}, Nums2: []int{7}}, Except: 4.5},

		{Input: Input{Nums1: []int{1, 2, 3, 4}, Nums2: []int{5, 6, 7, 8}}, Except: 4.5},
		{Input: Input{Nums1: []int{5, 6, 7, 8}, Nums2: []int{1, 2, 3, 4}}, Except: 4.5},
	}
}

type Func func(nums1 []int, nums2 []int) float64

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) (out float64) {
	return func(in Input) (out float64) {
		return f(in.Nums1, in.Nums2)
	}
}

func AddCases(c func() []tests.Case[Input, float64]) {
	_cases := cases()
	cases = func() []tests.Case[Input, float64] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

func Bench(b *testing.B) {
	tests.Bench(b, cases, funcs...)
}
