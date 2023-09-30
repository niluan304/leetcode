package interleaving_string

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	s1 string
	s2 string
	s3 string
}

type Output bool

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbcbcac",
			},
			Except: true,
		},
		{
			Input: Input{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbbaccc",
			},
			Except: false,
		},
		{
			Input: Input{
				s1: "",
				s2: "",
				s3: "",
			},
			Except: true,
		},
	}
}

type Func func(s1 string, s2 string, s3 string) bool

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.s1,
			in.s2,
			in.s3,
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
