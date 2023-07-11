package keyboard_row

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	words []string
}

type Output []string

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input:  Input{words: []string{"Hello", "Alaska", "Dad", "Peace"}},
			Except: []string{"Alaska", "Dad"},
		},
		{
			Input:  Input{words: []string{"omk"}},
			Except: []string{},
		},
		{
			Input:  Input{words: []string{"adsdf", "sfd"}},
			Except: []string{"adsdf", "sfd"},
		},
	}
}

type Func func(words []string) []string

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.words,
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

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
