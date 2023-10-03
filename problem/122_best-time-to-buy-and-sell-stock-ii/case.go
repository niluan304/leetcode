package best_time_to_buy_and_sell_stock_ii

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	prices []int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			Except: 7,
		},
		{
			Input: Input{
				prices: []int{1, 2, 3, 4, 5},
			},
			Except: 4,
		},
		{
			Input: Input{
				prices: []int{7, 6, 4, 3, 1},
			},
			Except: 0,
		},
	}
}

type Func func(prices []int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.prices,
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
