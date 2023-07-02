package three_sum_closest

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	Nums   []int
	Target int
}

var cases = func() []tests.Case[Input, int] {
	return []tests.Case[Input, int]{
		{
			Input:  Input{Nums: []int{-1, 2, 1, -4}, Target: 1},
			Except: 2,
		},
		{
			Input:  Input{Nums: []int{0, 0, 0}, Target: 1},
			Except: 0,
		},
		{
			Input:  Input{Nums: []int{0, 3, 97, 102, 200}, Target: 300},
			Except: 300,
		},
	}
}

type Func func(nums []int, target int) int

func adaptor(f Func) func(in Input) (out int) {
	return func(in Input) (out int) {
		return f(in.Nums, in.Target)
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
	tests.Unit(t, cases, funcs...)
}

func Bench(b *testing.B) {
	tests.Bench(b, cases, funcs...)
}
