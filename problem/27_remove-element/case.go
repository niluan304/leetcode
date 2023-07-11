package remove_element

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	Nums []int
	Val  int
}

var cases = func() []tests.Case[Input, int] {
	return []tests.Case[Input, int]{
		{
			Input:  Input{Nums: []int{3, 2, 2, 3}, Val: 3},
			Except: 2,
		},
		{
			Input:  Input{Nums: []int{3, 3, 3, 3}, Val: 3},
			Except: 0,
		},
		{
			Input:  Input{Nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, Val: 2},
			Except: 5,
		},
	}
}

type Func func(nums []int, val int) int

func adaptor(f Func) func(in Input) (out int) {
	return func(in Input) (out int) {
		return f(in.Nums, in.Val)
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
