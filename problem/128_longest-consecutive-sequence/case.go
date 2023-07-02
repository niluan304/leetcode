package longest_consecutive_sequence

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, int] {
	return []tests.Case[[]int, int]{
		{Input: []int{}, Except: 0},
		{Input: []int{3}, Except: 1},
		{Input: []int{100, 4, 200, 1, 3, 2}, Except: 4},
		{Input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, Except: 9},
		{Input: []int{1, 2, 0, 1}, Except: 3},
	}
}

var funcs = tests.NewFunc[[]int, int]()

func AddCases(c func() []tests.Case[[]int, int]) {
	_cases := cases()
	cases = func() []tests.Case[[]int, int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func([]int) int) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

func Bench(b *testing.B) {
	tests.Bench(b, cases, funcs...)
}
