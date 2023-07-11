package minimum_limit_of_balls_in_a_bag

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	Nums          []int
	MaxOperations int
}

var cases = func() []tests.Case[Input, int] {
	return []tests.Case[Input, int]{
		{
			Input: Input{
				Nums:          []int{9},
				MaxOperations: 2,
			},
			Except: 3,
		},
		{
			Input: Input{
				Nums:          []int{2, 4, 8, 2},
				MaxOperations: 4,
			},
			Except: 2,
		},
		{
			Input: Input{
				Nums:          []int{7, 17},
				MaxOperations: 2,
			},
			Except: 7,
		},
		{
			Input: Input{
				Nums:          []int{1},
				MaxOperations: 1,
			},
		},
	}
}

type Func func(nums []int, maxOperations int) int

func adaptor(f Func) func(in Input) (out int) {
	return func(in Input) (out int) {
		return f(in.Nums, in.MaxOperations)
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[Input, int]) {
	_cases := cases()
	cases = func() []tests.Case[Input, int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[Input, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
