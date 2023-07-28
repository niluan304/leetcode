package longest_subarray_of_1s_after_deleting_one_element

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums []int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input:  Input{nums: []int{1, 1, 0, 1}},
			Except: 3,
		},
		{
			Input:  Input{nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}},
			Except: 5,
		},
		{
			Input:  Input{nums: []int{1, 1, 1}},
			Except: 2,
		},
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
