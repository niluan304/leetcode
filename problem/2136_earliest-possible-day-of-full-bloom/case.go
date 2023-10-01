package earliest_possible_day_of_full_bloom

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	plantTime []int
	growTime  []int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				plantTime: []int{1, 4, 3},
				growTime:  []int{2, 3, 1},
			},
			Except: 9,
		},
		{
			Input: Input{
				plantTime: []int{1, 2, 3, 2},
				growTime:  []int{2, 1, 2, 1},
			},
			Except: 9,
		},
		{
			Input: Input{
				plantTime: []int{1},
				growTime:  []int{1},
			},
			Except: 2,
		},
	}
}

type Func func(plantTime []int, growTime []int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.plantTime,
			in.growTime,
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
