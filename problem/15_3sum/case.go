package three_sum

import (
	"sort"
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

func Sort(x [][]int) (x1 [][]int) {
	sort.Slice(x, func(i, j int) bool {
		if len(x[i]) == 0 {
			return true
		}
		if len(x[j]) == 0 {
			return false
		}
		if x[i][0] != x[j][0] {
			return x[i][0] < x[j][0]
		}
		if x[i][1] != x[j][1] {
			return x[i][1] < x[j][1]
		}
		return x[i][2] < x[j][2]
	})

	for i, v := range x {
		if len(v) == 0 {
			continue
		}
		x1 = append(x1, x[i])
	}

	return x1
}

func Equal(x, y [][]int) bool {
	x1, y1 := Sort(x), Sort(y)

	if len(x1) != len(y1) {
		return false
	}

	for i, v := range x1 {
		if len(v) != len(y1[i]) {
			return false
		}
		for j, v1 := range v {
			if v1 != y1[i][j] {
				return false
			}
		}
	}

	return true
}

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

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[[]int, [][]int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}
