package three_sum

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, [][]int] {
	return []tests.Case[[]int, [][]int]{
		{
			Input:  []int{-1, 0, 1, 2, -1, -4},
			Except: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			Input:  []int{0, 0, 0},
			Except: [][]int{{0, 0, 0}},
		},
		{
			Input:  []int{0, 1, 1},
			Except: [][]int{},
		},
		{
			Input:  []int{0, 0, 0, 0, 0},
			Except: [][]int{{0, 0, 0}},
		},
	}
}

var funcs = tests.NewFunc[[]int, [][]int]()

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
		IsEqual:  nil,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[[]int, [][]int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
