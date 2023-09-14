package queens_that_can_attack_the_king

import (
	"reflect"
	"sort"
	"testing"

	"leetcode/tests"
)

type Input struct {
	queens [][]int
	king   []int
}

type Output [][]int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				queens: [][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}},
				king:   []int{0, 0},
			},
			Except: [][]int{{0, 1}, {1, 0}, {3, 3}},
		},
		{
			Input: Input{
				queens: [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}},
				king:   []int{3, 3},
			},
			Except: [][]int{{2, 2}, {3, 4}, {4, 4}},
		},
		{
			Input: Input{
				queens: [][]int{{5, 6}, {7, 7}, {2, 1}, {0, 7}, {1, 6}, {5, 1}, {3, 7}, {0, 3}, {4, 0}, {1, 2}, {6, 3}, {5, 0}, {0, 4}, {2, 2}, {1, 1}, {6, 4}, {5, 4}, {0, 0}, {2, 6}, {4, 5}, {5, 2}, {1, 4}, {7, 5}, {2, 3}, {0, 5}, {4, 2}, {1, 0}, {2, 7}, {0, 1}, {4, 6}, {6, 1}, {0, 6}, {4, 3}, {1, 7}},
				king:   []int{3, 4},
			},
			Except: [][]int{{2, 3}, {1, 4}, {1, 6}, {3, 7}, {4, 3}, {5, 4}, {4, 5}},
		},
	}
}

type Func func(queens [][]int, king []int) [][]int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.queens,
			in.king,
		))
	}
}

func AddCases(c func() []tests.Case[Input, Output]) {
	_cases := cases()
	cases = func() []tests.Case[Input, Output] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Sort(x Output) {
	sort.Slice(x, func(i, j int) bool {
		xi, xj := x[i], x[j]
		if xi[0] == xj[0] {
			return xi[1] < xj[1]
		}
		return xi[0] < xj[0]
	})
}

func Equal(x, y Output) bool {
	Sort(x)
	Sort(y)
	return reflect.DeepEqual(x, y)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}
