package count_ways_to_build_good_strings

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	low  int
	high int
	zero int
	one  int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				low:  3,
				high: 3,
				zero: 1,
				one:  1,
			},
			Except: 8,
		},
		{
			Input: Input{
				low:  2,
				high: 3,
				zero: 1,
				one:  2,
			},
			Except: 5,
		},
	}
}

type Func func(low int, high int, zero int, one int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.low,
			in.high,
			in.zero,
			in.one,
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
