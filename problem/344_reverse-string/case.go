package reverse_string

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, string] {
	return []tests.Case[string, string]{
		{Input: "hello", Except: "olleh"},
		{Input: "Hannah", Except: "hannaH"},
	}
}

type Func func(s []byte)

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in string) (out string) {
	return func(in string) (out string) {
		input := []byte(in)
		f(input)

		if !checkResult {
			return
		}

		return string(input)
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
