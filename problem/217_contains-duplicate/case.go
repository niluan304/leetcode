package contains_duplicate

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, bool] {
	return []tests.Case[[]int, bool]{
		{Input: []int{1, 2, 3, 1}, Except: true},
		{Input: []int{1, 2, 3, 4}, Except: false},
		{Input: []int{1, 2}, Except: false},
		{Input: []int{1}, Except: false},
		{Input: []int{}, Except: false},
		{Input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, Except: true},
	}
}

var funcs = tests.NewFunc[[]int, bool]()

func AddCases(c func() []tests.Case[[]int, bool]) {
	_cases := cases()
	cases = func() []tests.Case[[]int, bool] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func([]int) bool) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

func Bench(b *testing.B) {
	tests.Bench(b, cases, funcs...)
}
