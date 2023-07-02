package merge_intervals

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[][]int, [][]int] {
	return []tests.Case[[][]int, [][]int]{
		{Input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, Except: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{Input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {7, 7}}, Except: [][]int{{1, 6}, {7, 7}, {8, 10}, {15, 18}}},
		{Input: [][]int{{1, 4}, {4, 5}}, Except: [][]int{{1, 5}}},
		{Input: [][]int{{1, 4}, {4, 4}, {4, 5}}, Except: [][]int{{1, 5}}},
		{Input: [][]int{{1, 4}, {4, 5}, {7, 7}}, Except: [][]int{{1, 5}, {7, 7}}},
		{Input: [][]int{{1, 4}, {4, 8}, {7, 7}}, Except: [][]int{{1, 8}}},
		{Input: [][]int{{1, 9}, {2, 5}, {19, 20}, {10, 11}, {12, 20}, {0, 3}, {0, 1}, {0, 2}}, Except: [][]int{{0, 9}, {10, 11}, {12, 20}}},
	}
}

var funcs = tests.NewFunc[[][]int, [][]int]()

func AddCases(c func() []tests.Case[[][]int, [][]int]) {
	_cases := cases()
	cases = func() []tests.Case[[][]int, [][]int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func([][]int) [][]int) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
