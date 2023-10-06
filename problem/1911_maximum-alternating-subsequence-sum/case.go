package maximum_alternating_subsequence_sum

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums []int
}

type Output int64

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				nums: []int{4, 2, 5, 3},
			},
			Except: 7,
		},
		{
			Input: Input{
				nums: []int{5, 6, 7, 8},
			},
			Except: 8,
		},
		{
			Input: Input{
				nums: []int{6, 2, 1, 2, 4, 5},
			},
			Except: 10,
		},
	}
}

type Func func(nums []int) int64

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
