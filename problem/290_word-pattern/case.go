package word_pattern

import (
	"testing"

	"leetcode/tests"
)

type Input struct{ Pattern, S string }

var cases = func() []tests.Case[Input, bool] {
	return []tests.Case[Input, bool]{
		{Input: Input{Pattern: "abba", S: "dog cat cat dog"}, Except: true},
		{Input: Input{Pattern: "abba", S: "dog cat cat fish"}, Except: false},
		{Input: Input{Pattern: "aaaa", S: "dog cat cat dog"}, Except: false},
		{Input: Input{Pattern: "abba", S: "dog dog dog dog"}, Except: false},
		{Input: Input{Pattern: "abab", S: "dog cat cat dog"}, Except: false},
	}
}

type Func func(pattern string, s string) bool

func adaptor(f Func) func(in Input) (out bool) {
	return func(in Input) (out bool) {
		return f(in.Pattern, in.S)
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[Input, bool]) {
	_cases := cases()
	cases = func() []tests.Case[Input, bool] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

func Bench(b *testing.B) {
	tests.Bench(b, cases, funcs...)
}
