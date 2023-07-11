package word_search

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	Board [][]byte
	Word  string
}

var cases = func() []tests.Case[Input, bool] {
	return []tests.Case[Input, bool]{
		{
			Input: Input{Board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFCS"),
				[]byte("ADEE")}, Word: "ABCCED"},
			Except: true,
		},
		{
			Input: Input{Board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFCS"),
				[]byte("ADEE")}, Word: "SEE"},
			Except: true,
		},
		{
			Input: Input{Board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFCS"),
				[]byte("ADEE")}, Word: "ABCB"},
			Except: false,
		},
		{
			Input: Input{Board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFES"),
				[]byte("ADEE")}, Word: "ABCESEEEFS"},
			Except: true,
		},
		{
			Input: Input{Board: [][]byte{
				[]byte("AA")}, Word: "AAA"},
			Except: false,
		},
		{
			Input: Input{Board: [][]byte{
				[]byte("AB")}, Word: "BA"},
			Except: true,
		},
	}
}

type Func func(board [][]byte, word string) bool

func adaptor(f Func) func(in Input) (out bool) {
	return func(in Input) (out bool) {
		return f(in.Board, in.Word)
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[Input, bool]) {
	_cases := cases()
	cases = func() []tests.Case[Input, bool] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, bool]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[Input, bool]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
