package non_overlapping_intervals

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[][]int, int] {
	return []tests.Case[[][]int, int]{
		{Input: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, Except: 1},
		{Input: [][]int{{1, 3}, {2, 3}, {3, 4}, {1, 2}}, Except: 1},
		{Input: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}, {0, 4}}, Except: 2},
		{Input: [][]int{{1, 2}, {1, 2}, {1, 2}}, Except: 2},
		{Input: [][]int{{1, 2}, {2, 3}}, Except: 0},
		{Input: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}, Except: 2},
	}
}

type Func func(intervals [][]int) int

func adaptor(f Func) func(in [][]int) (out int) {
	return func(in [][]int) (out int) {
		if !checkResult {
			return f(in)
		}

		var _in [][]int
		for _, v := range in {
			_in = append(_in, append([]int{}, v...))
		}

		return f(_in)
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[[][]int, int]) {
	_cases := cases()
	cases = func() []tests.Case[[][]int, int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[[][]int, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[[][]int, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
