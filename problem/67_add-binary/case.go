package add_binary

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	A, B string
}

var cases = func() []tests.Case[Input, string] {
	return []tests.Case[Input, string]{
		{Input: Input{A: "11", B: "1"}, Except: "100"},
		{Input: Input{A: "1010", B: "1011"}, Except: "10101"},
	}
}

type Func func(a string, b string) string

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) (out string) {
	return func(in Input) (out string) {
		return f(in.A, in.B)
	}
}

func AddCases(c func() []tests.Case[Input, string]) {
	_cases := cases()
	cases = func() []tests.Case[Input, string] {
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
