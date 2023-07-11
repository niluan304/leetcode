package sqrtx

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[int, int] {
	return []tests.Case[int, int]{
		{Input: 4, Except: 2},
		{Input: 8, Except: 2},
	}
}

type Func func(x int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in int) (out int) {
	return f
}

func AddCases(c func() []tests.Case[int, int]) {
	_cases := cases()
	cases = func() []tests.Case[int, int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[int, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[int, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
