package subsets

import (
	"sort"
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, [][]int] {
	return []tests.Case[[]int, [][]int]{
		{Input: []int{1, 2, 3}, Except: Sort([][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}})},
		{Input: []int{0}, Except: Sort([][]int{{}, {0}})},
		{Input: []int{0, 3, 5, 7, 9}, Except: Sort([][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}})},
	}
}

func Sort(result [][]int) [][]int {
	sort.Slice(result, func(i, j int) bool {
		xi, xj := result[i], result[j]
		if len(xi) != len(xj) {
			return len(xi) < len(xj)
		} else {
			for k := range xi {
				if xi[k] != xj[k] {
					return xi[k] < xj[k]
				}
			}
		}
		return true
	})

	return result
}

func adaptor(f func(nums []int) [][]int) func(in []int) (out [][]int) {
	return func(in []int) (out [][]int) {
		out = f(in)
		if checkResult {
			Sort(out)
		}
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
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
