package corporate_flight_bookings

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	bookings [][]int
	n        int
}

type Output []int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				bookings: [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}},
				n:        5,
			},
			Except: []int{10, 55, 45, 25, 25},
		},
		{
			Input: Input{
				bookings: [][]int{{1, 2, 10}, {2, 2, 15}},
				n:        2,
			},
			Except: []int{10, 25},
		},
	}
}

type Func func(bookings [][]int, n int) []int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.bookings,
			in.n,
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
