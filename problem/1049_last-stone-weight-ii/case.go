package last_stone_weight_ii

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	stones []int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				stones: []int{2, 7, 4, 1, 8, 1},
			},
			Except: 1,
		},
		{
			Input: Input{
				stones: []int{31, 26, 33, 21, 40},
			},
			Except: 5,
		},
	}
}

type Func func(stones []int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.stones,
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