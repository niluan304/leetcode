package replace_the_substring_for_balanced_string

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	s string
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input:  Input{s: "QWER"},
			Except: 0,
		},
		{
			Input:  Input{s: "QQWE"},
			Except: 1,
		},
		{
			Input:  Input{s: "QQQW"},
			Except: 2,
		},
		{
			Input:  Input{s: "QQQQ"},
			Except: 3,
		},
	}
}

type Func func(s string) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.s,
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
