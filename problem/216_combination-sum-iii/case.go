package combination_sum_iii

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	k int
	n int
}

type Output [][]int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{Input: Input{k: 3, n: 7}, Except: [][]int{{1, 2, 4}}},
		{Input: Input{k: 3, n: 9}, Except: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{Input: Input{k: 4, n: 1}, Except: [][]int{}},
	}
}

type Func func(k int, n int) [][]int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.k,
			in.n,
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

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
