package longest_palindromic_substring

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, string] {
	return []tests.Case[string, string]{
		{Input: "babad", Except: "bab"},
		{Input: "babad", Except: "aba"},
		{Input: "cbbd", Except: "bb"},
		{Input: "abba", Except: "abba"},
		{Input: "ababa", Except: "ababa"},
		{Input: "ababazxc", Except: "ababa"},
	}
}

var funcs = tests.NewFunc[string, string]()

func AddCases(c func() []tests.Case[string, string]) {
	_cases := cases()
	cases = func() []tests.Case[string, string] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func(string) string) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[string, string]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[string, string]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
