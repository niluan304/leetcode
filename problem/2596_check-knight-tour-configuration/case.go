package check_knight_tour_configuration

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	grid [][]int
}

type Output bool

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				grid: [][]int{
					{0, 11, 16, 5, 20},
					{17, 4, 19, 10, 15},
					{12, 1, 8, 21, 6},
					{3, 18, 23, 14, 9},
					{24, 13, 2, 7, 22},
				},
			},
			Except: true,
		},
		{
			Input: Input{
				grid: [][]int{
					{0, 3, 6},
					{5, 8, 1},
					{2, 7, 4},
				},
			},
			Except: false,
		},
	}
}

type Func func(grid [][]int) bool

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.grid,
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

func Equal(x, y Output) bool {

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
