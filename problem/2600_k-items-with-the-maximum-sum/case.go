package k_items_with_the_maximum_sum

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	numOnes    int
	numZeros   int
	numNegOnes int
	k          int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{Input: Input{numOnes: 3, numZeros: 2, numNegOnes: 0, k: 2}, Except: 2},
		{Input: Input{numOnes: 3, numZeros: 2, numNegOnes: 0, k: 4}, Except: 3},
	}
}

type Func func(numOnes int, numZeros int, numNegOnes int, k int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.numOnes,
			in.numZeros,
			in.numNegOnes,
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
