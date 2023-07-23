package minimum_size_subarray_sum

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	target int
	nums   []int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			Except: 2,
		},
		{
			Input: Input{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			Except: 1,
		},
		{
			Input: Input{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			Except: 0,
		},
	}
}

type Func func(target int, nums []int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.target,
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
