package count_primes

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[int, int] {
	return []tests.Case[int, int]{
		{Input: 10, Except: 4},
		{Input: 0, Except: 0},
		{Input: 1, Except: 0},
		{Input: 10000, Except: 1229},
		{Input: 999983, Except: 78497},
		{Input: 1500000, Except: 114155},
	}
}

var funcs = tests.NewFunc[int, int]()

func AddCases(c func() []tests.Case[int, int]) {
	_cases := cases()
	cases = func() []tests.Case[int, int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func(int) int) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
