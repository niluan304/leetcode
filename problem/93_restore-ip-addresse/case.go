package restore_ip_addresse

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, []string] {
	return []tests.Case[string, []string]{
		{Input: "25525511135", Except: []string{"255.255.11.135", "255.255.111.35"}},
		{Input: "0000", Except: []string{"0.0.0.0"}},
		{Input: "101023", Except: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
	}
}

type Func func(s string) []string

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in string) (out []string) {
	return func(in string) (out []string) {
		return f(in)
	}
}

func AddCases(c func() []tests.Case[string, []string]) {
	_cases := cases()
	cases = func() []tests.Case[string, []string] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[string, []string]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[string, []string]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
