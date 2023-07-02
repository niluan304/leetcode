package string_to_integer_atoi

import (
	"math"
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, int] {
	return []tests.Case[string, int]{
		{Input: "9223372036854775808", Except: math.MaxInt32},
		{Input: "-9223372036854775808", Except: math.MinInt32},
		{Input: "4193 with words", Except: 4193},
		{Input: "42", Except: 42},
		{Input: "   -42", Except: -42},
		{Input: "", Except: 0},
		{Input: "-+12", Except: 0},
		{Input: "2147483648", Except: math.MaxInt32},
		{Input: "-2147483649", Except: math.MinInt32},
		{Input: "-  9999999 sdfsdf", Except: 0},
		{Input: "-99999999999999999999999 sdfsdf", Except: math.MinInt32},
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
