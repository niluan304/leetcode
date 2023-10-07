package divide_intervals_into_minimum_number_of_groups

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	intervals [][]int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				intervals: [][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}},
			},
			Except: 3,
		},
		{
			Input: Input{
				intervals: [][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}},
			},
			Except: 1,
		},
	}
}

type Func func(intervals [][]int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.intervals,
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
