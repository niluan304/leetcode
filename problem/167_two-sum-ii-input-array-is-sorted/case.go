package two_sum_ii_input_array_is_sorted

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	numbers []int
	target  int
}

type Output []int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input:  Input{numbers: []int{2, 7, 11, 15}, target: 9},
			Except: Output{1, 2},
		},
		{
			Input:  Input{numbers: []int{2, 3, 4}, target: 6},
			Except: Output{1, 3},
		},
		{
			Input:  Input{numbers: []int{-1, 0}, target: -1},
			Except: Output{1, 2},
		},
	}
}

type Func func(numbers []int, target int) []int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.numbers,
			in.target,
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
