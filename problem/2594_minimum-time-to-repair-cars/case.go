package minimum_time_to_repair_cars

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	ranks []int
	cars  int
}

type Output int64

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				ranks: []int{4, 2, 3, 1},
				cars:  10,
			},
			Except: 16,
		},
		{
			Input: Input{
				ranks: []int{5, 1, 8},
				cars:  6,
			},
			Except: 16,
		},
	}
}

type Func func(ranks []int, cars int) int64

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.ranks,
			in.cars,
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
