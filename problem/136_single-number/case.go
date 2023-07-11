package single_number

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, int] {
	return []tests.Case[[]int, int]{
		{Input: []int{2, 2, 1}, Except: 1},
		{Input: []int{4, 1, 2, 1, 2}, Except: 4},
		{Input: []int{1}, Except: 1},
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
	tests.Unit(t, tests.Test[[]int, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[[]int, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
