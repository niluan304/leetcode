package decode_string

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, string] {
	return []tests.Case[string, string]{
		{Input: "3[a]2[bc]", Except: "aaabcbc"},
		{Input: "3[a2[c]]", Except: "accaccacc"},
		{Input: "2[abc]3[cd]ef", Except: "abcabccdcdcdef"},
		{Input: "abc3[cd]xyz", Except: "abccdcdcdxyz"},
	}
}

type Func func(string) string

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in string) (out string) {
	return func(in string) (out string) {
		return f(in)
	}
}

func AddCases(c func() []tests.Case[string, string]) {
	_cases := cases()
	cases = func() []tests.Case[string, string] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
