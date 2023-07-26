package new_21_game

import (
	"math"
	"testing"

	"leetcode/tests"
)

type Input struct {
	n      int
	k      int
	maxPts int
}

type Output float64

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input:  Input{n: 10, k: 1, maxPts: 10},
			Except: 1.0,
		},
		{
			Input:  Input{n: 6, k: 1, maxPts: 10},
			Except: 0.6,
		},
		{
			Input:  Input{n: 21, k: 17, maxPts: 10},
			Except: 0.73278,
		},
	}
}

type Func func(n int, k int, maxPts int) float64

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.n,
			in.k,
			in.maxPts,
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
	// 与实际答案误差不超过 10-5 的答案将被视为正确答案。

	const n = 100000

	return math.Round(float64(x*n)) == math.Round(float64(y*n))
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
