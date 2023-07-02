package merge_sorted_array

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	Nums1, Nums2 []int
	M, N         int
}

var cases = func() []tests.Case[Input, []int] {
	return []tests.Case[Input, []int]{
		{
			Input:  Input{Nums1: []int{1, 2, 3, 0, 0, 0}, Nums2: []int{2, 5, 6}, M: 3, N: 3},
			Except: []int{1, 2, 2, 3, 5, 6},
		},
		{
			Input:  Input{Nums1: []int{1}, Nums2: []int{}, M: 1, N: 0},
			Except: []int{1},
		},
		{
			Input:  Input{Nums1: []int{0}, Nums2: []int{2}, M: 0, N: 1},
			Except: []int{2},
		},
	}
}

type Func func(nums1 []int, m int, nums2 []int, n int)

func adaptor(f Func) func(in Input) (out []int) {
	return func(in Input) (out []int) {
		nums1, nums2 := in.Nums1, in.Nums2
		f(nums1, in.M, nums2, in.N)
		return nums1
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[Input, []int]) {
	_cases := cases()
	cases = func() []tests.Case[Input, []int] {
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
