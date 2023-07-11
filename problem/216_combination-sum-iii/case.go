package combination_sum_iii

import (
	"sort"
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

func Equal(a, b Output) bool {
	if len(a) != len(b) {
		return false
	}

	sortFunc := func(list Output) {
		sort.Slice(list, func(i, j int) bool {
			for k := 0; k < len(list[i]); k++ {
				if list[i][k] != list[j][k] {
					return list[i][k] < list[j][k]
				}
			}

			return true
		})
	}

	sortFunc(a)
	sortFunc(b)

	for i, v := range a {
		if len(v) != len(b[i]) {
			return false
		}

		for j, vv := range v {
			if vv != b[i][j] {
				return false
			}
		}
	}

	return true
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})

}
