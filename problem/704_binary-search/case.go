package binary_search

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
		{Input: Input{Nums: []int{-1, 0, 3, 5, 9, 12}, Target: 9}, Except: 4},
		{Input: Input{Nums: []int{-1, 0, 3, 5, 9, 12}, Target: 2}, Except: -1},
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

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
