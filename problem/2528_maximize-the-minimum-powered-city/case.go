package maximize_the_minimum_powered_city

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	stations []int
	r        int
	k        int
}

type Output int64

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				stations: []int{1, 2, 3, 5, 0},
				r:        1,
				k:        2,
			},
			Except: 5,
		},
		{
			Input: Input{
				stations: []int{4, 4, 4, 4},
				r:        0,
				k:        3,
			},
			Except: 4,
		},
	}
}

type Func func(stations []int, r int, k int) int64

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.stations,
			in.r,
			in.k,
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
