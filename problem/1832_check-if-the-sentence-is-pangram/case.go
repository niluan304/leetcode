package check_if_the_sentence_is_pangram

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, bool] {
	return []tests.Case[string, bool]{
		{Input: "thequickbrownfoxjumpsoverthelazydog", Except: true},
		{Input: "qwertyuiopasdfghjklzxcvbnnnnn", Except: false},
		{Input: "leetcode", Except: false},
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
	funcs = append(funcs, tests.NewFunc[string, bool](f...)...)
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
