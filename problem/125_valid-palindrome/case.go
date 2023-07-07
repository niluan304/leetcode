package valid_palindrome

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	s string
}

type Output bool

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{Input: Input{s: "A man, a plan, a canal: Panama"}, Except: true},
		{Input: Input{s: "race a car"}, Except: false},
		{Input: Input{s: " "}, Except: true},
	}
}

type Func func(s string) bool

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.s,
		))
	}
}

func AddCases(c func() []tests.Case[Input, Output]) {
	_cases := cases()
	cases = func() []tests.Case[Input, Output] {
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
