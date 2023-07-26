package contains_duplicate_ii

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums []int
	k    int
}

type Output bool

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			Except: true,
		},
		{
			Input: Input{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			Except: true,
		},
		{
			Input: Input{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			Except: false,
		},
	}
}

type Func func(nums []int, k int) bool

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.nums,
			in.k,
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
