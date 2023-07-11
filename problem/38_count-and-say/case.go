package count_and_say

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[int, string] {
	return []tests.Case[int, string]{
		{Input: 1, Except: "1"},
		{Input: 2, Except: "11"},
		{Input: 3, Except: "21"},
		{Input: 4, Except: "1211"},
		{Input: 5, Except: "111221"},
	}
}

type Func func(n int) string

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in int) (out string) {
	return func(in int) (out string) {
		return f(in)
	}
}

func AddCases(c func() []tests.Case[int, string]) {
	_cases := cases()
	cases = func() []tests.Case[int, string] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[int, string]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[int, string]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
