package word_break

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	S        string
	WordDict []string
}

var cases = func() []tests.Case[Input, bool] {
	return []tests.Case[Input, bool]{
		{Input: Input{S: "leetcode", WordDict: []string{"leet", "code"}}, Except: true},
		{Input: Input{S: "applepenapple", WordDict: []string{"apple", "pen"}}, Except: true},
		{Input: Input{S: "catsandog", WordDict: []string{"cats", "dog", "sand", "and", "cat"}}, Except: false},
	}
}

type Func func(s string, wordDict []string) bool

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) (out bool) {
	return func(in Input) (out bool) {
		return f(in.S, in.WordDict)
	}
}

func AddCases(c func() []tests.Case[Input, bool]) {
	_cases := cases()
	cases = func() []tests.Case[Input, bool] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, bool]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[Input, bool]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
