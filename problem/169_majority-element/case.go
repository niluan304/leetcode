package majority_element

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, int] {
	return []tests.Case[[]int, int]{
		{Input: []int{3, 2, 3}, Except: 3},
		{Input: []int{3}, Except: 3},
		{Input: []int{2, 2, 1, 1, 1, 2, 2}, Except: 2},
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
