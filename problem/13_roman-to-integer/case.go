package roman_to_integer

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[string, int] {
	return []tests.Case[string, int]{
		{Input: "I", Except: 1}, {Input: "II", Except: 2}, {Input: "III", Except: 3}, {Input: "IV", Except: 4}, {Input: "V", Except: 5},
		{Input: "VI", Except: 6}, {Input: "VII", Except: 7}, {Input: "VIII", Except: 8}, {Input: "IX", Except: 9}, {Input: "X", Except: 10},
		{Input: "XI", Except: 11}, {Input: "XII", Except: 12}, {Input: "XIII", Except: 13}, {Input: "XIV", Except: 14}, {Input: "XV", Except: 15},
		{Input: "XVI", Except: 16}, {Input: "XVII", Except: 17}, {Input: "XVIII", Except: 18}, {Input: "XIX", Except: 19}, {Input: "XX", Except: 20},
		{Input: "XXX", Except: 30}, {Input: "XL", Except: 40}, {Input: "XLIX", Except: 49}, {Input: "L", Except: 50}, {Input: "LVIII", Except: 58},
		{Input: "LX", Except: 60}, {Input: "LXX", Except: 70}, {Input: "LXXX", Except: 80}, {Input: "XC", Except: 90}, {Input: "XCIX", Except: 99},
		{Input: "C", Except: 100}, {Input: "CI", Except: 101}, {Input: "CII", Except: 102}, {Input: "CXCIX", Except: 199}, {Input: "CC", Except: 200},
		{Input: "CCC", Except: 300}, {Input: "CD", Except: 400}, {Input: "D", Except: 500}, {Input: "DC", Except: 600}, {Input: "DCC", Except: 700},
		{Input: "DCCC", Except: 800}, {Input: "CM", Except: 900}, {Input: "CMXCIX ", Except: 999}, {Input: "M", Except: 1000}, {Input: "MCD", Except: 1400},
		{Input: "MCDXXXVII", Except: 1437}, {Input: "MD", Except: 1500}, {Input: "MDCCC", Except: 1800}, {Input: "MCM", Except: 1900}, {Input: "MCMXCIV", Except: 1994},
		{Input: "MM", Except: 2000}, {Input: "MMM", Except: 3000}, {Input: "MMMCCCXXXIII", Except: 3333}, {Input: "MMMCMXCIX", Except: 3999},
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

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[string, int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
