package subsets

import (
	"sort"
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, [][]int] {
	return []tests.Case[[]int, [][]int]{
		{Input: []int{1, 2, 3}, Except: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{Input: []int{0}, Except: [][]int{{}, {0}}},
		{Input: []int{0, 3, 5, 7, 9}, Except: [][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}}},
	}
}

func Sort(output [][]int) {
	sort.Slice(output, func(i, j int) bool {
		x, y := output[i], output[j]
		if len(x) != len(y) {
			return len(x) < len(y)
		}

		for k := range x {
			if x[k] != y[k] {
				return x[k] < y[k]
			}
		}

		return true
	})
}

func Equal(x, y [][]int) bool {
	Sort(x)
	Sort(y)
	return tests.Equal(x, y)
}

func adaptor(f func(nums []int) [][]int) func(in []int) (out [][]int) {
	return func(in []int) (out [][]int) {
		out = f(in)
		return out
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[[]int, [][]int]) {
	_cases := cases()
	cases = func() []tests.Case[[]int, [][]int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func([]int) [][]int) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[[]int, [][]int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[[]int, [][]int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}
