package coin_change_ii

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	amount int
	coins  []int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				amount: 5,
				coins:  []int{1, 2, 5},
			},
			Except: 4,
		},
		{
			Input: Input{
				amount: 3,
				coins:  []int{2},
			},
			Except: 0,
		},
		{
			Input: Input{
				amount: 10,
				coins:  []int{10},
			},
			Except: 1,
		},
	}
}

type Func func(amount int, coins []int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.amount,
			in.coins,
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
