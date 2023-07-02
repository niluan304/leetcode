package palindrome_partitioning

import (
	"sort"
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, [][]string] {
	return []tests.Case[string, [][]string]{
		{Input: "aab", Except: [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{Input: "a", Except: [][]string{{"a"}}},
	}
}

type Func func(string) [][]string

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in string) (out [][]string) {
	return func(in string) (out [][]string) {
		out = f(in)
		if checkResult {
			sort.Slice(out, func(i, j int) bool {
				var idx = len(out[i])
				if len(out[i]) > len(out[j]) {
					idx = len(out[i])
				}

				for k := 0; k < idx; k++ {
					if out[i][k] != out[j][k] {
						return out[i][k] < out[j][k]
					}
				}
				return true
			})
		}
		return
	}
}

func AddCases(c func() []tests.Case[string, [][]string]) {
	_cases := cases()
	cases = func() []tests.Case[string, [][]string] {
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
