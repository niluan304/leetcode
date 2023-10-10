package powx_n

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	x float64
	n int
}

type Output float64

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				x: 2.0,
				n: 10,
			},
			Except: 1024.0,
		},
		{
			Input: Input{
				x: 2.1,
				n: 3,
			},
			Except: 9.26100,
		},
		{
			Input: Input{
				x: 2.0,
				n: -2,
			},
			Except: 0.25,
		},
	}
}

type Func func(x float64, n int) float64

func adaptor(f Func) func(in Input) (out Output) {
	return func(in Input) (out Output) {
		return Output(f(in.x, in.n))
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

func Equal(x, y Output) bool {
	if x < y {
		x, y = y, x
	}
	return x-y < 1e-10
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
