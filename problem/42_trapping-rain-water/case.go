package trapping_rain_water

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, int] {
	return []tests.Case[[]int, int]{
		{Input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, Except: 6},
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
