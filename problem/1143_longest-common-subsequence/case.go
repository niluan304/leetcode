package longest_common_subsequence

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	text1 string
	text2 string
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				text1: "abcde",
				text2: "ace",
			},
			Except: 3,
		},
		{
			Input: Input{
				text1: "abc",
				text2: "abc",
			},
			Except: 3,
		},
		{
			Input: Input{
				text1: "abc",
				text2: "def",
			},
			Except: 0,
		},
	}
}

type Func func(text1 string, text2 string) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.text1,
			in.text2,
		))
	}
}

func AddCases(c func() []tests.Case[Input, Output]) {
	_cases := cases()
	cases = func() []tests.Case[Input, Output] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Equal(x, y Output) bool {

	return reflect.DeepEqual(x, y)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}
