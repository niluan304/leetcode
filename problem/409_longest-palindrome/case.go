package longest_palindrome

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, int] {
	return []tests.Case[string, int]{
		{Input: "abccccdd", Except: 7},
		{Input: "a", Except: 1},
		{Input: "aA", Except: 1},
		{Input: "aaAA", Except: 4},
		{Input: "aabAA", Except: 5},
		{Input: "aaaaaccc", Except: 7},
		{Input: "bananas", Except: 5},
		{Input: "zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez", Except: 55},
	}
}

var funcs = tests.NewFunc[string, int]()

func AddCases(c func() []tests.Case[string, int]) {
	_cases := cases()
	cases = func() []tests.Case[string, int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func(string) int) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[string, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[string, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
