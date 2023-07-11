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
		return
	}
}

func AddCases(c func() []tests.Case[string, [][]string]) {
	_cases := cases()
	cases = func() []tests.Case[string, [][]string] {
		return append(_cases, c()...)
	}
}

func Equal(x, y [][]string) bool {
	sortFunc := func(output [][]string) {
		sort.Slice(output, func(i, j int) bool {
			a, b := output[i], output[j]
			if len(a) != len(b) {
				return len(a) < len(b)
			}

			for k := 0; k < len(a); k++ {
				if a[k] != b[k] {
					return a[k] < b[k]
				}
			}
			return i < j
		})
	}

	sortFunc(x)
	sortFunc(y)

	return tests.Equal(x, y)
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[string, [][]string]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[string, [][]string]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}
