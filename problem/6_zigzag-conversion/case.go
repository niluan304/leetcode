package zigzag_conversion

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	S       string
	NumRows int
}

var cases = func() []tests.Case[Input, string] {
	return []tests.Case[Input, string]{
		{Input: Input{S: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", NumRows: 2}, Except: "ACEGIKMOQSUWYBDFHJLNPRTVXZ"},
		{Input: Input{S: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", NumRows: 3}, Except: "AEIMQUYBDFHJLNPRTVXZCGKOSW"},
		{Input: Input{S: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", NumRows: 4}, Except: "AGMSYBFHLNRTXZCEIKOQUWDJPV"},
		{Input: Input{S: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", NumRows: 5}, Except: "AIQYBHJPRXZCGKOSWDFLNTVEMU"},
		{Input: Input{S: "PAYPALISHIRING", NumRows: 3}, Except: "PAHNAPLSIIGYIR"},
		{Input: Input{S: "PAYPALISHIRING", NumRows: 4}, Except: "PINALSIGYAHRPI"},
		{Input: Input{S: "A", NumRows: 1}, Except: "A"},
	}
}

type Func func(string, int) string

func adaptor(f Func) func(in Input) (out string) {
	return func(in Input) (out string) {
		return f(in.S, in.NumRows)
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

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
	// checkResult = false
	tests.Bench(b, cases, funcs...)
}
