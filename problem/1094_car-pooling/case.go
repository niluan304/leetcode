package car_pooling

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	trips    [][]int
	capacity int
}

type Output bool

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
				capacity: 4,
			},
			Except: false,
		},
		{
			Input: Input{
				trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
				capacity: 5,
			},
			Except: true,
		},
	}
}

type Func func(trips [][]int, capacity int) bool

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.trips,
			in.capacity,
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
