package pascals_triangle

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[int, [][]int] {
	return []tests.Case[int, [][]int]{
		{Input: 1, Except: [][]int{{1}}},
		{Input: 2, Except: [][]int{{1}, {1, 1}}},
		{Input: 3, Except: [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{Input: 4, Except: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
		{Input: 5, Except: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{Input: 6, Except: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}}},
		{Input: 7, Except: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}}},
		{Input: 8, Except: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}, {1, 7, 21, 35, 35, 21, 7, 1}}},
	}
}

var funcs = tests.NewFunc[int, [][]int]()

func AddCases(c func() []tests.Case[int, [][]int]) {
	_cases := cases()
	cases = func() []tests.Case[int, [][]int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func(int) [][]int) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

func Bench(b *testing.B) {
	tests.Bench(b, cases, funcs...)
}
