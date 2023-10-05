package best_time_to_buy_and_sell_stock_with_transaction_fee

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	prices []int
	fee    int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				prices: []int{1, 3, 2, 8, 4, 9},
				fee:    2,
			},
			Except: 8,
		},
		{
			Input: Input{
				prices: []int{1, 3, 7, 5, 10, 3},
				fee:    3,
			},
			Except: 6,
		},
	}
}

type Func func(prices []int, fee int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.prices,
			in.fee,
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
