package max_dot_product_of_two_subsequences

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums1 []int
	nums2 []int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				nums1: []int{2, 1, -2, 5},
				nums2: []int{3, 0, -6},
			},
			Except: 18,
		},
		{
			Input: Input{
				nums1: []int{3, -2},
				nums2: []int{2, -6, 7},
			},
			Except: 21,
		},
		{
			Input: Input{
				nums1: []int{-1, -1},
				nums2: []int{1, 1},
			},
			Except: -1,
		},
	}
}

type Func func(nums1 []int, nums2 []int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.nums1,
			in.nums2,
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

func Equal(x, y Output) bool {

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
