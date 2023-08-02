package swap_for_longest_repeated_character_substring

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	text string
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input:  Input{text: "ababa"},
			Except: 3,
		},
		{
			Input:  Input{text: "aaabaaa"},
			Except: 6,
		},
		{
			Input:  Input{text: "aaabbaaa"},
			Except: 4,
		},
		{
			Input:  Input{text: "aaaaa"},
			Except: 5,
		},
		{
			Input:  Input{text: "abcdef"},
			Except: 1,
		},
		{
			Input:  Input{text: "aaabaaa"},
			Except: 6,
		},
	}
}

type Func func(text string) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.text,
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