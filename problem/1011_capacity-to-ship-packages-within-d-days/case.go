package capacity_to_ship_packages_within_d_days

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	weights []int
	days    int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{Input: Input{weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, days: 5}, Except: 15},
		{Input: Input{weights: []int{3, 2, 2, 4, 1, 4}, days: 3}, Except: 6},
		{Input: Input{weights: []int{1, 2, 3, 1, 1}, days: 4}, Except: 3},
	}
}

type Func func(weights []int, days int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.weights,
			in.days,
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

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
