package simplify_path

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, string] {
	return []tests.Case[string, string]{
		{Input: "/home/", Except: "/home"},
		{Input: "/../", Except: "/"},
		{Input: "/home//foo/", Except: "/home/foo"},
		{Input: "/a/./b/../../c/", Except: "/c"},
	}
}

type Func func(path string) string

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
	tests.Unit(t, tests.Test[string, string]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[string, string]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
