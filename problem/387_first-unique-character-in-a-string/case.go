package first_unique_character_in_a_string

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, int] {
	return []tests.Case[string, int]{
		{Input: "leetcode", Except: 0},
		{Input: "loveleetcode", Except: 2},
		{Input: "aabb", Except: -1},
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
	tests.Unit(t, cases, funcs...)
}

func Bench(b *testing.B) {
	tests.Bench(b, cases, funcs...)
}
