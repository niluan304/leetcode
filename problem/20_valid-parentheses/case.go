package valid_parentheses

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, bool] {
	return []tests.Case[string, bool]{
		{Input: "()", Except: true},
		{Input: "()[]{}", Except: true},
		{Input: "([]{})", Except: true},
		{Input: "([{}])", Except: true},
		{Input: "()[{}]()", Except: true},
		{Input: "(]", Except: false},
		{Input: "{}{{{}}}({})(){{{}}}", Except: true},
		{Input: "{[]}()", Except: true},
		{Input: "[({(())}[()])]", Except: true},
		{Input: "[(])", Except: false},
	}
}

var funcs = tests.NewFunc[string, bool]()

func AddCases(c func() []tests.Case[string, bool]) {
	_cases := cases()
	cases = func() []tests.Case[string, bool] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func(string) bool) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[string, bool]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[string, bool]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
