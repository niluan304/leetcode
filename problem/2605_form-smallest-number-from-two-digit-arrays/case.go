package form_smallest_number_from_two_digit_arrays

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
				nums1: []int{4, 1, 3},
				nums2: []int{5, 7},
			},
			Except: 15,
		},
		{
			Input: Input{
				nums1: []int{3, 5, 2, 6},
				nums2: []int{3, 1, 7},
			},
			Except: 3,
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
