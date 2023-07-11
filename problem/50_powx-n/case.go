package powx_n

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	X float64
	N int
}

type Output float64

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{Input: Input{X: 0.00001, N: 2147483647}, Except: 0},
		{Input: Input{X: 2.0, N: 10}, Except: 1024.00000},
		{Input: Input{X: 2.1, N: 3}, Except: 9.26100},
		{Input: Input{X: 2.0, N: -2}, Except: 0.25000},
	}
}

type Func func(x float64, n int) float64

func adaptor(f Func) func(in Input) (out Output) {
	return func(in Input) (out Output) {
		return Output(f(in.X, in.N))
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

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
