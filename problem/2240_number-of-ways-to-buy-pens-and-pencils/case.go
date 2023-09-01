package number_of_ways_to_buy_pens_and_pencils

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	total int
	cost1 int
	cost2 int
}

type Output int64

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				total: 20,
				cost1: 10,
				cost2: 5,
			},
			Except: 9,
		},
		{
			Input: Input{
				total: 5,
				cost1: 10,
				cost2: 10,
			},
			Except: 1,
		},
	}
}

type Func func(total int, cost1 int, cost2 int) int64

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.total,
			in.cost1,
			in.cost2,
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
