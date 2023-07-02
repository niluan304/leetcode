package length_of_last_word

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, int] {
	return []tests.Case[string, int]{
		{Input: "Hello World", Except: 5},
		{Input: "   fly me   to   the moon  ", Except: 4},
		{Input: "luffy is still joyboy", Except: 6},
	}
}

type Func func(string) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in string) (out int) {
	return func(in string) (out int) {
		return f(in)
	}
}

func AddCases(c func() []tests.Case[string, int]) {
	_cases := cases()
	cases = func() []tests.Case[string, int] {
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
