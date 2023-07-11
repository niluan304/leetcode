package integer_to_roman

import (
	"testing"

	"leetcode/tests"
)

type (
	Input  int
	Output string
)

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{Input: 1, Except: "I"}, {Input: 2, Except: "II"}, {Input: 3, Except: "III"}, {Input: 4, Except: "IV"}, {Input: 5, Except: "V"}, {Input: 6, Except: "VI"}, {Input: 7, Except: "VII"}, {Input: 8, Except: "VIII"}, {Input: 9, Except: "IX"}, {Input: 10, Except: "X"},
		{Input: 11, Except: "XI"}, {Input: 12, Except: "XII"}, {Input: 13, Except: "XIII"}, {Input: 14, Except: "XIV"}, {Input: 15, Except: "XV"}, {Input: 16, Except: "XVI"}, {Input: 17, Except: "XVII"}, {Input: 18, Except: "XVIII"}, {Input: 19, Except: "XIX"}, {Input: 20, Except: "XX"}, {Input: 30, Except: "XXX"},
		{Input: 40, Except: "XL"}, {Input: 49, Except: "XLIX"}, {Input: 50, Except: "L"}, {Input: 58, Except: "LVIII"}, {Input: 60, Except: "LX"}, {Input: 70, Except: "LXX"}, {Input: 80, Except: "LXXX"}, {Input: 90, Except: "XC"}, {Input: 99, Except: "XCIX"}, {Input: 100, Except: "C"},
		{Input: 101, Except: "CI"}, {Input: 102, Except: "CII"}, {Input: 199, Except: "CXCIX"}, {Input: 200, Except: "CC"}, {Input: 300, Except: "CCC"}, {Input: 400, Except: "CD"}, {Input: 500, Except: "D"}, {Input: 600, Except: "DC"}, {Input: 700, Except: "DCC"}, {Input: 800, Except: "DCCC"},
		{Input: 900, Except: "CM"}, {Input: 999, Except: "CMXCIX"}, {Input: 1000, Except: "M"}, {Input: 1400, Except: "MCD"}, {Input: 1437, Except: "MCDXXXVII"}, {Input: 1500, Except: "MD"}, {Input: 1800, Except: "MDCCC"}, {Input: 1900, Except: "MCM"}, {Input: 1994, Except: "MCMXCIV"},
		{Input: 2000, Except: "MM"}, {Input: 3000, Except: "MMM"}, {Input: 3333, Except: "MMMCCCXXXIII"}, {Input: 3999, Except: "MMMCMXCIX"},
	}
}

type Func func(num int) string

var funcs = tests.NewFunc[Input, Output]()

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			int(in),
		))
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func AddCases(c func() []tests.Case[Input, Output]) {
	_cases := cases()
	cases = func() []tests.Case[Input, Output] {
		return append(_cases, c()...)
	}
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
