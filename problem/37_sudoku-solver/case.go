package sudoku_solver

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	board [][]byte
}

type Output [][]byte

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{board: [][]byte{
				[]byte("53." + ".7." + "..."),
				[]byte("6.." + "195" + "..."),
				[]byte(".98" + "..." + ".6."),
				[]byte("8.." + ".6." + "..3"),
				[]byte("4.." + "8.3" + "..1"),
				[]byte("7.." + ".2." + "..6"),
				[]byte(".6." + "..." + "28."),
				[]byte("..." + "419" + "..5"),
				[]byte("..." + ".8." + ".79"),
			}},
			Except: Output{
				[]byte("534" + "678" + "912"),
				[]byte("672" + "195" + "348"),
				[]byte("198" + "342" + "567"),
				[]byte("859" + "761" + "423"),
				[]byte("426" + "853" + "791"),
				[]byte("713" + "924" + "856"),
				[]byte("961" + "537" + "284"),
				[]byte("287" + "419" + "635"),
				[]byte("345" + "286" + "179"),
			},
		},
	}
}

type Func func(board [][]byte)

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		if ignore {
			f(in.board)
			return in.board
		}

		var output = make([][]byte, len(in.board))
		for i, rows := range in.board {
			output[i] = append([]byte{}, rows...)
		}

		f(
			output,
		)
		return output
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

var ignore = false

func Bench(b *testing.B) {
	ignore = true
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}
