package set_matrix_zeroes

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[][]int, [][]int] {
	return []tests.Case[[][]int, [][]int]{
		{
			Input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			Except: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			Input: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			Except: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			Input: [][]int{
				{0, 1, 2, 0},
				{3, 0, 5, 2},
				{1, 3, 1, 5},
			},
			Except: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
			},
		},
		{
			Input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{4, 3, 2, 1},
			},
			Except: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{4, 3, 2, 1},
			},
		},
	}
}

type Func func(matrix [][]int)

func adaptor(f Func) func(in [][]int) (out [][]int) {
	return func(in [][]int) (out [][]int) {
		if !checkResult {
			f(in)
			return in
		}

		for _, v := range in {
			out = append(out, append([]int{}, v...))
		}
		f(out)
		return out
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[[][]int, [][]int]) {
	_cases := cases()
	cases = func() []tests.Case[[][]int, [][]int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
