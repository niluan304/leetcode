package longest_increasing_subsequence

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
			Input: Input{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			Except: 4,
		},
		{
			Input: Input{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			Except: 4,
		},
		{
			Input: Input{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			Except: 1,
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
