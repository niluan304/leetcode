package palindrome_number

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[int, bool] {
	return []tests.Case[int, bool]{
		{Input: 101, Except: true},
		{Input: 0, Except: true},
		{Input: 10, Except: false},
		{Input: 7, Except: true},
		{Input: 77, Except: true},
		{Input: 111, Except: true},
		{Input: 102, Except: false},
		{Input: 25102, Except: false},
		{Input: 2552, Except: true},
		{Input: 25152, Except: true},
	}
}

var funcs = tests.NewFunc[int, bool]()

func AddCases(c func() []tests.Case[int, bool]) {
	_cases := cases()
	cases = func() []tests.Case[int, bool] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func(int) bool) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

func Bench(b *testing.B) {
	tests.Bench(b, cases, funcs...)
}
