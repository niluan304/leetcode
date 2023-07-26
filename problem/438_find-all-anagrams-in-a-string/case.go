package find_all_anagrams_in_a_string

import (
	"reflect"
	"sort"
	"testing"

	"leetcode/tests"
)

type Input struct {
	s string
	p string
}

type Output []int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input:  Input{s: "cbaebabacd", p: "abc"},
			Except: Output{0, 6},
		},
		{
			Input:  Input{s: "abab", p: "ab"},
			Except: Output{0, 1, 2},
		},
	}
}

type Func func(s string, p string) []int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.s,
			in.p,
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
	if len(x) == 0 && len(y) == 0 {
		return true
	}

	sort.Ints(x)
	sort.Ints(y)

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
