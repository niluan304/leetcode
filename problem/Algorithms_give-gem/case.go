package give_gem

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	gem        []int
	operations [][]int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				gem:        []int{3, 1, 2},
				operations: [][]int{{0, 2}, {2, 1}, {2, 0}},
			},
			Except: 2,
		},
		{
			Input: Input{
				gem:        []int{100, 0, 50, 100},
				operations: [][]int{{0, 2}, {0, 1}, {3, 0}, {3, 0}},
			},
			Except: 75,
		},
		{
			Input: Input{
				gem:        []int{0, 0, 0, 0},
				operations: [][]int{{1, 2}, {3, 1}, {1, 2}},
			},
			Except: 0,
		},
	}
}

type Func func(gem []int, operations [][]int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.gem,
			in.operations,
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
