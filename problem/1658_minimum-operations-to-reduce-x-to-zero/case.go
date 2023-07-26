package minimum_operations_to_reduce_x_to_zero

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums []int
	x    int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				nums: []int{1, 1, 4, 2, 3},
				x:    5,
			},
			Except: 2,
		},
		{
			Input: Input{
				nums: []int{5, 6, 7, 8, 9},
				x:    4,
			},
			Except: -1,
		},
		{
			Input: Input{
				nums: []int{3, 2, 20, 1, 1, 3},
				x:    10,
			},
			Except: 5,
		},
	}
}

type Func func(nums []int, x int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.nums,
			in.x,
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
