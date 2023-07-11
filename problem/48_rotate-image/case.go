package rotate_image

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[][]int, [][]int] {
	return []tests.Case[[][]int, [][]int]{
		{
			Input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			Except: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			Input: [][]int{
				{05, 01, 9, 11},
				{02, 04, 8, 10},
				{13, 03, 06, 07},
				{15, 14, 12, 16},
			},
			Except: [][]int{
				{15, 13, 02, 05},
				{14, 03, 04, 01},
				{12, 06, 8, 9},
				{16, 07, 10, 11},
			},
		},
		{
			Input: [][]int{
				{01, 02, 03, 04},
				{11, 12, 13, 14},
				{21, 22, 23, 24},
				{31, 32, 33, 34},
			},
			Except: [][]int{
				{31, 21, 11, 01},
				{32, 22, 12, 02},
				{33, 23, 13, 03},
				{34, 24, 14, 04},
			},
		},
	}
}

func adaptor(f func([][]int)) func(in [][]int) (out [][]int) {
	return func(in [][]int) (out [][]int) {
		// 会影响 bench 基准测试
		out = make([][]int, len(in))
		for i, v := range in {
			out[i] = make([]int, len(v))
			copy(out[i], v)
		}

		f(out)
		return out
	}
}

func Equal(x, y [][]int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if len(v) != len(y[i]) {
			return false
		}
		for j, vv := range v {
			if vv != y[i][j] {
				return false
			}
		}
	}
	return true
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[[][]int, [][]int]) {
	_cases := cases()
	cases = func() []tests.Case[[][]int, [][]int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func([][]int)) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[[][]int, [][]int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[[][]int, [][]int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}
